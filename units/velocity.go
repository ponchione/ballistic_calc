package units

type VelocityUnit int

const (
	VelocityMeterPerSecond VelocityUnit = iota
	VelocityKilometerPerHour
	VelocityFootPerSecond
	VelocityMilePerHour
	VelocityKnot
)

type Velocity struct {
	feetPerSecond float64
}

func NewVelocity(value float64, unit VelocityUnit) Velocity {
	switch unit {
	case VelocityMeterPerSecond:
		return Velocity{feetPerSecond: value * 3.2808399}
	case VelocityKilometerPerHour:
		return Velocity{feetPerSecond: (value / 3.6) * 3.2808399}
	case VelocityFootPerSecond:
		return Velocity{feetPerSecond: value}
	case VelocityMilePerHour:
		return Velocity{feetPerSecond: (value / 2.23693629) * 3.2808399}
	case VelocityKnot:
		return Velocity{feetPerSecond: (value / 1.94384449) * 3.2808399}
	default:
		panic("unknown velocity unit")
	}
}

func (v Velocity) In(unit VelocityUnit) float64 {
	switch unit {
	case VelocityMeterPerSecond:
		return v.feetPerSecond / 3.2808399
	case VelocityKilometerPerHour:
		return (v.feetPerSecond / 3.2808399) * 3.6
	case VelocityFootPerSecond:
		return v.feetPerSecond
	case VelocityMilePerHour:
		return (v.feetPerSecond / 3.2808399) * 2.23693629
	case VelocityKnot:
		return (v.feetPerSecond / 3.2808399) * 1.94384449
	default:
		panic("unknown velocity unit")
	}
}
