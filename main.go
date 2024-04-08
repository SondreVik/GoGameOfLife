package main

import (
	"GameOfLife/gameLogic"
	"GameOfLife/gameRenderer"
	"GameOfLife/settings"
	"image"
	"log"
	"math/rand"

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

func initState(width, height int) (state [][]bool) {
	for yi := 0; yi < height; yi++ {
		line := []bool{}
		for xi := 0; xi < width; xi++ {
			line = append(line, randomFlag(width, height))
		}
		state = append(state, line)
	}
	return
}

func randomFlag(width, height int) bool {
	percentedValue := ((settings.RandomPercent / 100) * float32(width) * float32(height))
	modulo := rand.Int() % (width * height)
	result := modulo < int(percentedValue)
	return result
}

func main() {
	ebiten.SetWindowSize(settings.GlobalScreenWidth, settings.GlobalScreenHeight)
	ebiten.SetWindowTitle("Hello, Conway!")
	ebiten.SetTPS(settings.Tps)
	if err := ebiten.RunGame(&Game{canvas: image.NewRGBA(image.Rect(0, 0, settings.GlobalScreenWidth, settings.GlobalScreenHeight)), state: initState(settings.GlobalScreenWidth/settings.Scale, settings.GlobalScreenHeight/settings.Scale)}); err != nil {
		log.Fatal(err)
	}
}
