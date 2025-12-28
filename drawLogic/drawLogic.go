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
	
	// Calculate the brush area based on brush size
	brushRadius := input.BrushSize / 2
	
	// Draw cells in a square around the center point
	for dy := -brushRadius; dy <= brushRadius; dy++ {
		for dx := -brushRadius; dx <= brushRadius; dx++ {
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
