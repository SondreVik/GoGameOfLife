package main

import (
	"GameOfLife/drawLogic"
	"GameOfLife/gameInput"
	"GameOfLife/gameLogic"
	"GameOfLife/gameRenderer"
	"GameOfLife/settings"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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
	err := ebiten.RunGame(&Game{
		canvas: image.NewRGBA(image.Rect(0, 0, settings.GlobalScreenWidth, settings.GlobalScreenHeight)),
		state:  gameLogic.InitBlankGame(settings.GlobalScreenWidth/settings.CellWidth, settings.GlobalScreenHeight/settings.CellWidth),
		// state:  gameLogic.InitGame(settings.GlobalScreenWidth/settings.CellWidth, settings.GlobalScreenHeight/settings.CellWidth),
		input: gameInput.NewInput()})
	if err != nil {
		log.Fatal(err)
	}
}
