package common

import (
	"github.com/nemesis567/fyne"
	"github.com/nemesis567/fyne/internal/cache"
)

// CanvasForObject returns the canvas for the specified object.
func CanvasForObject(obj fyne.CanvasObject) fyne.Canvas {
	return cache.GetCanvasForObject(obj)
}
