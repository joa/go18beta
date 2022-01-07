package future

import (
	"errors"
	"testing"
	"time"

	"github.com/joa/go18beta/try"
)

func TestSuccess(t *testing.T) {
	w := PromiseOf[string](try.Success("foo"))
	r := w.Future()

	if !r.Done() {
		t.Error("promise must be done")
	}

	if r.Value().Empty() {
		t.Error("promise result must not be empty")
	}

	if w.TryComplete(try.Success("bar")) {
		t.Error("try-complete must not succeed")
	}

	ch := make(chan string)

	r.Then(func(s string) {
		ch <- s
	}).Catch(func(_ error) {
		t.Fatal("must not call catch")
	})

	select {
	case <-time.After(10 * time.Second):
		t.Error("test timed out")
	case s := <-ch:
		if s != "foo" {
			t.Errorf("expected 'foo', got '%s'", s)
		}
	}
}
func TestFailure(t *testing.T) {
	err := errors.New("fail")
	w := PromiseOf[string](try.Failure[string](err))
	r := w.Future()

	if !r.Done() {
		t.Error("promise must be done")
	}

	if r.Value().Empty() {
		t.Error("promise result must not be empty")
	}

	if w.TryComplete(try.Success("bar")) {
		t.Error("try-complete must not succeed")
	}

	ch := make(chan error)

	r.Then(func(s string) {
		t.Fatal("must not call then")
	}).Catch(func(err error) {
		ch <- err
	})

	select {
	case <-time.After(10 * time.Second):
		t.Error("test timed out")
	case e := <-ch:
		if e != err {
			t.Errorf("expected '%s', got '%s'", err, e)
		}
	}
}
