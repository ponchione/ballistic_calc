package atmosphere

import (
	"errors"
	"math"

	"github.com/ponchione/ballistic_calc/units"
)

var errInvalidHumidity = errors.New("humidity must be in [0, 1] or (1, 100]")

type Atmosphere struct {
	altitude    units.Distance
	pressure    units.Pressure
	temperature units.Temperature
	humidity    float64
}

func Default() Atmosphere {
	return Atmosphere{
		altitude:    units.NewDistance(0, units.DistanceFoot),
		pressure:    units.NewPressure(29.92, units.PressureInchMercury),
		temperature: units.NewTemperature(59, units.TemperatureFahrenheit),
		humidity:    0.78,
	}
}

func New(altitude units.Distance, pressure units.Pressure, temperature units.Temperature, humidity float64) (Atmosphere, error) {
	switch {
	case humidity < 0 || humidity > 100:
		return Atmosphere{}, errInvalidHumidity
	case humidity > 1:
		humidity /= 100
	}

	return Atmosphere{
		altitude:    altitude,
		pressure:    pressure,
		temperature: temperature,
		humidity:    humidity,
	}, nil
}

func (a Atmosphere) Altitude() units.Distance {
	return a.altitude
}

func (a Atmosphere) Pressure() units.Pressure {
	return a.pressure
}

func (a Atmosphere) Temperature() units.Temperature {
	return a.temperature
}

func (a Atmosphere) Humidity() float64 {
	return a.humidity
}

func (a Atmosphere) DensityFactor() float64 {
	temperatureFahrenheit := a.temperature.In(units.TemperatureFahrenheit)
	pressureInchMercury := a.pressure.In(units.PressureInchMercury)

	humidityCorrection := 1.0
	if temperatureFahrenheit > 0 {
		saturationPressure := 1.24871 +
			0.0988438*temperatureFahrenheit +
			0.00152907*math.Pow(temperatureFahrenheit, 2) -
			3.07031e-06*math.Pow(temperatureFahrenheit, 3) +
			4.21329e-07*math.Pow(temperatureFahrenheit, 4)
		vaporPressure := 3.342e-04 * a.humidity * saturationPressure
		humidityCorrection = (pressureInchMercury - 0.3783*vaporPressure) / 29.92
	}

	densityPoundsPerCubicFoot := 0.076474 * (518.67 / (temperatureFahrenheit + 459.67)) * humidityCorrection
	return densityPoundsPerCubicFoot / 0.076474
}

func (a Atmosphere) SpeedOfSound() units.Velocity {
	temperatureFahrenheit := a.temperature.In(units.TemperatureFahrenheit)
	speedOfSoundFeetPerSecond := math.Sqrt(temperatureFahrenheit+459.67) * 49.0223
	return units.NewVelocity(speedOfSoundFeetPerSecond, units.VelocityFootPerSecond)
}
