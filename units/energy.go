package units

type EnergyUnit int

const (
	EnergyFootPound EnergyUnit = iota
	EnergyJoule
)

type Energy struct {
	footPounds float64
}

func NewEnergy(value float64, unit EnergyUnit) Energy {
	switch unit {
	case EnergyFootPound:
		return Energy{footPounds: value}
	case EnergyJoule:
		return Energy{footPounds: value * 0.737562149277}
	default:
		panic("unknown energy unit")
	}
}

func (e Energy) In(unit EnergyUnit) float64 {
	switch unit {
	case EnergyFootPound:
		return e.footPounds
	case EnergyJoule:
		return e.footPounds / 0.737562149277
	default:
		panic("unknown energy unit")
	}
}
