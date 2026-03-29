package atmosphere

import (
	"testing"

	"github.com/ponchione/ballistic_calc/units"
)

func TestDefaultReturnsStandardAtmosphere(t *testing.T) {
	atmosphere := Default()

	if got := atmosphere.Altitude().In(units.DistanceFoot); got != 0 {
		t.Fatalf("altitude = %v ft, want 0", got)
	}

	if got := atmosphere.Pressure().In(units.PressureInchMercury); got != 29.92 {
		t.Fatalf("pressure = %v inHg, want 29.92", got)
	}

	if got := atmosphere.Temperature().In(units.TemperatureFahrenheit); got != 59 {
		t.Fatalf("temperature = %v F, want 59", got)
	}

	if got := atmosphere.Humidity(); got != 0.78 {
		t.Fatalf("humidity = %v, want 0.78", got)
	}
}
