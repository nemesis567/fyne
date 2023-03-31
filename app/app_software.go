//go:build ci
// +build ci

package app

import (
	"github.com/nemesis567/fyne"
	"github.com/nemesis567/fyne/internal/painter/software"
	"github.com/nemesis567/fyne/test"
)

// NewWithID returns a new app instance using the test (headless) driver.
// The ID string should be globally unique to this app.
func NewWithID(id string) fyne.App {
	return newAppWithDriver(test.NewDriverWithPainter(software.NewPainter()), id)
}
