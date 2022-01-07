package try

import (
	"errors"
	"testing"
)

func TestFlatMap(t *testing.T) {
	s := Success("xxx")
	f := Failure[string](errors.New("yyy"))
	fun := func(t string) Try[int] { return Success(len(t + t)) }

	a := FlatMap(s, fun)
	b := FlatMap(f, fun)

	switch {
	case a.Failure():
		t.Error("a.Reject() must be false")
	case a.Err() != nil:
		t.Error("a.Err() must be nil")
	case a.Success() != true:
		t.Error("a.Resolve() must be true")
	case a.Must() != 6:
		t.Errorf("a.Must() must return '6', got '%v'", a.Must())
	}

	switch {
	case b.Failure() != true:
		t.Error("b.Reject() must be true")
	case b.Err() == nil:
		t.Error("b.Err() must not be nil")
	case b.Success():
		t.Error("b.Resolve() must be false")
	}
}
