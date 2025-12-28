package drawLogic

import (
	"GameOfLife/gameInput"
	"GameOfLife/utils"
)

func UpdateState(state [][]bool, input *gameInput.Input) [][]bool {
	if input.MouseState == gameInput.KeyNone {
		return state
	}
	centerX := utils.NormalizeLength(input.MousePosX)
	centerY := utils.NormalizeLength(input.MousePosY)
	
	// Draw cells in a square pattern based on brush size
	// For brush size N, we draw an NxN square with the top-left corner at the mouse position
	for dy := 0; dy < input.BrushSize; dy++ {
		for dx := 0; dx < input.BrushSize; dx++ {
			x := centerX + dx
			y := centerY + dy
			
			// Check for valid indices
			if y < 0 || y >= len(state) {
				continue
			}
			if x < 0 || x >= len(state[0]) {
				continue
			}
			
			state[y][x] = true
		}
	}
	
	return state
}
