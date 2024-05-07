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
	if y > len(state) {
		return state
	}
	if x > len(state[0]) {
		return state
	}
	state[y][x] = true
	return state
}
