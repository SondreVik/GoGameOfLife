package main

import (
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	globalScreenWidth        = 320
	globalScreenHeight       = 240
	tps                      = 30
	scale                    = 2
	rgbMax             uint8 = 0xff
	rgbMin             uint8 = 0
	randomPercent            = 100
)

type Game struct {
	state  [][]bool
	canvas *image.RGBA
}

func (g *Game) Update() error {
	g.state = runGame(g.state)
	renderGame(g.state, g.canvas)
	return nil
}

func runGame(state [][]bool) [][]bool {
	for yi := 0; yi < len(state); yi++ {
		for xi := 0; xi < len(state[yi]); xi++ {
			state[yi][xi] = newPixelState(countLiveNeighbors(state, xi, yi), state[yi][xi])
		}
	}
	return state
}

func renderGame(state [][]bool, canvas *image.RGBA) {
	for yi := 0; yi < len(state); yi++ {
		for xi := 0; xi < len(state[yi]); xi++ {
			drawPixel(canvas, xi, yi, state[yi][xi])
		}
	}
}

func countLiveNeighbors(state [][]bool, x, y int) (liveNeighbors int) {
	startx := max(x-1, 0)
	starty := max(y-1, 0)
	endy := min(y+2, len(state))
	for yi := starty; yi < endy; yi++ {
		endx := min(x+2, len(state[yi]))
		for xi := startx; xi < endx; xi++ {
			if yi == x && xi == y {
				continue
			}
			if state[yi][xi] {
				liveNeighbors = liveNeighbors + 1
			}
		}
	}
	return
}

func newPixelState(liveNeighbors int, originalState bool) bool {
	if liveNeighbors < 2 {
		return false
	}
	if originalState && liveNeighbors == 2 {
		return true
	}
	if liveNeighbors == 3 {
		return true
	}
	return false
}

func drawPixel(img *image.RGBA, x, y int, active bool) *image.RGBA {
	for yi := 0; yi < scale; yi++ {
		for xi := 0; xi < scale; xi++ {
			color := rgbMin
			if active {
				color = rgbMax
			}
			index := (4 * xi) + yi*img.Stride + y*img.Stride*scale + (x * 4 * scale)
			if img.Pix[index] == color {
				continue
			}
			img.Pix[index] = color
			img.Pix[index+1] = color
			img.Pix[index+2] = color
			img.Pix[index+3] = rgbMax
		}
	}
	return img
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.canvas.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return globalScreenWidth, globalScreenHeight
}

func initState(x, y int) (state [][]bool) {
	for yi := 0; yi < y; yi++ {
		line := []bool{}
		for xi := 0; xi < x; xi++ {
			line = append(line, randomFlag(x, y))
		}
		state = append(state, line)
	}
	return
}

func randomFlag(x, y int) bool {
	return rand.Int()%(x*y) < ((randomPercent / 100) * x * y)
}

func main() {
	ebiten.SetWindowSize(globalScreenWidth*2, globalScreenHeight*2)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(tps)
	if err := ebiten.RunGame(&Game{canvas: image.NewRGBA(image.Rect(0, 0, globalScreenWidth, globalScreenHeight)), state: initState(globalScreenWidth/4, globalScreenHeight/4)}); err != nil {
		log.Fatal(err)
	}
}
