package promise

import (
	"testing"

	"github.com/joa/go18beta/attempt"
)

func TestCreate(t *testing.T) {
	w := Create[string]()
	r := w.Future()

	if r.Done() {
		t.Error("promise must not be done")
	}

	if !r.Value().Empty() {
		t.Error("promise result must be empty")
	}

	if !w.TryComplete(attempt.Success("foo")) {
		t.Error("try-complete must succeed")
	}

	if !r.Done() {
		t.Error("promise must be done")
	}

	if r.Value().Empty() {
		t.Error("promise result must not be empty")
	}

	if v := r.Value().Get().Get(); v != "foo" {
		t.Errorf("expected 'foo', got '%s'", v)
	}

	if w.TryComplete(attempt.Success("bar")) {
		t.Error("try-complete must not succeed the second time")
	}

	if !r.Done() {
		t.Error("promise must still be done")
	}

	if r.Value().Empty() {
		t.Error("promise result must still not be empty")
	}

	if v := r.Value().Get().Get(); v != "foo" {
		t.Errorf("expected 'foo', got '%s'", v)
	}
}

func TestOnComplete(t *testing.T) {
	w := Create[string]()
	r := w.Future()

	order := make(chan int, 3)

	r.OnComplete(func(_ attempt.Attempt[string]) {
		order <- 1
	}).OnComplete(func(_ attempt.Attempt[string]) {
		order <- 2
	}).OnComplete(func(_ attempt.Attempt[string]) {
		order <- 3
	})

	select {
	case x := <-order:
		if x != 1 {
			t.Errorf("expected 1, got %d", x)
		}
	}
}
