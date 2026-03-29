package atmosphere

import (
	"math"

	"github.com/ponchione/ballistic_calc/units"
)

const (
	icaoLapseRate = -3.56616e-03
	icaoExponent  = -5.255876
)

func NewICAO(altitude units.Distance) Atmosphere {
	altitudeFeet := altitude.In(units.DistanceFoot)
	temperatureFahrenheit := 518.67 + altitudeFeet*icaoLapseRate - 459.67
	pressureInchMercury := 29.92 * math.Pow(518.67/(temperatureFahrenheit+459.67), icaoExponent)

	return Atmosphere{
		altitude:    altitude,
		pressure:    units.NewPressure(pressureInchMercury, units.PressureInchMercury),
		temperature: units.NewTemperature(temperatureFahrenheit, units.TemperatureFahrenheit),
		humidity:    0,
	}
}
