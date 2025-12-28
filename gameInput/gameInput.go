package gameInput

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type keyState int

const (
	KeyNone keyState = iota
	KeyDown
	KeyUp
)

type Input struct {
	MouseState         keyState
	MousePosX          int
	MousePosY          int
	SimInProgress      bool
	SimInProgressState keyState
	BrushSize          int
	BrushSizeUpState   keyState
	BrushSizeDownState keyState
}

func NewInput(simInProgress bool) *Input {
	return &Input{SimInProgress: simInProgress, BrushSize: 1}
}

func (i *Input) Update() *Input {
	switch i.MouseState {
	case KeyNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			i.MousePosX = x
			i.MousePosY = y
			i.MouseState = KeyDown
		}
	case KeyDown:
		x, y := ebiten.CursorPosition()
		i.MousePosX = x
		i.MousePosY = y
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			i.MouseState = KeyUp
		}
	case KeyUp:
		i.MouseState = KeyNone
	}
	switch i.SimInProgressState {
	case KeyNone:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			i.SimInProgress = !i.SimInProgress
			i.SimInProgressState = KeyDown
		}
	case KeyDown:
		if !ebiten.IsKeyPressed(ebiten.KeySpace) {
			i.SimInProgressState = KeyNone
		}
	}
	switch i.BrushSizeUpState {
	case KeyNone:
		if ebiten.IsKeyPressed(ebiten.KeyEqual) || ebiten.IsKeyPressed(ebiten.KeyKPAdd) {
			if i.BrushSize < 10 {
				i.BrushSize++
			}
			i.BrushSizeUpState = KeyDown
		}
	case KeyDown:
		if !ebiten.IsKeyPressed(ebiten.KeyEqual) && !ebiten.IsKeyPressed(ebiten.KeyKPAdd) {
			i.BrushSizeUpState = KeyNone
		}
	}
	switch i.BrushSizeDownState {
	case KeyNone:
		if ebiten.IsKeyPressed(ebiten.KeyMinus) || ebiten.IsKeyPressed(ebiten.KeyKPSubtract) {
			if i.BrushSize > 1 {
				i.BrushSize--
			}
			i.BrushSizeDownState = KeyDown
		}
	case KeyDown:
		if !ebiten.IsKeyPressed(ebiten.KeyMinus) && !ebiten.IsKeyPressed(ebiten.KeyKPSubtract) {
			i.BrushSizeDownState = KeyNone
		}
	}
	return i
}
