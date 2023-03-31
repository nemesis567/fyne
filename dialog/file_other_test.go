//go:build !windows
// +build !windows

package dialog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nfyne/storage"
)

func TestIsHidden(t *testing.T) {
	assert.True(t, isHidden(storage.NewFileURI(".nothing")))
	assert.False(t, isHidden(storage.NewFileURI("file.txt")))
}

func TestFileDialog_LoadPlaces(t *testing.T) {
	f := &fileDialog{}
	places := f.getPlaces()

	assert.Equal(t, 1, len(places))
	assert.Equal(t, "Computer", places[0].locName)
}
