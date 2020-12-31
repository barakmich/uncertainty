package uncertainty

import "testing"

const epsilon = 0.1

func TestTwoCoins(t *testing.T) {
	coinA := NewBernoulli(0.5)
	coinB := NewBernoulli(0.5)

	together := Add(coinA, coinB)

	x := Materialize(together, 10000).Average()
	t.Logf("%0.7f", x)
	if !Within(x, 1.0, epsilon) {
		t.Errorf("Expected heads is not 1.0, got %0.7f", x)
	}
}
