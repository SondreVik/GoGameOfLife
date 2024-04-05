package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	globalScreenWidth  = 320
	globalScreenHeight = 240
	tps                = 1
	scale              = 4
)

type Game struct {
	state *image.RGBA
}

func (g *Game) Update() error {
	flipBit(g.state, 0, 0, true)
	flipBit(g.state, 1, 1, true)
	flipBit(g.state, 2, 2, true)
	return nil
}

func flipBit(img *image.RGBA, x, y int, active bool) *image.RGBA {
	for xi := 0; xi < scale; xi++ {
		for yi := 0; yi < scale; yi++ {
			img.Pix[(4*xi)+yi*img.Stride+y*img.Stride*scale+(x*4*scale)] = 100
			img.Pix[(4*xi)+yi*img.Stride+y*img.Stride*scale+(x*4*scale)+1] = 100
			img.Pix[(4*xi)+yi*img.Stride+y*img.Stride*scale+(x*4*scale)+2] = 100
			img.Pix[(4*xi)+yi*img.Stride+y*img.Stride*scale+(x*4*scale)+3] = 0xff
		}
	}
	return img
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.state.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return globalScreenWidth, globalScreenHeight
}

func main() {
	ebiten.SetWindowSize(globalScreenWidth*2, globalScreenHeight*2)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(tps)
	if err := ebiten.RunGame(&Game{state: image.NewRGBA(image.Rect(0, 0, globalScreenWidth, globalScreenHeight))}); err != nil {
		log.Fatal(err)
	}
}
