package atmosphere

import "github.com/ponchione/ballistic_calc/units"

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
