package gameRenderer

import (
	"GameOfLife/settings"
	"image"
)

func RenderGame(state [][]bool, canvas *image.RGBA) {
	for yi := 0; yi < len(state); yi++ {
		for xi := 0; xi < len(state[yi]); xi++ {
			drawCell(canvas, xi, yi, state[yi][xi])
		}
	}
}

func drawCell(img *image.RGBA, x, y int, active bool) *image.RGBA {
	for yi := 0; yi < settings.CellWidth; yi++ {
		for xi := 0; xi < settings.CellWidth; xi++ {
			color := settings.Dead
			if active {
				color = settings.Alive
			}
			index := (4 * xi) + yi*img.Stride + y*img.Stride*settings.CellWidth + (x * 4 * settings.CellWidth)
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
