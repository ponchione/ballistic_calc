package units

type TemperatureUnit int

const (
	TemperatureFahrenheit TemperatureUnit = iota
	TemperatureCelsius
	TemperatureKelvin
	TemperatureRankine
)

type Temperature struct {
	fahrenheit float64
}

func NewTemperature(value float64, unit TemperatureUnit) Temperature {
	switch unit {
	case TemperatureFahrenheit:
		return Temperature{fahrenheit: value}
	case TemperatureCelsius:
		return Temperature{fahrenheit: value*9/5 + 32}
	case TemperatureKelvin:
		return Temperature{fahrenheit: (value-273.15)*9/5 + 32}
	case TemperatureRankine:
		return Temperature{fahrenheit: value - 459.67}
	default:
		panic("unknown temperature unit")
	}
}

func (t Temperature) In(unit TemperatureUnit) float64 {
	switch unit {
	case TemperatureFahrenheit:
		return t.fahrenheit
	case TemperatureCelsius:
		return (t.fahrenheit - 32) * 5 / 9
	case TemperatureKelvin:
		return (t.fahrenheit-32)*5/9 + 273.15
	case TemperatureRankine:
		return t.fahrenheit + 459.67
	default:
		panic("unknown temperature unit")
	}
}
