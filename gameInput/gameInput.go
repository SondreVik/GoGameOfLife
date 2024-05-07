package gameInput

import "github.com/hajimehoshi/ebiten/v2"

type keyState int

const (
	KeyNone keyState = iota
	KeyDown
	KeyUp
)

type Input struct {
	MouseState    keyState
	MousePosX     int
	MousePosY     int
	SimInProgress bool
}

func NewInput(simInProgress bool) *Input {
	return &Input{SimInProgress: simInProgress}
}

func (i *Input) Update() *Input {
	switch i.MouseState {
	case KeyNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			i.MouseState = KeyDown
		}
	case KeyDown:
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			i.MousePosX = x
			i.MousePosY = y
			i.MouseState = KeyUp
		}
	case KeyUp:
		i.MouseState = KeyNone
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		i.SimInProgress = !i.SimInProgress
	}
	return i
}
