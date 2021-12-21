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
}
