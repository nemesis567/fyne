package app

import (
	"os"

	"github.com/nemesis567/fyne"
	"github.com/nemesis567/fyne/internal"
	"github.com/nemesis567/fyne/storage"
)

type store struct {
	*internal.Docs
	a *fyneApp
}

func (s *store) RootURI() fyne.URI {
	if s.a.UniqueID() == "" {
		fyne.LogError("Storage API requires a unique ID, use app.NewWithID()", nil)
		return storage.NewFileURI(os.TempDir())
	}

	return storage.NewFileURI(s.a.storageRoot())
}

func (s *store) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
