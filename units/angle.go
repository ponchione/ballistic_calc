package units

import "math"

type AngleUnit int

const (
	AngleRadian AngleUnit = iota
	AngleDegree
	AngleMOA
	AngleMil
	AngleMilliradian
	AngleThousand
	AngleInPer100Yard
	AngleCmPer100Meter
)

type Angle struct {
	radians float64
}

func NewAngle(value float64, unit AngleUnit) Angle {
	switch unit {
	case AngleRadian:
		return Angle{radians: value}
	case AngleDegree:
		return Angle{radians: value / 180 * math.Pi}
	case AngleMOA:
		return Angle{radians: value / 180 * math.Pi / 60}
	case AngleMil:
		return Angle{radians: value / 3200 * math.Pi}
	case AngleMilliradian:
		return Angle{radians: value / 1000}
	case AngleThousand:
		return Angle{radians: value / 3000 * math.Pi}
	case AngleInPer100Yard:
		return Angle{radians: math.Atan(value / 3600)}
	case AngleCmPer100Meter:
		return Angle{radians: math.Atan(value / 10000)}
	default:
		panic("unknown angle unit")
	}
}

func (a Angle) In(unit AngleUnit) float64 {
	switch unit {
	case AngleRadian:
		return a.radians
	case AngleDegree:
		return a.radians * 180 / math.Pi
	case AngleMOA:
		return a.radians * 180 / math.Pi * 60
	case AngleMil:
		return a.radians / math.Pi * 3200
	case AngleMilliradian:
		return a.radians * 1000
	case AngleThousand:
		return a.radians / math.Pi * 3000
	case AngleInPer100Yard:
		return math.Tan(a.radians) * 3600
	case AngleCmPer100Meter:
		return math.Tan(a.radians) * 10000
	default:
		panic("unknown angle unit")
	}
}
