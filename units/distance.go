package units

type DistanceUnit int

const (
	DistanceInch DistanceUnit = iota
	DistanceFoot
	DistanceYard
	DistanceMile
	DistanceNauticalMile
	DistanceMillimeter
	DistanceCentimeter
	DistanceMeter
	DistanceKilometer
	DistanceLine
)

type Distance struct {
	inches float64
}

func NewDistance(value float64, unit DistanceUnit) Distance {
	switch unit {
	case DistanceInch:
		return Distance{inches: value}
	case DistanceFoot:
		return Distance{inches: value * 12}
	case DistanceYard:
		return Distance{inches: value * 36}
	case DistanceMile:
		return Distance{inches: value * 63360}
	case DistanceNauticalMile:
		return Distance{inches: value * 72913.3858}
	case DistanceMillimeter:
		return Distance{inches: value / 25.4}
	case DistanceCentimeter:
		return Distance{inches: value / 2.54}
	case DistanceMeter:
		return Distance{inches: value * 1000 / 25.4}
	case DistanceKilometer:
		return Distance{inches: value * 1000000 / 25.4}
	case DistanceLine:
		return Distance{inches: value * 0.1}
	default:
		panic("unknown distance unit")
	}
}

func (d Distance) In(unit DistanceUnit) float64 {
	switch unit {
	case DistanceInch:
		return d.inches
	case DistanceFoot:
		return d.inches / 12
	case DistanceYard:
		return d.inches / 36
	case DistanceMile:
		return d.inches / 63360
	case DistanceNauticalMile:
		return d.inches / 72913.3858
	case DistanceMillimeter:
		return d.inches * 25.4
	case DistanceCentimeter:
		return d.inches * 2.54
	case DistanceMeter:
		return d.inches * 25.4 / 1000
	case DistanceKilometer:
		return d.inches * 25.4 / 1000000
	case DistanceLine:
		return d.inches / 0.1
	default:
		panic("unknown distance unit")
	}
}
