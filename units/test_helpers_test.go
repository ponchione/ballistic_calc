package units

import (
	"math"
	"testing"
)

func almostEqual(t *testing.T, got, want, tol float64) {
	t.Helper()
	if math.Abs(got-want) > tol {
		t.Fatalf("got %v want %v tol %v", got, want, tol)
	}
}
