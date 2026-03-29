package atmosphere

import (
	"errors"

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
