package promise

import (
	"testing"
	"time"

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

	values := make(chan int)

	r.OnComplete(func(_ attempt.Attempt[string]) {
		values <- 1
	}).OnComplete(func(_ attempt.Attempt[string]) {
		values <- 2
	}).OnComplete(func(_ attempt.Attempt[string]) {
		values <- 3
	})

	res := make(chan int)

	go func() {
		sum := 0
		sum += <-values
		sum += <-values
		sum += <-values
		res <- sum
	}()

	go func() {
		w.Success("foo")
	}()

	select {
	case <-time.After(10 * time.Second):
		t.Errorf("test timed out")
	case res := <-res:
		if res != 6 {
			t.Errorf("expected 6, got %d", res)
		}
	}

}
