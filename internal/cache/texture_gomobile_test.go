//go:build android || ios || mobile
// +build android ios mobile

package cache_test

import (
	"sync"
	"testing"

	"github.com/nemesis567/fyne/internal/cache"
	"github.com/nemesis567/fyne/internal/driver/mobile/gl"
)

// go test -race
func TestTextureGomobile(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		cache.SetTexture(nil, gl.Texture{0}, nil)
	}()
	go func() {
		defer wg.Done()
		cache.SetTexture(nil, gl.Texture{1}, nil)
	}()

	wg.Wait()
}
