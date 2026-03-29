package units

import "testing"

func TestVelocityConversions(t *testing.T) {
	almostEqual(t, NewVelocity(1, VelocityMeterPerSecond).In(VelocityFootPerSecond), 3.2808399, 1e-7)
	almostEqual(t, NewVelocity(3.6, VelocityKilometerPerHour).In(VelocityMeterPerSecond), 1, 1e-12)
	almostEqual(t, NewVelocity(2.23693629, VelocityMilePerHour).In(VelocityMeterPerSecond), 1, 1e-8)
	almostEqual(t, NewVelocity(1.94384449, VelocityKnot).In(VelocityMeterPerSecond), 1, 1e-8)
}
