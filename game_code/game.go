package main

import (
	"math"
	"math/rand"
	"time"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 500
	screenHeight = 800
)

type Game struct {
	bird         Bird
	op           ebiten.DrawImageOptions
	acceleration float64
	pipesArray   []Pipe
}

type Pipe struct {
	x, y float64
}

type Bird struct {
	image           *ebiten.Image
	x, y            float64
	alive           bool
	is_start_flying bool
	angle           float64
	coins           int
}

var (
	BirdImageUp         *ebiten.Image
	StartMessageImageUp *ebiten.Image
	GameOverImage       *ebiten.Image
	BackgroundImage     *ebiten.Image
	PipeImage           *ebiten.Image

	NumberZeroImage  *ebiten.Image
	NumberOneImage   *ebiten.Image
	NumberTwoImage   *ebiten.Image
	NumberThreeImage *ebiten.Image
	NumberFourImage  *ebiten.Image
	NumberFiveImage  *ebiten.Image
	NumberSixImage   *ebiten.Image
	NumberSevenImage *ebiten.Image
	NumberEightImage *ebiten.Image
	NumberNineImage  *ebiten.Image
)

var (
	BirdImageWidth  float64
	BirdImageHeight float64

	PipeImageWidth  float64
	PipeImageHeight float64
)

func init() {
	var err error
	BirdImageUp, _, err = ebitenutil.NewImageFromFile("sprites/yellowbird-upflap-min.png")
	if err != nil {
		panic(err)
	}
	StartMessageImageUp, _, err = ebitenutil.NewImageFromFile("sprites/message-min.png")
	if err != nil {
		panic(err)
	}
	GameOverImage, _, err = ebitenutil.NewImageFromFile("sprites/gameover-min.png")
	if err != nil {
		panic(err)
	}
	BackgroundImage, _, err = ebitenutil.NewImageFromFile("sprites/background-day-min.png")
	if err != nil {
		panic(err)
	}
	PipeImage, _, err = ebitenutil.NewImageFromFile("sprites/pipe-green (1)-min.png")
	if err != nil {
		panic(err)
	}
	NumberZeroImage, _, err = ebitenutil.NewImageFromFile("sprites/0-min.png")
	if err != nil {
		panic(err)
	}

	NumberOneImage, _, err = ebitenutil.NewImageFromFile("sprites/1-min.png")
	if err != nil {
		panic(err)
	}

	NumberTwoImage, _, err = ebitenutil.NewImageFromFile("sprites/2-min.png")
	if err != nil {
		panic(err)
	}

	NumberThreeImage, _, err = ebitenutil.NewImageFromFile("sprites/3-min.png")
	if err != nil {
		panic(err)
	}

	NumberFourImage, _, err = ebitenutil.NewImageFromFile("sprites/4-min.png")
	if err != nil {
		panic(err)
	}

	NumberFiveImage, _, err = ebitenutil.NewImageFromFile("sprites/5-min.png")
	if err != nil {
		panic(err)
	}

	NumberSixImage, _, err = ebitenutil.NewImageFromFile("sprites/6-min.png")
	if err != nil {
		panic(err)
	}

	NumberSevenImage, _, err = ebitenutil.NewImageFromFile("sprites/7-min.png")
	if err != nil {
		panic(err)
	}

	NumberEightImage, _, err = ebitenutil.NewImageFromFile("sprites/8-min.png")
	if err != nil {
		panic(err)
	}

	NumberNineImage, _, err = ebitenutil.NewImageFromFile("sprites/9-min.png")
	if err != nil {
		panic(err)
	}

	w, h := BirdImageUp.Size()
	BirdImageWidth = float64(w)
	BirdImageHeight = float64(h)

	w, h = PipeImage.Size()
	PipeImageWidth = float64(w)
	PipeImageHeight = float64(h)
}
func ChechCollision(bird Bird, pipe Pipe) bool {
	x1 := pipe.x
	y1 := 0.0
	x2 := pipe.x + PipeImageWidth*2
	y3 := 800 - 100 - pipe.y - 150
	y5 := y3 + 150
	y7 := 800.0

	if bird.angle > 0 {
		b_x1 := bird.x - BirdImageWidth - bird.angle*0.5 + BirdImageWidth
		b_y1 := bird.y - BirdImageHeight*1.5 + bird.angle*10 + BirdImageHeight
		b_x2 := bird.x + BirdImageWidth*0.7 - bird.angle*25 + BirdImageWidth
		b_y2 := bird.y - BirdImageHeight*1.5 + bird.angle*55 + BirdImageHeight
		b_x3 := bird.x - BirdImageWidth - bird.angle*30 + BirdImageWidth
		b_y3 := bird.y + BirdImageHeight/2 - bird.angle*20 + BirdImageHeight
		b_x4 := bird.x + BirdImageWidth - bird.angle*65 + BirdImageWidth
		b_y4 := bird.y + BirdImageHeight/2 + bird.angle*40 + BirdImageHeight

		if (b_x1 > x1 && b_x1 < x2) && (b_y1 > y1 && b_y1 < y3) || (b_x2 > x1 && b_x2 < x2) && (b_y2 > y1 && b_y2 < y3) || (b_x3 > x1 && b_x3 < x2) && (b_y3 > y1 && b_y3 < y3) || (b_x4 > x1 && b_x4 < x2) && (b_y4 > y1 && b_y4 < y3) {
			return true
		} else if (b_x1 > x1 && b_x1 < x2) && (b_y1 > y5 && b_y1 < y7) || (b_x2 > x1 && b_x2 < x2) && (b_y2 > y5 && b_y2 < y7) || (b_x3 > x1 && b_x3 < x2) && (b_y3 > y5 && b_y3 < y7) || (b_x4 > x1 && b_x4 < x2) && (b_y4 > y5 && b_y4 < y7) {
			return true
		} else if b_y1 < 0 || b_y2 < 0 || b_y3 > 800 || b_y4 > 800 {
			return true
		} else {
			return false
		}
	} else {
		b_x1 := bird.x - BirdImageWidth - bird.angle*20 + BirdImageWidth
		b_y1 := bird.y - BirdImageHeight*1.3 + bird.angle*10 + BirdImageHeight
		b_x2 := bird.x + BirdImageWidth*0.5 + BirdImageWidth
		b_y2 := bird.y - BirdImageHeight*1.5 + bird.angle*50 + BirdImageHeight
		b_x3 := bird.x - BirdImageWidth - bird.angle*50 + BirdImageWidth
		b_y3 := bird.y + BirdImageHeight/2 + bird.angle*15 + BirdImageHeight
		b_x4 := bird.x + BirdImageWidth - bird.angle*19 + BirdImageWidth
		b_y4 := bird.y + BirdImageHeight/2 + bird.angle*75 + BirdImageHeight

		if (b_x1 > x1 && b_x1 < x2) && (b_y1 > y1 && b_y1 < y3) || (b_x2 > x1 && b_x2 < x2) && (b_y2 > y1 && b_y2 < y3) || (b_x3 > x1 && b_x3 < x2) && (b_y3 > y1 && b_y3 < y3) || (b_x4 > x1 && b_x4 < x2) && (b_y4 > y1 && b_y4 < y3) {
			return true
		} else if (b_x1 > x1 && b_x1 < x2) && (b_y1 > y5 && b_y1 < y7) || (b_x2 > x1 && b_x2 < x2) && (b_y2 > y5 && b_y2 < y7) || (b_x3 > x1 && b_x3 < x2) && (b_y3 > y5 && b_y3 < y7) || (b_x4 > x1 && b_x4 < x2) && (b_y4 > y5 && b_y4 < y7) {
			return true
		} else if b_y1 < 0 || b_y2 < 0 || b_y3 > 800 || b_y4 > 800 {
			return true
		} else {
			return false
		}
	}
}

func (g *Game) showCoins(screen *ebiten.Image) {
	show_num := func(num int, x float64) {
		g.op.GeoM.Reset()
		g.op.GeoM.Scale(2.0, 2.0)
		g.op.GeoM.Translate(x, 100)

		switch num {
		case 0:
			screen.DrawImage(NumberZeroImage, &g.op)
		case 1:
			screen.DrawImage(NumberOneImage, &g.op)
		case 2:
			screen.DrawImage(NumberTwoImage, &g.op)
		case 3:
			screen.DrawImage(NumberThreeImage, &g.op)
		case 4:
			screen.DrawImage(NumberFourImage, &g.op)
		case 5:
			screen.DrawImage(NumberFiveImage, &g.op)
		case 6:
			screen.DrawImage(NumberSixImage, &g.op)
		case 7:
			screen.DrawImage(NumberSevenImage, &g.op)
		case 8:
			screen.DrawImage(NumberEightImage, &g.op)
		case 9:
			screen.DrawImage(NumberNineImage, &g.op)
		}
	}

	given_coint := g.bird.coins
	var nums []int

	for given_coint > 0 {
		nums = append(nums, given_coint%10)
		given_coint = given_coint / 10
	}

	for i := len(nums) - 1; i >= 0; i-- {
		show_num(nums[i], float64(87+(3-i)*50))
	}
}

func CreateStartedPipes(g *Game) {
	g.pipesArray = nil
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		g.pipesArray = append(g.pipesArray, Pipe{float64(500 + i*300), float64(rand.Intn(450))})
	}
}

func (g *Game) FullReset() {
	g.acceleration = 0
	CreateStartedPipes(g)
	g.bird = Bird{BirdImageUp, screenWidth/2 - BirdImageWidth, screenHeight/2 - BirdImageHeight, true, false, 0, 0}
}

func (g *Game) Jump() {
	g.acceleration = 0
	g.bird.y -= 40
	g.bird.angle = -(math.Pi / 7)
}

func (g *Game) HandleMovement() {
	if inpututil.IsKeyJustPressed(100) || len(inpututil.JustPressedTouchIDs()) > 0 {
		g.Jump()
	}
}

func (g *Game) Update() error {
	if g.bird.is_start_flying {
		if g.bird.alive {
			g.HandleMovement()

			g.acceleration += 0.03
			if g.bird.angle < math.Pi/4 {
				g.bird.angle += (g.acceleration / 10) * (g.acceleration / 10)
			}
			g.bird.y += g.acceleration * g.acceleration

			for i := 0; i < 3; i++ {
				g.pipesArray[i].x -= 3

				if ChechCollision(g.bird, g.pipesArray[i]) {
					g.bird.alive = false
				}
			}

			if g.pipesArray[0].x < -120 {
				for i := 0; i < 2; i++ {
					g.pipesArray[i] = g.pipesArray[i+1]
				}

				rand.Seed(time.Now().UnixNano())
				g.pipesArray[2] = Pipe{g.pipesArray[2].x + 300, float64(rand.Intn(450))}
			}

			if g.bird.x >= g.pipesArray[0].x-1 && g.bird.x <= g.pipesArray[0].x+1 || g.bird.x >= g.pipesArray[1].x-1 && g.bird.x <= g.pipesArray[1].x+1 {
				g.bird.coins += 1
			}

		} else {
			g.bird.angle = -math.Pi / 1.7
			g.bird.x = g.bird.x - 5
			g.bird.y += 2

			if inpututil.IsKeyJustPressed(ebiten.KeySpace) || len(inpututil.JustPressedTouchIDs()) > 0 {
				g.FullReset()
			}
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || len(inpututil.JustPressedTouchIDs()) > 0 {
			g.bird.is_start_flying = true
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.bird.is_start_flying {

		g.op.GeoM.Reset()
		g.op.GeoM.Scale(2.0, 1.7)
		g.op.GeoM.Translate(0, 0)
		screen.DrawImage(BackgroundImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Scale(2.0, 2.0)
		g.op.GeoM.Rotate(g.bird.angle)
		g.op.GeoM.Translate(g.bird.x, g.bird.y)
		screen.DrawImage(g.bird.image, &g.op)

		for _, pipe := range g.pipesArray {
			g.op.GeoM.Reset()
			g.op.GeoM.Scale(2.0, 2.0)
			g.op.GeoM.Translate(pipe.x, 800-100-pipe.y)
			screen.DrawImage(PipeImage, &g.op)

			g.op.GeoM.Reset()
			g.op.GeoM.Scale(2.0, 2.0)
			g.op.GeoM.Translate(pipe.x, -784+(450-pipe.y))
			screen.DrawImage(PipeImage, &g.op)
		}

		g.showCoins(screen)

		if !g.bird.alive {
			g.op.GeoM.Reset()
			g.op.GeoM.Scale(2.0, 2.0)
			g.op.GeoM.Rotate(g.bird.angle)
			g.op.GeoM.Translate(g.bird.x, g.bird.y)
			screen.DrawImage(g.bird.image, &g.op)

			g.op.GeoM.Reset()
			g.op.GeoM.Scale(2.0, 2.0)
			g.op.GeoM.Translate(63, 351)
			screen.DrawImage(GameOverImage, &g.op)

			g.showCoins(screen)
		}

	} else {
		g.op.GeoM.Reset()
		g.op.GeoM.Scale(2.0, 1.7)
		g.op.GeoM.Translate(0, 0)
		screen.DrawImage(BackgroundImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Scale(2.0, 2.0)
		g.op.GeoM.Translate(66, 38)
		screen.DrawImage(StartMessageImageUp, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Scale(2.0, 2.0)
		g.op.GeoM.Translate(g.bird.x, g.bird.y)
		screen.DrawImage(g.bird.image, &g.op)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	return &Game{
		bird: Bird{BirdImageUp, screenWidth/2 - BirdImageWidth, screenHeight/2 - BirdImageHeight, true, false, 0, 0},
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Flappy Bird")
	ebiten.SetMaxTPS(60)
	ng := NewGame()
	CreateStartedPipes(ng)
	if err := ebiten.RunGame(ng); err != nil {
		panic(err)
	}

}
