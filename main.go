package main

import (
	"GameOfLife/drawLogic"
	"GameOfLife/gameInput"
	"GameOfLife/gameLogic"
	"GameOfLife/gameRenderer"
	"GameOfLife/settings"
	"flag"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	randFlag bool
)

type Game struct {
	state         [][]bool
	canvas        *image.RGBA
	input         *gameInput.Input
	isGameRunning bool
}

func (g *Game) Update() error {
	var input = g.input.Update()
	g.isGameRunning = input.SimInProgress
	g.state = drawLogic.UpdateState(g.state, input)
	if g.isGameRunning {
		g.state = gameLogic.RunGame(g.state)
	}
	gameRenderer.RenderGame(g.state, g.canvas)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.canvas.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.GlobalScreenWidth, settings.GlobalScreenHeight
}

func main() {
	ebiten.SetWindowSize(settings.GlobalScreenWidth, settings.GlobalScreenHeight)
	ebiten.SetWindowTitle("Hello, Conway!")
	ebiten.SetTPS(settings.Tps)
	flag.BoolVar(&randFlag, "r", false, "if set, the game will start with random")
	flag.Parse()
	var initialGameState [][]bool
	gameStarted := false
	if randFlag {
		initialGameState = gameLogic.InitGame(settings.GlobalScreenWidth/settings.CellWidth, settings.GlobalScreenHeight/settings.CellWidth)
		gameStarted = true
	} else {
		initialGameState = gameLogic.InitBlankGame(settings.GlobalScreenWidth/settings.CellWidth, settings.GlobalScreenHeight/settings.CellWidth)
	}
	err := ebiten.RunGame(&Game{
		canvas:        image.NewRGBA(image.Rect(0, 0, settings.GlobalScreenWidth, settings.GlobalScreenHeight)),
		state:         initialGameState,
		input:         gameInput.NewInput(gameStarted),
		isGameRunning: gameStarted,
	})
	if err != nil {
		log.Fatal(err)
	}
}
