package predicate

import "testing"

func TestAlwaysTrue(t *testing.T) {
	if !AlwaysTrue[string]().Apply("") {
		t.Errorf("AlwaysTrue must be true")
	}
}

func TestAlwaysFalse(t *testing.T) {
	if AlwaysFalse[string]().Apply("") {
		t.Errorf("AlwaysFalse must be false")
	}
}

func TestNot(t *testing.T) {
	if Not(AlwaysTrue[string]()).Apply("") {
		t.Errorf("Not(AlwaysTrue) must be false")
	}
}

func TestOr(t *testing.T) {
	if !Or(AlwaysFalse[string](), AlwaysTrue[string]()).Apply("") {
		t.Errorf("Or(AlwaysFalse, AlwaysTrue) must be true")
	}
}

func TestAnd(t *testing.T) {
	if And(AlwaysFalse[string](), AlwaysTrue[string]()).Apply("") {
		t.Errorf("And(AlwaysFalse, AlwaysTrue) must be true")
	}
}

func TestNotNil(t *testing.T) {
	type Foo struct{}

	if NotNil[*Foo]().Apply(nil) {
		t.Error("NotNil(nil) must be false")
	}

	if !NotNil[*Foo]().Apply(&Foo{}) {
		t.Error("NotNil(Foo{}) must be true")
	}
}
