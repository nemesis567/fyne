//go:build js || wasm || test_web_driver
// +build js wasm test_web_driver

package glfw

func (d *gLDriver) SetSystemTrayMenu(m *fyne.Menu) {
	// no-op for mobile apps using this driver
}
