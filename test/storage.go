package test

import (
	"os"

	"github.com/nemesis567/fyne"
	"github.com/nemesis567/fyne/internal"
	"github.com/nemesis567/fyne/storage"
)

type testStorage struct {
	*internal.Docs
}

func (s *testStorage) RootURI() fyne.URI {
	return storage.NewFileURI(os.TempDir())
}

func (s *testStorage) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
