//go:build js || wasm || test_web_driver
// +build js wasm test_web_driver

package glfw

import "github.com/nemesis567/fyne"

func (d *gLDriver) SetSystemTrayMenu(m *fyne.Menu) {
	// no-op for mobile apps using this driver
}
