package drag

type Kind int

const (
	KindStandardTable Kind = iota
	KindCustomFunction
)

type ValueType int

const (
	ValueTypeBC ValueType = iota
	ValueTypeFF
)

type Table string

const (
	TableG1  Table = "G1"
	TableG2  Table = "G2"
	TableG5  Table = "G5"
	TableG6  Table = "G6"
	TableG7  Table = "G7"
	TableG8  Table = "G8"
	TableGS  Table = "GS"
	TableRA4 Table = "RA4"
)

type RawFunction func(mach float64) float64

type CurvePoint struct {
	Mach float64
	Drag float64
}

type CurveCoefficients struct{}

type CurveEvaluator func(points []CurvePoint, coefficients CurveCoefficients, mach float64) float64

type Definition struct {
	kind        Kind
	table       Table
	valueType   ValueType
	value       float64
	rawFunction RawFunction
}

func NewStandardTable(table Table, bc float64) Definition {
	return Definition{
		kind:      KindStandardTable,
		table:     table,
		valueType: ValueTypeBC,
		value:     bc,
	}
}

func NewCustomFunction(valueType ValueType, value float64, rawFunction RawFunction) Definition {
	return Definition{
		kind:        KindCustomFunction,
		valueType:   valueType,
		value:       value,
		rawFunction: rawFunction,
	}
}

func NewCurveFunction(points []CurvePoint, coefficients CurveCoefficients, evaluator CurveEvaluator) RawFunction {
	return func(mach float64) float64 {
		return evaluator(points, coefficients, mach)
	}
}

func (d Definition) Kind() Kind {
	return d.kind
}

func (d Definition) Table() Table {
	return d.table
}

func (d Definition) ValueType() ValueType {
	return d.valueType
}

func (d Definition) Value() float64 {
	return d.value
}

func (d Definition) RawFunction() RawFunction {
	return d.rawFunction
}
