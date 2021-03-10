package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var main_temp = template.Must(template.ParseFiles("main_page.html"))
var game_temp = template.Must(template.ParseFiles("game_page.html"))

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "helloworld",
	DB:       0, // use default DB
})

func mainPage(w http.ResponseWriter, r *http.Request) {
	err := main_temp.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func birdGame(w http.ResponseWriter, r *http.Request) {
	err := game_temp.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveUserResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		fmt.Println(r.Form)
		user_coins_str := r.Form.Get("user_coins_str")
		user_coins_arr := strings.Split(user_coins_str, "~")
		if len(user_coins_arr) != 2 {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		username := user_coins_arr[0]
		coins, err := strconv.ParseFloat(user_coins_arr[1], 64)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		if user_score, err := rdb.ZScore(ctx, "bird_leaderboard", username).Result(); err != nil {
			if coins != 0 && username != "Anonymous" && username != "" {
				rdb.ZAdd(ctx, "bird_leaderboard", &redis.Z{Score: coins, Member: username})
			}
		} else {
			if coins > user_score {
				rdb.ZAdd(ctx, "bird_leaderboard", &redis.Z{Score: coins, Member: username})
			}
		}

	}
}

func getLeaderboardAndUserResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		username, ok := r.URL.Query()["username"]
		if len(username) != 1 || !ok {
			http.Error(w, "The request must contain one GET parameter 'username'", http.StatusBadRequest)
			return
		}

		leaderboard, err := rdb.ZRevRangeWithScores(ctx, "bird_leaderboard", 0, 9).Result()
		if err != nil {
			http.Error(w, "Something wrong in DataBase", http.StatusInternalServerError)
			return
		}
		user_rank, err := rdb.ZRevRank(ctx, "bird_leaderboard", username[0]).Result()
		if err != nil {
			leaderboard_json, err := json.Marshal(leaderboard)
			if err != nil {
				http.Error(w, "Something wrong in leaderboard", http.StatusInternalServerError)
				return
			}
			w.Write(leaderboard_json)
		} else {
			user_score, _ := rdb.ZScore(ctx, "bird_leaderboard", username[0]).Result()
			leaderboard := []interface{}{leaderboard, map[string]string{"Member": username[0], "Score": strconv.FormatInt(int64(user_score), 10), "Rank": strconv.FormatInt(user_rank, 10)}}
			leaderboard_json, err := json.Marshal(leaderboard)
			if err != nil {
				http.Error(w, "Something wrong in leaderboard", http.StatusInternalServerError)
				return
			}
			w.Write(leaderboard_json)

		}
	}
}

func main() {
	server_port := flag.Int("port", 8000, "The port on which the server is running")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", mainPage)
	r.HandleFunc("/bird_game", birdGame)
	r.HandleFunc("/save_user_result", saveUserResult)
	r.HandleFunc("/get_leaderboard", getLeaderboardAndUserResult)
	http.Handle("/", r)

	log.Println("Server start at port " + strconv.Itoa(*server_port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*server_port), nil))
}
