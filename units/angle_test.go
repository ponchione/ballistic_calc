package units

import (
	"math"
	"testing"
)

func TestAngleRadDegMOA(t *testing.T) {
	almostEqual(t, NewAngle(1, AngleRadian).In(AngleDegree), 180/math.Pi, 1e-9)
	almostEqual(t, NewAngle(180, AngleDegree).In(AngleRadian), math.Pi, 1e-12)
	almostEqual(t, NewAngle(1, AngleMOA).In(AngleRadian), math.Pi/10800, 1e-12)
}
