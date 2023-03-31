package desktop

import "github.com/nemesis567/fyne"

// App defines the desktop specific extensions to a fyne.App.
//
// Since: 2.2
type App interface {
	SetSystemTrayMenu(menu *fyne.Menu)
	SetSystemTrayIcon(icon fyne.Resource)
}
