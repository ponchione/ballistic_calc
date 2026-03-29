package atmosphere

import (
	"math"
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

func TestNewReturnsExplicitAtmosphereForFractionalHumidity(t *testing.T) {
	atmosphere, err := New(
		units.NewDistance(1500, units.DistanceFoot),
		units.NewPressure(27.5, units.PressureInchMercury),
		units.NewTemperature(41, units.TemperatureFahrenheit),
		0.42,
	)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	if got := atmosphere.Altitude().In(units.DistanceFoot); got != 1500 {
		t.Fatalf("altitude = %v ft, want 1500", got)
	}

	if got := atmosphere.Pressure().In(units.PressureInchMercury); got != 27.5 {
		t.Fatalf("pressure = %v inHg, want 27.5", got)
	}

	if got := atmosphere.Temperature().In(units.TemperatureFahrenheit); got != 41 {
		t.Fatalf("temperature = %v F, want 41", got)
	}

	almostEqual(t, atmosphere.Humidity(), 0.42, 1e-12)
}

func TestNewNormalizesPercentHumidity(t *testing.T) {
	atmosphere, err := New(
		units.NewDistance(0, units.DistanceFoot),
		units.NewPressure(29.92, units.PressureInchMercury),
		units.NewTemperature(59, units.TemperatureFahrenheit),
		78,
	)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	almostEqual(t, atmosphere.Humidity(), 0.78, 1e-12)
}

func TestNewRejectsHumidityOutsideSupportedRanges(t *testing.T) {
	tests := []struct {
		name     string
		humidity float64
	}{
		{name: "below zero", humidity: -0.01},
		{name: "above one hundred", humidity: 100.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(
				units.NewDistance(0, units.DistanceFoot),
				units.NewPressure(29.92, units.PressureInchMercury),
				units.NewTemperature(59, units.TemperatureFahrenheit),
				tt.humidity,
			)
			if err == nil {
				t.Fatalf("expected error for humidity %v", tt.humidity)
			}
		})
	}
}

func TestNewICAOReturnsSeaLevelAtmosphere(t *testing.T) {
	atmosphere := NewICAO(units.NewDistance(0, units.DistanceFoot))

	if got := atmosphere.Altitude().In(units.DistanceFoot); got != 0 {
		t.Fatalf("altitude = %v ft, want 0", got)
	}

	almostEqual(t, atmosphere.Temperature().In(units.TemperatureFahrenheit), 59.0, 1e-4)
	almostEqual(t, atmosphere.Pressure().In(units.PressureInchMercury), 29.92, 1e-5)
	almostEqual(t, atmosphere.Humidity(), 0, 0)
}

func TestNewICAOReturnsExpectedSpotCheckAt5000Feet(t *testing.T) {
	atmosphere := NewICAO(units.NewDistance(5000, units.DistanceFoot))

	if got := atmosphere.Altitude().In(units.DistanceFoot); got != 5000 {
		t.Fatalf("altitude = %v ft, want 5000", got)
	}

	almostEqual(t, atmosphere.Temperature().In(units.TemperatureFahrenheit), 41.1692, 1e-4)
	almostEqual(t, atmosphere.Pressure().In(units.PressureInchMercury), 24.89488, 1e-5)
	almostEqual(t, atmosphere.Humidity(), 0, 0)
}

func TestDefaultDensityFactorMatchesSection64Formula(t *testing.T) {
	almostEqual(t, Default().DensityFactor(), 0.9999443715552248, 1e-12)
}

func TestDefaultSpeedOfSoundMatchesSection64Formula(t *testing.T) {
	almostEqual(t, Default().SpeedOfSound().In(units.VelocityFootPerSecond), 1116.4499224539381, 1e-9)
}

func TestNewLocalConditionsStartWithBaseAtmosphereValues(t *testing.T) {
	base, err := New(
		units.NewDistance(500, units.DistanceFoot),
		units.NewPressure(28, units.PressureInchMercury),
		units.NewTemperature(70, units.TemperatureFahrenheit),
		0.5,
	)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	local := NewLocalConditions(base)

	almostEqual(t, local.DensityFactor(), 0.9163427629511216, 1e-12)
	almostEqual(t, local.SpeedOfSound().In(units.VelocityFootPerSecond), 1128.2266945155989, 1e-9)
}

func TestLocalConditionsUpdateRecomputesValuesBeyondThirtyFeetFromBaseAltitude(t *testing.T) {
	base, err := New(
		units.NewDistance(500, units.DistanceFoot),
		units.NewPressure(28, units.PressureInchMercury),
		units.NewTemperature(70, units.TemperatureFahrenheit),
		0.5,
	)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	local := NewLocalConditions(base)
	local.UpdateForAltitude(units.NewDistance(700, units.DistanceFoot))

	almostEqual(t, local.DensityFactor(), 0.9111037704270233, 1e-12)
	almostEqual(t, local.SpeedOfSound().In(units.VelocityFootPerSecond), 1127.466826622273, 1e-9)
}

func TestLocalConditionsUpdateReusesBaseValuesWithinThirtyFeetOfBaseAltitude(t *testing.T) {
	base, err := New(
		units.NewDistance(500, units.DistanceFoot),
		units.NewPressure(28, units.PressureInchMercury),
		units.NewTemperature(70, units.TemperatureFahrenheit),
		0.5,
	)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	local := NewLocalConditions(base)
	local.UpdateForAltitude(units.NewDistance(520, units.DistanceFoot))

	almostEqual(t, local.DensityFactor(), 0.9163427629511216, 1e-12)
	almostEqual(t, local.SpeedOfSound().In(units.VelocityFootPerSecond), 1128.2266945155989, 1e-9)
}

func TestLocalConditionsUpdateRecomputesAtThirtyFeetDifferenceFromBaseAltitude(t *testing.T) {
	base, err := New(
		units.NewDistance(500, units.DistanceFoot),
		units.NewPressure(28, units.PressureInchMercury),
		units.NewTemperature(70, units.TemperatureFahrenheit),
		0.5,
	)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	local := NewLocalConditions(base)
	local.UpdateForAltitude(units.NewDistance(530, units.DistanceFoot))

	almostEqual(t, local.DensityFactor(), 0.9155554495382884, 1e-12)
	almostEqual(t, local.SpeedOfSound().In(units.VelocityFootPerSecond), 1128.1127469606085, 1e-9)
}

func TestLocalConditionsUpdateRefreshesOnlyAfterAltitudeChangesByMoreThanOneMeter(t *testing.T) {
	base, err := New(
		units.NewDistance(500, units.DistanceFoot),
		units.NewPressure(28, units.PressureInchMercury),
		units.NewTemperature(70, units.TemperatureFahrenheit),
		0.5,
	)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	local := NewLocalConditions(base)
	local.UpdateForAltitude(units.NewDistance(700, units.DistanceFoot))
	almostEqual(t, local.DensityFactor(), 0.9111037704270233, 1e-12)
	almostEqual(t, local.SpeedOfSound().In(units.VelocityFootPerSecond), 1127.466826622273, 1e-9)

	local.UpdateForAltitude(units.NewDistance(703.28084, units.DistanceFoot))
	almostEqual(t, local.DensityFactor(), 0.9111037704270233, 1e-12)
	almostEqual(t, local.SpeedOfSound().In(units.VelocityFootPerSecond), 1127.466826622273, 1e-9)

	local.UpdateForAltitude(units.NewDistance(704, units.DistanceFoot))
	almostEqual(t, local.DensityFactor(), 0.9109992246965521, 1e-12)
	almostEqual(t, local.SpeedOfSound().In(units.VelocityFootPerSecond), 1127.4516240407038, 1e-9)
}

func almostEqual(t *testing.T, got, want, tol float64) {
	t.Helper()
	if math.Abs(got-want) > tol {
		t.Fatalf("got %v want %v tol %v", got, want, tol)
	}
}
