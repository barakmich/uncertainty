package uncertainty

import "testing"

func TestConstantSample(t *testing.T) {
	c := NewConstant(5.0)
	m := Materialize(c, 10)
	if m.First() != 5.0 {
		t.Fatal("Couldn't retrieve constant after sampling")
	}
}

func TestConstantBNNSample(t *testing.T) {
	a := NewConstant(5.0)
	b := NewConstant(6.0)
	c := Add(a, b)

	m := Materialize(c, 10)
	if m.First() != 11.0 {
		t.Fatal("Addition on constants failed")
	}
}

func TestConstantBernoulliSample(t *testing.T) {
	a := NewConstant(5.0)
	b := NewConstant(6.0)
	c := LessThan(b, a)

	m := Materialize(c, 10)
	if m.First() != 0.0 {
		t.Fatal("LessThan on constants failed")
	}
}

func TestBernoulliConditional(t *testing.T) {
	x := NewConstant(5.0)
	y := NewConstant(6.0)

	if LessThan(y, x).Pr() {
		t.Fatal("6 tests less than 5")
	}
	if !LessThan(x, y).Pr() {
		t.Fatal("5 tests not less than 6")
	}
}

func TestBernoulliEqual(t *testing.T) {
	x := NewConstant(5.0)
	y := NewConstant(5.0)

	if LessThan(x, y).Pr() {
		t.Fatal("x is less than y, incorrectly (they should be equal)")
	}
	if LessThan(y, x).Pr() {
		t.Fatal("y is less than x, incorrectly (they should be equal)")
	}
}
