package units

import "testing"

func TestTemperatureConversions(t *testing.T) {
	almostEqual(t, NewTemperature(32, TemperatureFahrenheit).In(TemperatureCelsius), 0, 1e-12)
	almostEqual(t, NewTemperature(0, TemperatureCelsius).In(TemperatureFahrenheit), 32, 1e-12)
	almostEqual(t, NewTemperature(273.15, TemperatureKelvin).In(TemperatureFahrenheit), 32, 1e-12)
	almostEqual(t, NewTemperature(491.67, TemperatureRankine).In(TemperatureFahrenheit), 32, 1e-12)
}
