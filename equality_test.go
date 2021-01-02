package uncertainty

import "testing"

func TestGaussianEquality(t *testing.T) {
	x := NewGaussian(1.0, 1.0)
	y := NewGaussian(4.0, 2.0)

	if GreaterThan(x, y).Pr() {
		t.Error("x > y")
	}
	if LessThan(y, x).Pr() {
		t.Error("y < x")
	}
	if !GreaterThan(y, x).Pr() {
		t.Error("!y > x")
	}
	if !LessThan(x, y).Pr() {
		t.Error("!x < y")
	}
}
