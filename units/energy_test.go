package units

import "testing"

func TestEnergyConversions(t *testing.T) {
	almostEqual(t, NewEnergy(1, EnergyJoule).In(EnergyFootPound), 0.737562149277, 1e-12)
	almostEqual(t, NewEnergy(1, EnergyFootPound).In(EnergyJoule), 1/0.737562149277, 1e-12)
}
