# Epic 1: Measurement Foundation Detailed Plan

> For Hermes: Use subagent-driven-development skill to implement this plan task-by-task.

Goal: Build the full unit-conversion layer from Section 5 so every later package can rely on one shared measurement system.

Architecture: Use one Go package, `units`, with one type per measurement family. Store values in canonical units chosen to match the later solver math: radians, inches, feet-per-second, foot-pounds, inches of mercury, Fahrenheit, and grains. Keep conversions isolated to constructors/getters; no downstream package should hand-roll its own conversion math.

Tech Stack: Go 1.25, standard library only, table-driven tests.

This plan refines Epic 1 from `docs/plans/2026-03-29-spec-epic-breakdown.md` and replaces the earlier single-test-file idea with per-family tests plus one shared round-trip matrix.

---

## Target file layout

Create:
- `units/angle.go`
- `units/angle_test.go`
- `units/distance.go`
- `units/distance_test.go`
- `units/velocity.go`
- `units/velocity_test.go`
- `units/energy.go`
- `units/energy_test.go`
- `units/pressure.go`
- `units/pressure_test.go`
- `units/temperature.go`
- `units/temperature_test.go`
- `units/weight.go`
- `units/weight_test.go`
- `units/roundtrip_test.go`
- `units/test_helpers_test.go`

Shared design rules:
- `Angle` stores radians.
- `Distance` stores inches.
- `Velocity` stores feet per second.
- `Energy` stores foot-pounds.
- `Pressure` stores inches of mercury.
- `Temperature` stores Fahrenheit.
- `Weight` stores grains.
- Constructors follow the pattern `NewAngle(value float64, unit AngleUnit) Angle`.
- Getters follow the pattern `func (a Angle) In(unit AngleUnit) float64`.
- Unknown unit constants may panic; they are programmer errors, not user-input parsing.

---

### Task 1: Create the package skeleton and first angle conversions

Objective: establish the `units` package shape with radians, degrees, and MOA working end-to-end.

Files:
- Create `units/angle.go`
- Create `units/angle_test.go`
- Create `units/test_helpers_test.go`

Step 1: Write failing tests

```go
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
```

```go
package units

import (
    "math"
    "testing"
)

func almostEqual(t *testing.T, got, want, tol float64) {
    t.Helper()
    if math.Abs(got-want) > tol {
        t.Fatalf("got %v want %v tol %v", got, want, tol)
    }
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestAngleRadDegMOA -v`

Expected:
FAIL — package or symbols do not exist yet.

Step 3: Write minimal implementation

```go
package units

import "math"

type AngleUnit int

const (
    AngleRadian AngleUnit = iota
    AngleDegree
    AngleMOA
    AngleMil
    AngleMilliradian
    AngleThousand
    AngleInPer100Yard
    AngleCmPer100Meter
)

type Angle struct {
    radians float64
}

func NewAngle(value float64, unit AngleUnit) Angle {
    switch unit {
    case AngleRadian:
        return Angle{radians: value}
    case AngleDegree:
        return Angle{radians: value / 180 * math.Pi}
    case AngleMOA:
        return Angle{radians: value / 180 * math.Pi / 60}
    default:
        panic("unknown angle unit")
    }
}

func (a Angle) In(unit AngleUnit) float64 {
    switch unit {
    case AngleRadian:
        return a.radians
    case AngleDegree:
        return a.radians * 180 / math.Pi
    case AngleMOA:
        return a.radians * 180 / math.Pi * 60
    default:
        panic("unknown angle unit")
    }
}
```

Step 4: Run test to verify pass

Run:
`go test ./units -run TestAngleRadDegMOA -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/angle.go units/angle_test.go units/test_helpers_test.go
git commit -m "feat: add base angle conversions"
```

---

### Task 2: Add mil, milliradian, and thousand angle units

Objective: complete the remaining pure-angle units before introducing the geometric units.

Files:
- Modify `units/angle.go`
- Modify `units/angle_test.go`

Step 1: Write failing test

```go
func TestAngleMilMradThousand(t *testing.T) {
    almostEqual(t, NewAngle(1, AngleMil).In(AngleRadian), math.Pi/3200, 1e-12)
    almostEqual(t, NewAngle(1, AngleMilliradian).In(AngleRadian), 0.001, 1e-12)
    almostEqual(t, NewAngle(1, AngleThousand).In(AngleRadian), math.Pi/3000, 1e-12)
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestAngleMilMradThousand -v`

Expected:
FAIL — unknown angle unit panic or missing conversion logic.

Step 3: Write minimal implementation

Add these branches to `NewAngle` and `In`:

```go
case AngleMil:
    return Angle{radians: value / 3200 * math.Pi}
case AngleMilliradian:
    return Angle{radians: value / 1000}
case AngleThousand:
    return Angle{radians: value / 3000 * math.Pi}
```

And inverse conversions:

```go
case AngleMil:
    return a.radians / math.Pi * 3200
case AngleMilliradian:
    return a.radians * 1000
case AngleThousand:
    return a.radians / math.Pi * 3000
```

Step 4: Run test to verify pass

Run:
`go test ./units -run TestAngleMilMradThousand -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/angle.go units/angle_test.go
git commit -m "feat: add mil mrad and thousand angle units"
```

---

### Task 3: Add geometric angle units for in/100yd and cm/100m

Objective: implement the two non-linear angle units correctly with `atan` and `tan`.

Files:
- Modify `units/angle.go`
- Modify `units/angle_test.go`

Step 1: Write failing tests

```go
func TestAngleLinearPerDistanceUnitsUseGeometry(t *testing.T) {
    almostEqual(t, NewAngle(1, AngleInPer100Yard).In(AngleRadian), math.Atan(1.0/3600.0), 1e-12)
    almostEqual(t, NewAngle(1, AngleCmPer100Meter).In(AngleRadian), math.Atan(1.0/10000.0), 1e-12)
}

func TestAngleLinearPerDistanceRoundTrips(t *testing.T) {
    almostEqual(t, NewAngle(1, AngleInPer100Yard).In(AngleInPer100Yard), 1, 1e-12)
    almostEqual(t, NewAngle(1, AngleCmPer100Meter).In(AngleCmPer100Meter), 1, 1e-12)
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run 'TestAngleLinearPerDistance' -v`

Expected:
FAIL

Step 3: Write minimal implementation

Add these branches:

```go
case AngleInPer100Yard:
    return Angle{radians: math.Atan(value / 3600)}
case AngleCmPer100Meter:
    return Angle{radians: math.Atan(value / 10000)}
```

And inverse conversions:

```go
case AngleInPer100Yard:
    return math.Tan(a.radians) * 3600
case AngleCmPer100Meter:
    return math.Tan(a.radians) * 10000
```

Step 4: Run test to verify pass

Run:
`go test ./units -run 'TestAngleLinearPerDistance' -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/angle.go units/angle_test.go
git commit -m "feat: add geometric linear-angle units"
```

---

### Task 4: Implement distance units

Objective: add the full distance family using inches as the canonical unit.

Files:
- Create `units/distance.go`
- Create `units/distance_test.go`

Step 1: Write failing test

```go
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
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestDistanceConversions -v`

Expected:
FAIL

Step 3: Write minimal implementation

Implement `DistanceUnit`, `Distance`, `NewDistance`, and `In` in `units/distance.go` using the exact Appendix A constants.

Step 4: Run test to verify pass

Run:
`go test ./units -run TestDistanceConversions -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/distance.go units/distance_test.go
git commit -m "feat: add distance conversions"
```

---

### Task 5: Implement velocity units

Objective: add the full velocity family using feet per second as the canonical unit.

Files:
- Create `units/velocity.go`
- Create `units/velocity_test.go`

Step 1: Write failing test

```go
package units

import "testing"

func TestVelocityConversions(t *testing.T) {
    almostEqual(t, NewVelocity(1, VelocityMeterPerSecond).In(VelocityFootPerSecond), 3.2808399, 1e-7)
    almostEqual(t, NewVelocity(3.6, VelocityKilometerPerHour).In(VelocityMeterPerSecond), 1, 1e-12)
    almostEqual(t, NewVelocity(2.23693629, VelocityMilePerHour).In(VelocityMeterPerSecond), 1, 1e-8)
    almostEqual(t, NewVelocity(1.94384449, VelocityKnot).In(VelocityMeterPerSecond), 1, 1e-8)
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestVelocityConversions -v`

Expected:
FAIL

Step 3: Write minimal implementation

Implement `VelocityUnit`, `Velocity`, `NewVelocity`, and `In` using Appendix A constants.

Step 4: Run test to verify pass

Run:
`go test ./units -run TestVelocityConversions -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/velocity.go units/velocity_test.go
git commit -m "feat: add velocity conversions"
```

---

### Task 6: Implement energy units

Objective: add the small energy family cleanly and keep the conversion constant exact.

Files:
- Create `units/energy.go`
- Create `units/energy_test.go`

Step 1: Write failing test

```go
package units

import "testing"

func TestEnergyConversions(t *testing.T) {
    almostEqual(t, NewEnergy(1, EnergyJoule).In(EnergyFootPound), 0.737562149277, 1e-12)
    almostEqual(t, NewEnergy(1, EnergyFootPound).In(EnergyJoule), 1/0.737562149277, 1e-12)
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestEnergyConversions -v`

Expected:
FAIL

Step 3: Write minimal implementation

Implement `EnergyUnit`, `Energy`, `NewEnergy`, and `In`.

Step 4: Run test to verify pass

Run:
`go test ./units -run TestEnergyConversions -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/energy.go units/energy_test.go
git commit -m "feat: add energy conversions"
```

---

### Task 7: Implement pressure units

Objective: add pressure conversions with inches of mercury as the canonical unit because the atmosphere formulas consume it directly.

Files:
- Create `units/pressure.go`
- Create `units/pressure_test.go`

Step 1: Write failing test

```go
package units

import "testing"

func TestPressureConversions(t *testing.T) {
    almostEqual(t, NewPressure(25.4, PressureMillimeterMercury).In(PressureInchMercury), 1, 1e-12)
    almostEqual(t, NewPressure(1, PressureBar).In(PressureMillimeterMercury), 750.061683, 1e-6)
    almostEqual(t, NewPressure(1, PressureHectopascal).In(PressureMillimeterMercury), 750.061683/1000, 1e-9)
    almostEqual(t, NewPressure(1, PressurePSI).In(PressureMillimeterMercury), 51.714924102396, 1e-9)
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestPressureConversions -v`

Expected:
FAIL

Step 3: Write minimal implementation

Implement `PressureUnit`, `Pressure`, `NewPressure`, and `In`.

Step 4: Run test to verify pass

Run:
`go test ./units -run TestPressureConversions -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/pressure.go units/pressure_test.go
git commit -m "feat: add pressure conversions"
```

---

### Task 8: Implement temperature units

Objective: add Fahrenheit, Celsius, Kelvin, and Rankine with Fahrenheit canonical to match the atmosphere formulas.

Files:
- Create `units/temperature.go`
- Create `units/temperature_test.go`

Step 1: Write failing test

```go
package units

import "testing"

func TestTemperatureConversions(t *testing.T) {
    almostEqual(t, NewTemperature(32, TemperatureFahrenheit).In(TemperatureCelsius), 0, 1e-12)
    almostEqual(t, NewTemperature(0, TemperatureCelsius).In(TemperatureFahrenheit), 32, 1e-12)
    almostEqual(t, NewTemperature(273.15, TemperatureKelvin).In(TemperatureFahrenheit), 32, 1e-12)
    almostEqual(t, NewTemperature(491.67, TemperatureRankine).In(TemperatureFahrenheit), 32, 1e-12)
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestTemperatureConversions -v`

Expected:
FAIL

Step 3: Write minimal implementation

Implement `TemperatureUnit`, `Temperature`, `NewTemperature`, and `In`.

Step 4: Run test to verify pass

Run:
`go test ./units -run TestTemperatureConversions -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/temperature.go units/temperature_test.go
git commit -m "feat: add temperature conversions"
```

---

### Task 9: Implement weight units, including the Newton quirk

Objective: add the weight family and explicitly preserve the spec's observed Newton conversion constant.

Files:
- Create `units/weight.go`
- Create `units/weight_test.go`

Step 1: Write failing test

```go
package units

import "testing"

func TestWeightConversions(t *testing.T) {
    almostEqual(t, NewWeight(1, WeightGram).In(WeightGrain), 15.4323584, 1e-9)
    almostEqual(t, NewWeight(1, WeightKilogram).In(WeightGrain), 15432.3584, 1e-6)
    almostEqual(t, NewWeight(1, WeightNewton).In(WeightGrain), 151339.73750336, 1e-6)
    almostEqual(t, NewWeight(1, WeightOunce).In(WeightGrain), 437.5, 1e-12)
    almostEqual(t, NewWeight(1, WeightPound).In(WeightGrain), 7000, 1e-9)
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestWeightConversions -v`

Expected:
FAIL

Step 3: Write minimal implementation

Implement `WeightUnit`, `Weight`, `NewWeight`, and `In` using `1 N = 151339.73750336 gr` exactly.

Step 4: Run test to verify pass

Run:
`go test ./units -run TestWeightConversions -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/weight.go units/weight_test.go
git commit -m "feat: add weight conversions"
```

---

### Task 10: Add the full round-trip matrix for every listed unit

Objective: prove Section 12.9 across every unit, not just spot checks.

Files:
- Create `units/roundtrip_test.go`

Step 1: Write failing test

```go
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
        {name: "angle-in/100yd", got: NewAngle(1.75, AngleInPer100Yard).In(AngleInPer100Yard), want: 1.75},
        {name: "angle-cm/100m", got: NewAngle(3.25, AngleCmPer100Meter).In(AngleCmPer100Meter), want: 3.25},
        {name: "distance-inch", got: NewDistance(12.5, DistanceInch).In(DistanceInch), want: 12.5},
        {name: "distance-foot", got: NewDistance(8.25, DistanceFoot).In(DistanceFoot), want: 8.25},
        {name: "distance-yard", got: NewDistance(9.75, DistanceYard).In(DistanceYard), want: 9.75},
        {name: "distance-mile", got: NewDistance(1.5, DistanceMile).In(DistanceMile), want: 1.5},
        {name: "distance-nautical-mile", got: NewDistance(0.75, DistanceNauticalMile).In(DistanceNauticalMile), want: 0.75},
        {name: "distance-mm", got: NewDistance(15.5, DistanceMillimeter).In(DistanceMillimeter), want: 15.5},
        {name: "distance-cm", got: NewDistance(12.25, DistanceCentimeter).In(DistanceCentimeter), want: 12.25},
        {name: "distance-m", got: NewDistance(99.125, DistanceMeter).In(DistanceMeter), want: 99.125},
        {name: "distance-km", got: NewDistance(1.25, DistanceKilometer).In(DistanceKilometer), want: 1.25},
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
        {name: "weight-gr", got: NewWeight(168, WeightGrain).In(WeightGrain), want: 168},
        {name: "weight-oz", got: NewWeight(2.5, WeightOunce).In(WeightOunce), want: 2.5},
        {name: "weight-g", got: NewWeight(10.8, WeightGram).In(WeightGram), want: 10.8},
        {name: "weight-lb", got: NewWeight(1.25, WeightPound).In(WeightPound), want: 1.25},
        {name: "weight-kg", got: NewWeight(0.85, WeightKilogram).In(WeightKilogram), want: 0.85},
        {name: "weight-n", got: NewWeight(2.0, WeightNewton).In(WeightNewton), want: 2.0},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            almostEqual(t, tc.got, tc.want, 1e-7)
        })
    }
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestRoundTripAllUnits -v`

Expected:
FAIL until all unit families are present.

Step 3: Write minimal implementation

No new production files should be needed if earlier tasks are complete. Fix any missing branches or precision mistakes until this test passes.

Step 4: Run test to verify pass

Run:
`go test ./units -run TestRoundTripAllUnits -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/roundtrip_test.go units/*.go
git commit -m "test: add full unit round-trip matrix"
```

---

### Task 11: Add the explicit angular invariants from Scenario 12.9

Objective: lock down the two clean-room-sensitive angle invariants called out by the spec.

Files:
- Modify `units/angle_test.go`

Step 1: Write failing test

```go
func TestAngularSpecInvariants(t *testing.T) {
    a := NewAngle(1, AngleInPer100Yard)
    almostEqual(t, a.In(AngleMOA), 0.954930, 1e-5)

    got := fmt.Sprintf("%.2fcm/100m", a.In(AngleCmPer100Meter))
    if got != "2.78cm/100m" {
        t.Fatalf("got %q want %q", got, "2.78cm/100m")
    }
}
```

Step 2: Run test to verify failure

Run:
`go test ./units -run TestAngularSpecInvariants -v`

Expected:
FAIL if the geometric conversions are wrong or any constant drifted.

Step 3: Write minimal implementation

If needed, only adjust the angle conversion logic in `units/angle.go`. Do not add special-case formatting code; plain numeric formatting should be enough.

Step 4: Run test to verify pass

Run:
`go test ./units -run TestAngularSpecInvariants -v`

Expected:
PASS

Step 5: Commit

```bash
git add units/angle_test.go units/angle.go
git commit -m "test: lock down angular spec invariants"
```

---

### Task 12: Run the final package test pass and stop

Objective: finish Epic 1 with one clear green checkpoint and no extra polish.

Files:
- No new files expected

Step 1: Run the focused package tests

Run:
`go test ./units -v`

Expected:
PASS

Step 2: Run the repo test suite

Run:
`go test ./...`

Expected:
PASS

Step 3: Commit only if there were final cleanup edits

```bash
git add units
git commit -m "test: finish measurement foundation epic"
```

Stop condition:
- Do not start atmosphere work in the same change.
- Once `units` is green, pause and plan Epic 2 separately.

---

## Done definition for Epic 1

Epic 1 is complete when:
- All seven measurement families exist.
- Every unit listed in Section 5 is supported.
- Round-tripping each unit through its own constructor/getter pair stays within `1e-7`.
- `1 in/100yd` equals `0.954930 moa +/- 1e-5`.
- `1 in/100yd` converts to a formatted `2.78cm/100m`.
- `go test ./units` passes cleanly.

## Anti-scope notes

Do not add in Epic 1:
- String parsing from user input
- JSON/YAML serialization helpers
- Generic unit registries
- Shared numeric abstractions across families
- Public API compatibility shims for the original repo

Keep it boring and exact.
