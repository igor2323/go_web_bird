# Funny Bird

Funny Bird is a small browser game

![Bird image](sprites/yellowbird-upflap.png)

## Download and Install

After downloading repository, run the file "server".

    ./server

You can use --help to see a list of parameters

    ./server --help

My server is designed for nginx or another similar server to serve static files, sprites and the game.

## FAQ

Some information about how this project is made

### Game

The game is written in the Go language using the Ebiten library. WASM was used to launch the game in the browser

### Server

The server was also written in the Go language. The project uses the standard "net / http" package

