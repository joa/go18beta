package future

import (
	"errors"
	"testing"
	"time"

	"github.com/joa/go18beta/pair"
)

func TestFlatMap(t *testing.T) {
	s := Resolve("xxx")
	f := Reject[string](errors.New("yyy"))
	fun := func(t string) Future[int] { return Resolve[int](len(t + t)) }

	a := FlatMap(s, fun)
	b := FlatMap(f, fun)

	done := make(chan bool, 1)

	a.Then(func(d int) {
		if d != 6 {
			t.Errorf("expected 6, got %d", d)
		}
		done <- true
	}).Catch(func(err error) {
		done <- true
		t.Fatal("must not fail")
	})

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Error("test timed out")
	}

	b.Then(func(d int) {
		done <- true
		t.Fatal("must fail")
	}).Catch(func(err error) {
		if err.Error() != "yyy" {
			t.Errorf("expected err, got %v", b.Value().Must())
		}
		done <- true
	})

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Error("test timed out")
	}
}

func TestFlatMapPanic(t *testing.T) {
	fut := FlatMap(Resolve("xxx"), func(t string) Future[int] { panic("err") })
	done := make(chan bool, 1)

	fut.Then(func(d int) {
		done <- true
		t.Fatal("must fail")
	}).Catch(func(err error) {
		if err.Error() != "err" {
			t.Errorf("expected err, got %s", err)
		}
		done <- true
	})

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Error("test timed out")
	}
}

func TestJoin(t *testing.T) {
	a := Create[string]()
	b := Create[string]()

	c := a.Future()
	d := b.Future()

	go func() {
		b.Resolve("b")
		a.Resolve("a")
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
