package gameLogic

import (
	"GameOfLife/settings"
	"math/rand"
)

func RunGame(state [][]bool) (newState [][]bool) {
	for yi := 0; yi < len(state); yi++ {
		line := []bool{}
		for xi := 0; xi < len(state[yi]); xi++ {
			line = append(line, newCellState(countLiveNeighbors(state, xi, yi), state[yi][xi]))
		}
		newState = append(newState, line)
	}
	return
}

func InitGame(width, height int) (state [][]bool) {
	for yi := 0; yi < height; yi++ {
		line := []bool{}
		for xi := 0; xi < width; xi++ {
			line = append(line, randomFlag(width, height))
		}
		state = append(state, line)
	}
	return
}

func InitBlankGame(width, height int) (state [][]bool) {
	for yi := 0; yi < height; yi++ {
		line := []bool{}
		for xi := 0; xi < width; xi++ {
			line = append(line, false)
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

func newCellState(liveNeighbors int, originalState bool) bool {
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

func countLiveNeighbors(state [][]bool, x, y int) (liveNeighbors int) {
	startx := max(x-1, 0)
	starty := max(y-1, 0)
	endy := min(y+2, len(state))
	for yi := starty; yi < endy; yi++ {
		endx := min(x+2, len(state[yi]))
		for xi := startx; xi < endx; xi++ {
			if yi == y && xi == x {
				continue
			}
			if state[yi][xi] {
				liveNeighbors = liveNeighbors + 1
			}
		}
	}
	return
}
