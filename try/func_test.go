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
		t.Error("a.Failure() must be false")
	case a.Err() != nil:
		t.Error("a.Err() must be nil")
	case a.Success() != true:
		t.Error("a.Success() must be true")
	case a.Must() != 6:
		t.Errorf("a.Must() must return '6', got '%v'", a.Must())
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

func TestFlatMapPanic(t *testing.T) {
	res := FlatMap(Success("foo"), func(x string) Try[int] {
		// Try[T] must recover from panics
		panic("at the discotheque")
	})

	switch {
	case res.Failure() != true:
		t.Error("b.Failure() must be true")
	case res.Err() == nil:
		t.Error("b.Err() must not be nil")
	case res.Success():
		t.Error("b.Success() must be false")
	}

	if res.Err().Error() != "at the discotheque" {
		t.Errorf("got unexpected error: %s", res.Err())
	}
}
