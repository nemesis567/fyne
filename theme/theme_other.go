//go:build !hints
// +build !hints

package theme

import (
	"image/color"
)

var (
	fallbackColor = color.Transparent
	fallbackIcon  = &fyne.StaticResource{}
)
