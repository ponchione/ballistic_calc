package units

type WeightUnit int

const (
	WeightGrain WeightUnit = iota
	WeightOunce
	WeightGram
	WeightPound
	WeightKilogram
	WeightNewton
)

type Weight struct {
	grains float64
}

func NewWeight(value float64, unit WeightUnit) Weight {
	switch unit {
	case WeightGrain:
		return Weight{grains: value}
	case WeightOunce:
		return Weight{grains: value * 437.5}
	case WeightGram:
		return Weight{grains: value * 15.4323584}
	case WeightPound:
		return Weight{grains: value * 7000}
	case WeightKilogram:
		return Weight{grains: value * 15432.3584}
	case WeightNewton:
		return Weight{grains: value * 151339.73750336}
	default:
		panic("unknown weight unit")
	}
}

func (w Weight) In(unit WeightUnit) float64 {
	switch unit {
	case WeightGrain:
		return w.grains
	case WeightOunce:
		return w.grains / 437.5
	case WeightGram:
		return w.grains / 15.4323584
	case WeightPound:
		return w.grains / 7000
	case WeightKilogram:
		return w.grains / 15432.3584
	case WeightNewton:
		return w.grains / 151339.73750336
	default:
		panic("unknown weight unit")
	}
}
