package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nfyne/canvas"
	"nfyne/internal/cache"
	"nfyne/theme"
)

func TestTabButton_Icon_Change(t *testing.T) {
	b := &tabButton{icon: theme.CancelIcon()}
	r := cache.Renderer(b)
	icon := r.Objects()[3].(*canvas.Image)
	oldResource := icon.Resource

	b.icon = theme.ConfirmIcon()
	b.Refresh()
	assert.NotEqual(t, oldResource, icon.Resource)
}
