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
	default:
		panic("unknown angle unit")
	}
}
