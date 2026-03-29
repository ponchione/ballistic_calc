package units

import "testing"

func TestPressureConversions(t *testing.T) {
	almostEqual(t, NewPressure(25.4, PressureMillimeterMercury).In(PressureInchMercury), 1, 1e-12)
	almostEqual(t, NewPressure(1, PressureBar).In(PressureMillimeterMercury), 750.061683, 1e-6)
	almostEqual(t, NewPressure(1, PressureHectopascal).In(PressureMillimeterMercury), 750.061683/1000, 1e-9)
	almostEqual(t, NewPressure(1, PressurePSI).In(PressureMillimeterMercury), 51.714924102396, 1e-9)
}
