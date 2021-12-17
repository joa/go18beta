package attempt

import (
	"errors"
	"testing"
)

func TestFlatMap(t *testing.T) {
	s := Success("xxx")
	f := Failure[string](errors.New("yyy"))
	fun := func(t string) Attempt[string] { return Success(t + t) }

	a := FlatMap(s, fun)
	b := FlatMap(f, fun)

	switch {
	case a.Failure():
		t.Error("a.Failure() must be false")
	case a.Err() != nil:
		t.Error("a.Err() must be nil")
	case a.Success() != true:
		t.Error("a.Success() must be true")
	case a.Get() != "xxxxxx":
		t.Errorf("a.Get() must return 'xxxxxx', got '%s'", a.Get())
	}

	switch {
	case b.Failure() != true:
		t.Error("b.Failure() must be true")
	case b.Err() == nil:
		t.Error("b.Err() must not be nil")
	case b.Success():
		t.Error("b.Success() must be false")
	}
}
