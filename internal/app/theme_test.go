package app_test

import (
	"testing"

	"github.com/nemesis567/fyne/internal/app"
	"github.com/nemesis567/fyne/test"
)

func TestApplySettings_BeforeContentSet(t *testing.T) {
	a := test.NewApp()
	_ = a.NewWindow("NoContent")

	app.ApplySettings(a.Settings(), a)
}
