package software

import (
	"github.com/nemesis567/fyne/internal/painter/software"
	"github.com/nemesis567/fyne/test"
)

// NewCanvas creates a new canvas in memory that can render without hardware support.
func NewCanvas() test.WindowlessCanvas {
	return test.NewCanvasWithPainter(software.NewPainter())
}

// NewTransparentCanvas creates a new canvas in memory that can render without hardware support without a background color.
//
// Since: 2.2
func NewTransparentCanvas() test.WindowlessCanvas {
	return test.NewTransparentCanvasWithPainter(software.NewPainter())
}
