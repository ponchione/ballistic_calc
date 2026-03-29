package units

import "testing"

func TestRoundTripAllUnits(t *testing.T) {
	type roundTripCase struct {
		name string
		got  float64
		want float64
	}

	cases := []roundTripCase{
		{name: "angle-rad", got: NewAngle(1.2345, AngleRadian).In(AngleRadian), want: 1.2345},
		{name: "angle-deg", got: NewAngle(12.345, AngleDegree).In(AngleDegree), want: 12.345},
		{name: "angle-moa", got: NewAngle(7.89, AngleMOA).In(AngleMOA), want: 7.89},
		{name: "angle-mil", got: NewAngle(5.67, AngleMil).In(AngleMil), want: 5.67},
		{name: "angle-mrad", got: NewAngle(2.5, AngleMilliradian).In(AngleMilliradian), want: 2.5},
		{name: "angle-thousand", got: NewAngle(4.25, AngleThousand).In(AngleThousand), want: 4.25},
		{name: "angle-in-per-100yd", got: NewAngle(1.75, AngleInPer100Yard).In(AngleInPer100Yard), want: 1.75},
		{name: "angle-cm-per-100m", got: NewAngle(3.25, AngleCmPer100Meter).In(AngleCmPer100Meter), want: 3.25},
		{name: "distance-inch", got: NewDistance(12.5, DistanceInch).In(DistanceInch), want: 12.5},
		{name: "distance-foot", got: NewDistance(8.25, DistanceFoot).In(DistanceFoot), want: 8.25},
		{name: "distance-yard", got: NewDistance(9.75, DistanceYard).In(DistanceYard), want: 9.75},
		{name: "distance-mile", got: NewDistance(1.5, DistanceMile).In(DistanceMile), want: 1.5},
		{name: "distance-nautical-mile", got: NewDistance(0.75, DistanceNauticalMile).In(DistanceNauticalMile), want: 0.75},
		{name: "distance-mm", got: NewDistance(15.5, DistanceMillimeter).In(DistanceMillimeter), want: 15.5},
		{name: "distance-cm", got: NewDistance(12.25, DistanceCentimeter).In(DistanceCentimeter), want: 12.25},
		{name: "distance-meter", got: NewDistance(99.125, DistanceMeter).In(DistanceMeter), want: 99.125},
		{name: "distance-kilometer", got: NewDistance(1.25, DistanceKilometer).In(DistanceKilometer), want: 1.25},
		{name: "distance-line", got: NewDistance(6, DistanceLine).In(DistanceLine), want: 6},
		{name: "velocity-mps", got: NewVelocity(250.5, VelocityMeterPerSecond).In(VelocityMeterPerSecond), want: 250.5},
		{name: "velocity-kph", got: NewVelocity(80.25, VelocityKilometerPerHour).In(VelocityKilometerPerHour), want: 80.25},
		{name: "velocity-fps", got: NewVelocity(2750, VelocityFootPerSecond).In(VelocityFootPerSecond), want: 2750},
		{name: "velocity-mph", got: NewVelocity(12.5, VelocityMilePerHour).In(VelocityMilePerHour), want: 12.5},
		{name: "velocity-knot", got: NewVelocity(18.75, VelocityKnot).In(VelocityKnot), want: 18.75},
		{name: "energy-ftlb", got: NewEnergy(1250, EnergyFootPound).In(EnergyFootPound), want: 1250},
		{name: "energy-joule", got: NewEnergy(850, EnergyJoule).In(EnergyJoule), want: 850},
		{name: "pressure-mmhg", got: NewPressure(760, PressureMillimeterMercury).In(PressureMillimeterMercury), want: 760},
		{name: "pressure-inhg", got: NewPressure(29.92, PressureInchMercury).In(PressureInchMercury), want: 29.92},
		{name: "pressure-bar", got: NewPressure(1.01325, PressureBar).In(PressureBar), want: 1.01325},
		{name: "pressure-hpa", got: NewPressure(1013.25, PressureHectopascal).In(PressureHectopascal), want: 1013.25},
		{name: "pressure-psi", got: NewPressure(14.6959, PressurePSI).In(PressurePSI), want: 14.6959},
		{name: "temperature-f", got: NewTemperature(59, TemperatureFahrenheit).In(TemperatureFahrenheit), want: 59},
		{name: "temperature-c", got: NewTemperature(15, TemperatureCelsius).In(TemperatureCelsius), want: 15},
		{name: "temperature-k", got: NewTemperature(288.15, TemperatureKelvin).In(TemperatureKelvin), want: 288.15},
		{name: "temperature-r", got: NewTemperature(518.67, TemperatureRankine).In(TemperatureRankine), want: 518.67},
		{name: "weight-grain", got: NewWeight(168, WeightGrain).In(WeightGrain), want: 168},
		{name: "weight-ounce", got: NewWeight(2.5, WeightOunce).In(WeightOunce), want: 2.5},
		{name: "weight-gram", got: NewWeight(10.8, WeightGram).In(WeightGram), want: 10.8},
		{name: "weight-pound", got: NewWeight(1.25, WeightPound).In(WeightPound), want: 1.25},
		{name: "weight-kilogram", got: NewWeight(0.85, WeightKilogram).In(WeightKilogram), want: 0.85},
		{name: "weight-newton", got: NewWeight(2.0, WeightNewton).In(WeightNewton), want: 2.0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			almostEqual(t, tc.got, tc.want, 1e-7)
		})
	}
}
