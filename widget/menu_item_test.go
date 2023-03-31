package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nemesis567/fyne"
	"github.com/nemesis567/fyne/internal/cache"
	"github.com/nemesis567/fyne/theme"
)

func TestMenuItem_Disabled(t *testing.T) {
	i := fyne.NewMenuItem("Disabled", func() {})
	m := fyne.NewMenu("top", []*fyne.MenuItem{i}...)
	i.Disabled = true
	w := newMenuItem(i, NewMenu(m))
	r := cache.Renderer(w)

	assert.Equal(t, theme.DisabledColor(), r.(*menuItemRenderer).text.Color)
}
