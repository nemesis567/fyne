//go:build !darwin && !js && !wasm && (android || ios || mobile)
// +build !darwin
// +build !js
// +build !wasm
// +build android ios mobile

package gl

import (
	"github.com/nemesis567/fyne/internal/driver/mobile/gl"
)

const (
	singleChannelColorFormat = gl.LUMINANCE
)

var _ = singleChannelColorFormat
