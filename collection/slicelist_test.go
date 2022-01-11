package collection

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var x SliceList[string]
	y := List[string](&x)

	x.Add("foo")
	y.Add("bar")
	y.Add("baz")

	for i := y.Iterator(); i.HasNext(); {
		fmt.Println(i.Next())
	}

	i := y.Iterator()
	exp := []string{"foo", "bar", "baz"}
	for _, e := range exp {
		if act := i.Next(); act != e {
			t.Errorf("expected %s, got %s", e, act)
		}
	}
}
