package test

import "github.com/nemesis567/fyne"

type testClipboard struct {
	content string
}

func (c *testClipboard) Content() string {
	return c.content
}

func (c *testClipboard) SetContent(content string) {
	c.content = content
}

// NewClipboard returns a single use in-memory clipboard used for testing
func NewClipboard() fyne.Clipboard {
	return &testClipboard{}
}
