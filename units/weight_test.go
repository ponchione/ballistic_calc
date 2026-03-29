package units

import "testing"

func TestWeightConversions(t *testing.T) {
	almostEqual(t, NewWeight(1, WeightGram).In(WeightGrain), 15.4323584, 1e-9)
	almostEqual(t, NewWeight(1, WeightKilogram).In(WeightGrain), 15432.3584, 1e-6)
	almostEqual(t, NewWeight(1, WeightNewton).In(WeightGrain), 151339.73750336, 1e-6)
	almostEqual(t, NewWeight(1, WeightOunce).In(WeightGrain), 437.5, 1e-12)
	almostEqual(t, NewWeight(1, WeightPound).In(WeightGrain), 7000, 1e-9)
}
