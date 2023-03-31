package desktop

import "github.com/nemesis567/fyne"

// Canvas defines the desktop specific extensions to a fyne.Canvas.
type Canvas interface {
	OnKeyDown() func(*fyne.KeyEvent)
	SetOnKeyDown(func(*fyne.KeyEvent))
	OnKeyUp() func(*fyne.KeyEvent)
	SetOnKeyUp(func(*fyne.KeyEvent))
}
