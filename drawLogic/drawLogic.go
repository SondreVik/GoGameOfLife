package drawLogic

import (
	"GameOfLife/gameInput"
	"GameOfLife/utils"
)

func UpdateState(state [][]bool, input *gameInput.Input) [][]bool {
	if input.MouseState != gameInput.KeyUp {
		return state
	}
	x := utils.NormalizeLength(input.MousePosX)
	y := utils.NormalizeLength(input.MousePosY)
	state[y][x] = !state[y][x]
	return state
}
