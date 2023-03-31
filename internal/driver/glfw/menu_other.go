//go:build !darwin || js || wasm || test_web_driver || no_native_menus
// +build !darwin js wasm test_web_driver no_native_menus

package glfw

import "github.com/nemesis567/fyne"

func hasNativeMenu() bool {
	return false
}

func setupNativeMenu(_ *window, _ *fyne.MainMenu) {
	// no-op
}
