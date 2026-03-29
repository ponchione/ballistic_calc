package atmosphere

import (
	"math"

	"github.com/ponchione/ballistic_calc/units"
)

const (
	localRefreshThresholdFeet     = 3.28084
	localRefreshComparisonEpsilon = 1e-9
)

type LocalConditions struct {
	base          Atmosphere
	lastAltitude  units.Distance
	densityFactor float64
	speedOfSound  units.Velocity
}

func NewLocalConditions(base Atmosphere) LocalConditions {
	return LocalConditions{
		base:          base,
		lastAltitude:  base.Altitude(),
		densityFactor: base.DensityFactor(),
		speedOfSound:  base.SpeedOfSound(),
	}
}

func (c LocalConditions) DensityFactor() float64 {
	return c.densityFactor
}

func (c LocalConditions) SpeedOfSound() units.Velocity {
	return c.speedOfSound
}

func (c *LocalConditions) UpdateForAltitude(altitude units.Distance) {
	if math.Abs(c.lastAltitude.In(units.DistanceFoot)-altitude.In(units.DistanceFoot)) <= localRefreshThresholdFeet+localRefreshComparisonEpsilon {
		return
	}

	adjusted := adjustedAtmosphereForAltitude(c.base, altitude)
	c.lastAltitude = altitude
	c.densityFactor = adjusted.DensityFactor()
	c.speedOfSound = adjusted.SpeedOfSound()
}

func adjustedAtmosphereForAltitude(base Atmosphere, altitude units.Distance) Atmosphere {
	baseAltitudeFeet := base.Altitude().In(units.DistanceFoot)
	altitudeFeet := altitude.In(units.DistanceFoot)
	if math.Abs(baseAltitudeFeet-altitudeFeet) < 30 {
		return Atmosphere{
			altitude:    altitude,
			pressure:    base.Pressure(),
			temperature: base.Temperature(),
			humidity:    base.Humidity(),
		}
	}

	baseTemperatureFahrenheit := base.Temperature().In(units.TemperatureFahrenheit)
	basePressureInchMercury := base.Pressure().In(units.PressureInchMercury)
	baseReferenceTemperature := 518.67 + baseAltitudeFeet*icaoLapseRate - 459.67
	altitudeReferenceTemperature := 518.67 + altitudeFeet*icaoLapseRate - 459.67
	adjustedTemperatureFahrenheit := baseTemperatureFahrenheit - baseReferenceTemperature + altitudeReferenceTemperature
	adjustedPressureInchMercury := basePressureInchMercury * math.Pow((baseTemperatureFahrenheit+459.67)/(adjustedTemperatureFahrenheit+459.67), icaoExponent)

	return Atmosphere{
		altitude:    altitude,
		pressure:    units.NewPressure(adjustedPressureInchMercury, units.PressureInchMercury),
		temperature: units.NewTemperature(adjustedTemperatureFahrenheit, units.TemperatureFahrenheit),
		humidity:    base.Humidity(),
	}
}
