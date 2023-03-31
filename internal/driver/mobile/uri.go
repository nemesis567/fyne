//go:build !android
// +build !android

package mobile

import (
	"github.com/nemesis567/fyne"
	"github.com/nemesis567/fyne/storage"
)

func nativeURI(path string) fyne.URI {
	uri, err := storage.ParseURI(path)
	if err != nil {
		fyne.LogError("Error on parsing uri", err)
	}
	return uri
}
