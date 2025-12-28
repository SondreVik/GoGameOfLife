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

	// Mouse wheel for brush size
	_, wheelY := ebiten.Wheel()
	if wheelY > 0 && i.BrushSize < 10 {
		i.BrushSize++
	} else if wheelY < 0 && i.BrushSize > 1 {
		i.BrushSize--
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
	i.updateBrushSize(ebiten.KeyE, 1, &i.BrushSizeUpState)
	i.updateBrushSize(ebiten.KeyQ, -1, &i.BrushSizeDownState)
	return i
}

func (i *Input) updateBrushSize(key ebiten.Key, delta int, state *keyState) {
	switch *state {
	case KeyNone:
		// Use E for increase
		if ebiten.IsKeyPressed(key) {
			if i.BrushSize < 10 && delta > 0 {
				i.BrushSize++
			} else if i.BrushSize > 1 && delta < 0 {
				i.BrushSize--
			}
			*state = KeyDown
		}
	case KeyDown:
		if !ebiten.IsKeyPressed(key) {
			*state = KeyNone
		}
	}
}
