package playground

import (
	"github.com/nemesis567/fyne/driver/software"
	"github.com/nemesis567/fyne/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return software.NewCanvas()
}
