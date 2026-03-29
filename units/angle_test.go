package units

import (
	"fmt"
	"math"
	"testing"
)

func TestAngleRadDegMOA(t *testing.T) {
	almostEqual(t, NewAngle(1, AngleRadian).In(AngleDegree), 180/math.Pi, 1e-9)
	almostEqual(t, NewAngle(180, AngleDegree).In(AngleRadian), math.Pi, 1e-12)
	almostEqual(t, NewAngle(1, AngleMOA).In(AngleRadian), math.Pi/10800, 1e-12)
}

func TestAngleMilMradThousand(t *testing.T) {
	almostEqual(t, NewAngle(1, AngleMil).In(AngleRadian), math.Pi/3200, 1e-12)
	almostEqual(t, NewAngle(1, AngleMilliradian).In(AngleRadian), 0.001, 1e-12)
	almostEqual(t, NewAngle(1, AngleThousand).In(AngleRadian), math.Pi/3000, 1e-12)
}

func TestAngleLinearPerDistanceUnitsUseGeometry(t *testing.T) {
	almostEqual(t, NewAngle(1, AngleInPer100Yard).In(AngleRadian), math.Atan(1.0/3600.0), 1e-12)
	almostEqual(t, NewAngle(1, AngleCmPer100Meter).In(AngleRadian), math.Atan(1.0/10000.0), 1e-12)
}

func TestAngleLinearPerDistanceRoundTrips(t *testing.T) {
	almostEqual(t, NewAngle(1, AngleInPer100Yard).In(AngleInPer100Yard), 1, 1e-12)
	almostEqual(t, NewAngle(1, AngleCmPer100Meter).In(AngleCmPer100Meter), 1, 1e-12)
}

func TestAngularSpecInvariants(t *testing.T) {
	a := NewAngle(1, AngleInPer100Yard)
	almostEqual(t, a.In(AngleMOA), 0.954930, 1e-5)

	got := fmt.Sprintf("%.2fcm/100m", a.In(AngleCmPer100Meter))
	if got != "2.78cm/100m" {
		t.Fatalf("got %q want %q", got, "2.78cm/100m")
	}
}
