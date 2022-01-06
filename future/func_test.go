package future

import (
	"testing"
	"time"

	"github.com/joa/go18beta/pair"
)

func TestJoin(t *testing.T) {
	a := Create[string]()
	b := Create[string]()

	c := a.Future()
	d := b.Future()

	go func() {
		b.Success("b")
		a.Success("a")
	}()

	res := make(chan string)

	Join(c, d).Then(func(p pair.Pair[string, string]) {
		res <- p.X + p.Y
	})

	select {
	case <-time.After(10 * time.Second):
		t.Error("test timed out")
	case res := <-res:
		if res != "ab" {
			t.Errorf("expected 'ab', got '%s'", res)
		}
	}
}
