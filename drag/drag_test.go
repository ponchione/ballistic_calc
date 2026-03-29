package drag_test

import (
	"testing"

	"github.com/ponchione/ballistic_calc/drag"
)

func TestNewStandardTableReturnsDefinitionForEachSupportedTable(t *testing.T) {
	tables := []drag.Table{
		drag.TableG1,
		drag.TableG2,
		drag.TableG5,
		drag.TableG6,
		drag.TableG7,
		drag.TableG8,
		drag.TableGS,
		drag.TableRA4,
	}

	for _, table := range tables {
		t.Run(string(table), func(t *testing.T) {
			definition, err := drag.NewStandardTable(table, 0.223)
			if err != nil {
				t.Fatalf("NewStandardTable(%q) returned error: %v", table, err)
			}

			if got := definition.Kind(); got != drag.KindStandardTable {
				t.Fatalf("Kind() = %v, want %v", got, drag.KindStandardTable)
			}

			if got := definition.Table(); got != table {
				t.Fatalf("Table() = %v, want %v", got, table)
			}

			if got := definition.ValueType(); got != drag.ValueTypeBC {
				t.Fatalf("ValueType() = %v, want %v", got, drag.ValueTypeBC)
			}

			if got := definition.Value(); got != 0.223 {
				t.Fatalf("Value() = %v, want 0.223", got)
			}

			if got := definition.RawFunction(); got != nil {
				t.Fatalf("RawFunction() = %v, want nil", got)
			}
		})
	}
}

func TestNewStandardTableRejectsUnknownIdentifierWithoutPanic(t *testing.T) {
	definition, err := drag.NewStandardTable(drag.Table("NOPE"), 0.223)
	if err == nil {
		t.Fatal("NewStandardTable() error = nil, want rejection")
	}

	if got := definition.Kind(); got != 0 {
		t.Fatalf("Kind() = %v, want zero definition after rejection", got)
	}
}

func TestNewCustomFunctionReturnsCustomFunctionDefinition(t *testing.T) {
	rawFunction := func(mach float64) float64 {
		return 0.25 + 0.01*mach
	}

	definition := drag.NewCustomFunction(drag.ValueTypeFF, 1.184, rawFunction)

	if got := definition.Kind(); got != drag.KindCustomFunction {
		t.Fatalf("Kind() = %v, want %v", got, drag.KindCustomFunction)
	}

	if got := definition.ValueType(); got != drag.ValueTypeFF {
		t.Fatalf("ValueType() = %v, want %v", got, drag.ValueTypeFF)
	}

	if got := definition.Value(); got != 1.184 {
		t.Fatalf("Value() = %v, want 1.184", got)
	}

	if definition.RawFunction() == nil {
		t.Fatal("RawFunction() = nil, want custom function")
	}

	if got := definition.RawFunction()(1.0); got != 0.26 {
		t.Fatalf("RawFunction()(1.0) = %v, want 0.26", got)
	}
}

func TestNewCurveFunctionWrapsCurveEvaluatorAsCustomRawFunction(t *testing.T) {
	points := []drag.CurvePoint{
		{Mach: 0.70, Drag: 0.119},
		{Mach: 0.85, Drag: 0.120},
	}
	coefficients := drag.CurveCoefficients{}
	called := false

	rawFunction := drag.NewCurveFunction(points, coefficients, func(gotPoints []drag.CurvePoint, gotCoefficients drag.CurveCoefficients, mach float64) float64 {
		called = true

		if len(gotPoints) != len(points) {
			t.Fatalf("len(points) = %d, want %d", len(gotPoints), len(points))
		}
		for i := range points {
			if gotPoints[i] != points[i] {
				t.Fatalf("points[%d] = %#v, want %#v", i, gotPoints[i], points[i])
			}
		}

		if gotCoefficients != coefficients {
			t.Fatalf("coefficients = %#v, want %#v", gotCoefficients, coefficients)
		}

		if mach != 0.8 {
			t.Fatalf("mach = %v, want 0.8", mach)
		}

		return 0.123
	})

	definition := drag.NewCustomFunction(drag.ValueTypeBC, 1.0, rawFunction)

	if got := definition.RawFunction()(0.8); got != 0.123 {
		t.Fatalf("RawFunction()(0.8) = %v, want 0.123", got)
	}

	if !called {
		t.Fatal("curve evaluator was not called")
	}
}
