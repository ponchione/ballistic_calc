package units

import "testing"

func TestDistanceConversions(t *testing.T) {
	almostEqual(t, NewDistance(1, DistanceFoot).In(DistanceInch), 12, 1e-12)
	almostEqual(t, NewDistance(1, DistanceYard).In(DistanceInch), 36, 1e-12)
	almostEqual(t, NewDistance(1, DistanceMile).In(DistanceInch), 63360, 1e-9)
	almostEqual(t, NewDistance(1, DistanceNauticalMile).In(DistanceInch), 72913.3858, 1e-6)
	almostEqual(t, NewDistance(1, DistanceLine).In(DistanceInch), 0.1, 1e-12)
	almostEqual(t, NewDistance(25.4, DistanceMillimeter).In(DistanceInch), 1, 1e-12)
	almostEqual(t, NewDistance(2.54, DistanceCentimeter).In(DistanceInch), 1, 1e-12)
	almostEqual(t, NewDistance(1, DistanceMeter).In(DistanceInch), 1000.0/25.4, 1e-9)
	almostEqual(t, NewDistance(1, DistanceKilometer).In(DistanceInch), 1000000.0/25.4, 1e-6)
}
