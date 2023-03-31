package glfw

import "github.com/nemesis567/fyne"

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
}
