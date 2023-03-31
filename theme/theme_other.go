//go:build !hints
// +build !hints

package theme

import (
	"image/color"

	"github.com/nemesis567/fyne"
)

var (
	fallbackColor = color.Transparent
	fallbackIcon  = &fyne.StaticResource{}
)
