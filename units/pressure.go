package units

type PressureUnit int

const (
	PressureMillimeterMercury PressureUnit = iota
	PressureInchMercury
	PressureBar
	PressureHectopascal
	PressurePSI
)

type Pressure struct {
	inchesMercury float64
}

func NewPressure(value float64, unit PressureUnit) Pressure {
	switch unit {
	case PressureMillimeterMercury:
		return Pressure{inchesMercury: value / 25.4}
	case PressureInchMercury:
		return Pressure{inchesMercury: value}
	case PressureBar:
		return Pressure{inchesMercury: 750.061683 / 25.4 * value}
	case PressureHectopascal:
		return Pressure{inchesMercury: (750.061683 / 1000) / 25.4 * value}
	case PressurePSI:
		return Pressure{inchesMercury: 51.714924102396 / 25.4 * value}
	default:
		panic("unknown pressure unit")
	}
}

func (p Pressure) In(unit PressureUnit) float64 {
	switch unit {
	case PressureMillimeterMercury:
		return p.inchesMercury * 25.4
	case PressureInchMercury:
		return p.inchesMercury
	case PressureBar:
		return (p.inchesMercury * 25.4) / 750.061683
	case PressureHectopascal:
		return (p.inchesMercury * 25.4) / (750.061683 / 1000)
	case PressurePSI:
		return (p.inchesMercury * 25.4) / 51.714924102396
	default:
		panic("unknown pressure unit")
	}
}
