package main

import (
	"GameOfLife/gameLogic"
	"GameOfLife/gameRenderer"
	"GameOfLife/settings"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	state  [][]bool
	canvas *image.RGBA
}

func (g *Game) Update() error {
	g.state = gameLogic.RunGame(g.state)
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
	if err := ebiten.RunGame(&Game{canvas: image.NewRGBA(image.Rect(0, 0, settings.GlobalScreenWidth, settings.GlobalScreenHeight)), state: gameLogic.InitGame(settings.GlobalScreenWidth/settings.Scale, settings.GlobalScreenHeight/settings.Scale)}); err != nil {
		log.Fatal(err)
	}
}
