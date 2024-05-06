package utils

import "GameOfLife/settings"

func NormalizeLength(length int) int {
	return length / settings.CellWidth
}
