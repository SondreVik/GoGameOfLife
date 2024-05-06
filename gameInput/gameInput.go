package gameInput

import "github.com/hajimehoshi/ebiten/v2"

type mouseState int

const (
	MouseNone mouseState = iota
	MouseDown
	MouseUp
)

type Input struct {
	MouseState mouseState
	MousePosX  int
	MousePosY  int
	RPressed   bool
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update() *Input {
	switch i.MouseState {
	case MouseNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			i.MouseState = MouseDown
		}
	case MouseDown:
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			i.MousePosX = x
			i.MousePosY = y
			i.MouseState = MouseUp
		}
	case MouseUp:
		i.MouseState = MouseNone
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		i.RPressed = true
	}
	return i
}
