package main

import (
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	globalScreenWidth          = 1920
	globalScreenHeight         = 1080
	tps                        = 10
	scale                      = 4
	alive              uint8   = 0
	dead               uint8   = 0xff
	randomPercent      float32 = 5
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
	if liveNeighbors > 3 {
		return false
	}
	return originalState
}

func drawPixel(img *image.RGBA, x, y int, active bool) *image.RGBA {
	for yi := 0; yi < scale; yi++ {
		for xi := 0; xi < scale; xi++ {
			color := dead
			if active {
				color = alive
			}
			index := (4 * xi) + yi*img.Stride + y*img.Stride*scale + (x * 4 * scale)
			if img.Pix[index] == color {
				continue
			}
			img.Pix[index] = color
			img.Pix[index+1] = color
			img.Pix[index+2] = color
			img.Pix[index+3] = 0xff
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
	percentedValue := ((randomPercent / 100) * float32(width) * float32(height))
	modulo := rand.Int() % (width * height)
	result := modulo < int(percentedValue)
	return result
}

func main() {
	ebiten.SetWindowSize(globalScreenWidth, globalScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(tps)
	if err := ebiten.RunGame(&Game{canvas: image.NewRGBA(image.Rect(0, 0, globalScreenWidth, globalScreenHeight)), state: initState(globalScreenWidth/scale, globalScreenHeight/scale)}); err != nil {
		log.Fatal(err)
	}
}
