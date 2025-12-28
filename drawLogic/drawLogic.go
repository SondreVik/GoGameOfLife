package drawLogic

import (
	"GameOfLife/gameInput"
	"GameOfLife/utils"
)

func UpdateState(state [][]bool, input *gameInput.Input) [][]bool {
	if input.MouseState == gameInput.KeyNone {
		return state
	}
	x := utils.NormalizeLength(input.MousePosX)
	y := utils.NormalizeLength(input.MousePosY)
	
	// Check for valid indices (including negative values)
	if y < 0 || y >= len(state) {
		return state
	}
	if x < 0 || x >= len(state[0]) {
		return state
	}
	
	state[y][x] = true
	return state
}
