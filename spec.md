# Clean-Room Reconstruction Spec: External Ballistics Library

## 1. Purpose

This document defines the behavior required to recreate the current repository as an independent implementation.

The target is a working ballistic-calculation library plus verification assets. It does **not** need to preserve the original Go API, file names, type names, or internal structure. It **does** need to reproduce the same user-visible capabilities and the same test-visible numerical behavior within stated tolerances.

This is a behavioral reconstruction contract, not a source-transcription guide.

## 2. Reconstruction Target

The rebuilt system shall provide:

- A 3-axis ballistic trajectory engine for small-arms and similar projectiles.
- A zero solver that computes the sight angle needed to hit a stated zero distance.
- Support for standard drag models `G1`, `G2`, `G5`, `G6`, `G7`, `G8`, `GS`, and `RA4`.
- Support for custom drag functions and custom drag curves.
- Support for wind, inclined/declined shots, and optional spin drift.
- Unit-aware inputs and outputs across the measurement families listed below.
- Verification assets that prove the rebuilt system matches the normative scenarios in this document.

The rebuilt system does **not** need to preserve:

- Original package paths.
- Original public symbol names.
- Original file/module decomposition.
- Original string formatting for types, except where a verification asset intentionally checks it.

## 3. Clean-Room Boundary

The rebuilt implementation should be written from this specification and independently sourced domain knowledge.

- Use public/reference drag-table data for standard drag tables rather than transcribing table literals from the original repository.
- Preserve observed behavior and verified numerical results, not the original code structure.
- If a more robust numerical method is chosen, it is acceptable as long as the acceptance contract is met.

## 4. Domain Model

The rebuilt system should model the following domain concepts.

### 4.1 Projectile

A projectile has:

- A drag description.
- A projectile weight.
- Optional diameter.
- Optional length.

Observed data requirements at the current repository head:

- Form-factor-to-BC conversion uses projectile weight and diameter.
- Spin-drift calculation additionally uses projectile length and weapon twist.

### 4.2 Ammunition

A round of ammunition has:

- A projectile.
- A muzzle velocity.

### 4.3 Atmosphere

An atmosphere has:

- Altitude above sea level.
- Pressure.
- Temperature.
- Relative humidity.

The model shall expose, either directly or internally:

- Local air-density factor relative to standard density.
- Local speed of sound.

### 4.4 Weapon And Zeroing Data

A weapon has:

- Sight height above bore.
- Zeroing distance.
- Optional rifling twist direction and twist pitch.
- Optional click-value metadata.

Observed behavior at the current repository head:

- Zero distance affects calculations.
- Optional alternate zero ammunition is stored but not consumed by the solver.
- Optional zero atmosphere is stored but not consumed by the solver.
- Click value is stored metadata and is not consumed by the solver.

A clean reimplementation may omit those passive metadata fields unless broader compatibility is desired.

### 4.5 Shot Definition

A shot definition has:

- Sight angle.
- Maximum range to report.
- Report step.
- Optional shot angle relative to the horizon.
- Optional cant angle.

If shot angle and cant angle are omitted, both default to zero.

### 4.6 Wind

Wind input is a piecewise-constant list of segments. Each segment has:

- An `until distance`.
- A wind speed.
- A wind direction.

Wind direction semantics:

- `0 deg`: blowing from the rear / tailwind.
- `+90 deg`: blowing from the left.
- `-90 deg` or `270 deg`: blowing from the right.
- `180 deg`: blowing into the shooter / headwind.

Segments are ordered from nearest to farthest.

Equivalent "no wind" representations should be accepted, including:

- No segments.
- A single zero-valued segment.

## 5. Measurement Support

The rebuilt system shall support the following measurement families.

| Family | Required units |
| --- | --- |
| Angle | radian, degree, MOA, mil (1/6400 circle), milliradian, thousand (1/6000 circle), inches per 100 yards, centimeters per 100 meters |
| Distance | inch, foot, yard, mile, nautical mile, millimeter, centimeter, meter, kilometer, line (0.1 inch) |
| Velocity | m/s, km/h, ft/s, mph, knot |
| Energy | foot-pound, joule |
| Pressure | mmHg, inHg, bar, hPa, psi |
| Temperature | Fahrenheit, Celsius, Kelvin, Rankine |
| Weight | grain, ounce, gram, pound, kilogram, newton |

The implementation may use any internal canonical units, but it must preserve the following behavior:

- Converting a value expressed in any supported unit and reading it back in that same unit should round-trip within `1e-7` for ordinary values.
- The angular "linear per distance" units are geometric, not linearized approximations:
  - `in/100yd` means `atan(value / 3600)` radians.
  - `cm/100m` means `atan(value / 10000)` radians.

Appendix A lists the exact conversion constants observed at the current repository head.

## 6. Atmosphere Model

### 6.1 Default Atmosphere

The default atmosphere used by the current repository is:

- Altitude: `0 ft`
- Pressure: `29.92 inHg`
- Temperature: `59 F`
- Relative humidity: `0.78`

This humidity value is material because the normative trajectory scenarios in this document use the default atmosphere.

### 6.2 Explicit Atmosphere Construction

An explicit atmosphere constructor shall accept:

- Altitude
- Pressure
- Temperature
- Humidity

Humidity rules:

- Values in `[0, 1]` are treated as a fraction.
- Values in `(1, 100]` are treated as percents and divided by `100`.
- Values outside those ranges are rejected.

Observed compatibility notes:

- At the current repository head, humidity is the only numeric input that is range-checked. Altitude, pressure, and temperature are otherwise accepted as provided.
- On humidity validation failure, the Go constructor returns an error together with a fully initialized default atmosphere. That fallback is an observed HEAD quirk, not required rebuild behavior.

### 6.3 ICAO Atmosphere By Altitude

The rebuilt system should also support an ICAO-style atmosphere defined from altitude alone:

- Humidity is `0`.
- Temperature and pressure vary with altitude using the formulas below.

Observed compatible formulas:

```text
lapse = -3.56616e-03
exponent = -5.255876

temperature_F = 518.67 + altitude_ft * lapse - 459.67
pressure_inHg = 29.92 * (518.67 / (temperature_F + 459.67)) ^ exponent
humidity = 0
```

### 6.4 Density And Speed-Of-Sound Formulas

For a temperature `T_F` in Fahrenheit, pressure `P_inHg` in inches of mercury, and normalized humidity `RH`:

```text
if T_F > 0:
    et0 = 1.24871
        + 0.0988438 * T_F
        + 0.00152907 * T_F^2
        - 3.07031e-06 * T_F^3
        + 4.21329e-07 * T_F^4
    et = 3.342e-04 * RH * et0
    hc = (P_inHg - 0.3783 * et) / 29.92
else:
    hc = 1

density_lb_per_ft3 = 0.076474 * (518.67 / (T_F + 459.67)) * hc
mach_fps = sqrt(T_F + 459.67) * 49.0223
density_factor = density_lb_per_ft3 / 0.076474
```

Notes:

- The original repository disables humidity correction at and below `0 F`. This is an observed behavior, not necessarily a physics requirement.
- A rebuild may use a more complete humidity treatment if it does not regress the normative scenarios.

### 6.5 Altitude-Adjusted Local Conditions During Flight

The current repository varies density factor and local speed of sound with projectile altitude during trajectory integration.

At projectile altitude `alt_ft`, starting from a base atmosphere defined at `base_alt_ft`, `base_T_F`, `base_P_inHg`:

```text
if abs(base_alt_ft - alt_ft) < 30:
    reuse base density and base mach
else:
    lapse = -3.56616e-03
    exponent = -5.255876

    ta = 518.67 + base_alt_ft * lapse - 459.67
    tb = 518.67 + alt_ft * lapse - 459.67
    T_alt_F = base_T_F - ta + tb
    P_alt_inHg = base_P_inHg * ((base_T_F + 459.67) / (T_alt_F + 459.67)) ^ exponent

    recompute density and mach from T_alt_F, P_alt_inHg, RH
```

The current repository refreshes those values only when altitude changes by more than `1 meter` (`3.28084 ft`) since the last refresh.

## 7. Drag Model

### 7.1 Supported Drag Modes

The rebuilt system shall support:

- Standard drag table mode with table identifiers `G1`, `G2`, `G5`, `G6`, `G7`, `G8`, `GS`, and `RA4`.
- Custom drag-function mode, where the caller supplies an arbitrary function `f(mach) -> raw_drag`.

The rebuilt system shall also provide a way to use tabulated `(Mach, drag)` points with custom drag-function mode. Observed behavior at the current repository head is a two-step helper path: one helper precomputes interpolation coefficients from the tabulated points, and a separate evaluator consumes those coefficients together with the original table to produce raw drag at a requested Mach. A caller-supplied wrapper around that evaluator is then passed to the custom drag-function constructor. See Section 7.4 for the interpolation behavior.

### 7.2 Ballistic-Coefficient Value Modes

The drag description supports two value modes.

`BC` mode:

- The supplied value is used directly as the ballistic coefficient.

`FF` mode:

- The supplied value is a form factor.
- A compatible implementation must derive an equivalent ballistic coefficient from projectile properties:

```text
equivalent_bc = (weight_gr / 7000) / (diameter_in^2 * form_factor)
```

Because of that formula, form-factor mode requires projectile weight and a positive projectile diameter.

Projectile length is not used in the form-factor-to-BC conversion.

The current repository does not explicitly validate that projectile diameter is present and nonzero before performing the FF calculation. If dimensions are omitted, or if diameter is explicitly zero, the equivalent BC calculation can overflow to `+Inf`. A clean-room rebuild shall reject FF-mode inputs whose effective projectile diameter is missing or nonpositive rather than relying on that behavior.

Validation requirements for the rebuild:

- BC and form-factor values must be strictly greater than zero.
- External standard drag-table identifiers must be one of `G1`, `G2`, `G5`, `G6`, `G7`, `G8`, `GS`, or `RA4`; any other identifier must be rejected without panic.
- Custom-drag value types must be either `BC` or `FF`.

Standard drag-table definitions always use BC mode. FF mode is only available for custom drag-function definitions.

Observed compatibility note:

- The current repository has an internal custom-table sentinel that can slip through the standard-table constructor's numeric range check and then panic immediately during construction when drag-function lookup runs. That is a HEAD bug, not required rebuild behavior.

### 7.3 Drag Scaling

After retrieving a drag value from a standard table or custom curve/function, scale it by:

```text
PIR = 2.08551e-04
scaled_drag = raw_drag * PIR
```

### 7.4 Custom Curve Interpolation

The current repository derives a local polynomial curve from tabulated points:

- First point: linear interpolation from the first two data points.
- Interior points: quadratic interpolation through three adjacent points.
- Final point: constant tail using the last drag value.

At evaluation time:

- Find the two neighboring data points around the requested Mach.
- Choose whichever tabulated point is closer in Mach.
- Evaluate that point's precomputed polynomial.

A rebuild may use an equivalent interpolation strategy if it reproduces the normative scenarios. If off-table behavior must match the current repository, use:

- Linear extrapolation below the minimum Mach.
- Constant tail beyond the maximum Mach.

Construction path at the current repository head: the curve-precomputation utility accepts tabulated data points and returns interpolation coefficients, not a callable drag function. A separate evaluator consumes the original table, those precomputed coefficients, and a Mach value to produce raw drag. In the shipped tests, a small caller-defined wrapper around that evaluator is what gets passed to the custom drag-function constructor together with a BC or FF value. PIR scaling from Section 7.3 is applied by the drag-function evaluation path, not by the curve-precomputation utility itself. The curve helpers operate on raw (unscaled) drag values.

Observed input assumptions: the current repository expects at least two tabulated points with strictly increasing Mach values. Fewer than two points panic during curve precomputation. Duplicate Mach values can produce divide-by-zero coefficients and `Inf`/`NaN` evaluation results. Descending or otherwise out-of-order Mach values do not have a defined compatibility contract. A clean-room rebuild should reject fewer-than-two, duplicate-Mach, and out-of-order inputs explicitly rather than inheriting that behavior.

## 8. Zero Solver

The rebuilt system shall provide a zero solver that computes the sight angle required for a weapon/ammunition/atmosphere combination at a stated zero distance.

The zero solver shall account for:

- Sight height over bore.
- Gravity.
- Drag.
- Atmospheric density factor.
- Local speed of sound.

Observed behavior at the current repository head:

- Zero solving uses the atmosphere at the muzzle and does not refresh with altitude.
- The solver starts with zero barrel elevation and iteratively corrects elevation from the vertical miss at the zero distance.
- The current repository compares the vertical miss magnitude against `5e-6 ft` and stops after at most `10` iterations.

A compatible solver may use any numerical method that meets the acceptance scenarios below.

### 8.1 Compatible Zero-Solver Requirements

A compatible zero solver may use any numerical method, but it shall preserve these observed behaviors:

- Use the muzzle atmosphere's density factor and local speed of sound for the whole solve; do not refresh with projectile altitude.
- Start from the muzzle state `position = (0, -sight_height_ft, 0)` with zero barrel azimuth.
- Solve for the barrel elevation required to bring the predicted projectile height to the zero distance.
- If reproducing HEAD's convergence behavior, use a reduced internal step derived from `10` units in the zero-distance unit family, update barrel elevation from the observed miss at the zero distance, compare the miss magnitude against `5e-6 ft`, and stop after at most `10` iterations.
- The normative zero-solver outputs are the acceptance scenarios in Sections 12.1 and 12.2; those outputs take precedence over any particular implementation recipe.

## 9. Trajectory Engine

### 9.1 Required Inputs

The trajectory engine shall accept:

- Ammunition
- Weapon
- Atmosphere
- Shot definition
- Optional wind segments

### 9.2 Reported Sample Ranges

For ordinary shots, define report thresholds at:

- `0`
- `step`
- `2 * step`
- ...
- `floor(maximum_distance / step) * step`

If `maximum distance` is an exact multiple of `step`, the final threshold equals `maximum distance`.

Observed sampling behavior at the current repository head:

- A sample is emitted on the first computed state whose travelled distance is greater than or equal to the next threshold.
- The current repository does **not** interpolate samples back to the exact threshold.
- Reported `travel distance` and `flat distance` values can therefore overshoot the nominal threshold slightly.

For normal trajectories that reach the requested range, the observed sample count is:

```text
floor(maximum_distance / step) + 1
```

If a projectile becomes unusable earlier, the current repository still returns the preallocated full-length slice and leaves any unfilled tail entries at zero values. A compatible rebuild shall preserve that output shape.

### 9.3 Internal Integration Step

The current repository uses a smaller internal integration step than the reporting step.

Observed reducer:

```text
step_ft = report_step_ft / 2
maximum_step_ft = 1.0 by default

if step_ft > maximum_step_ft:
    step_order = floor(log10(step_ft))
    max_order = floor(log10(maximum_step_ft))
    step_ft = step_ft / 10^(step_order - max_order + 1)
```

This reducer materially affects the observed outputs. A rebuild may use a different integrator if it still matches the normative scenarios.

### 9.4 State And Axes

Use a right-handed motion model with:

- `x`: downrange toward the target
- `y`: vertical
- `z`: cross-range / windage

The projectile starts at:

```text
position = (0, -sight_height_ft, 0)
```

Gravity is:

```text
gravity = (0, -32.17405, 0) ft/s^2
```

### 9.5 Barrel Elevation And Line Of Sight

Let:

- `sight_angle` be the angle between optic line and bore line.
- `shot_angle` be the incline/decline angle relative to the horizon.

Then:

- Line of sight uses `shot_angle`.
- Barrel elevation uses `sight_angle + shot_angle` for inclined/declined shots.
- Barrel azimuth is effectively zero in the current repository.

### 9.6 Wind Vector

The current repository converts a wind segment into a 3D vector using:

```text
combined_angle = sight_angle
if shot_angle != 0:
    combined_angle += shot_angle

sight_cos = cos(combined_angle)
sight_sin = sin(combined_angle)
cant_cos = cos(cant_angle)
cant_sin = sin(cant_angle)

range_velocity = wind_speed_fps * cos(direction_rad)
cross_component = wind_speed_fps * sin(direction_rad)
range_factor = -range_velocity * sight_sin

wind_x = range_velocity * sight_cos
wind_y = range_factor * cant_cos + cross_component * cant_sin
wind_z = cross_component * cant_cos - range_factor * cant_sin
```

Observed wind-segment switching behavior at the current repository head:

- Each segment's `until distance` acts as the handoff boundary to the next segment.
- The active segment switches when travelled slant distance first reaches or exceeds that boundary.
- Once the final segment becomes active, it remains active for the rest of the shot; the last segment's own `until distance` is not used as a shutoff point.

Observed limitation at the current repository head:

- Cant angle rotates wind decomposition.
- Cant angle does not rotate gravity or the drop reference frame.

If nonzero cant support is advertised, preserving that limitation is the HEAD-compatible behavior. A rebuild may adopt a more complete cant model only if its verification assets intentionally choose and verify that compatibility break.

### 9.7 Spin Drift

If and only if both of the following are present:

- Weapon twist information
- Projectile diameter and length

the engine shall compute a stability coefficient and add a spin-drift term to windage.

Observed formulas:

```text
twist_calibers = twist_in / diameter_in
length_calibers = length_in / diameter_in

stability =
    30 * weight_gr
    / (twist_calibers^2 * diameter_in^3 * length_calibers * (1 + length_calibers^2))

velocity_factor = (muzzle_velocity_fps / 2800)^(1/3)
atmosphere_factor = ((temperature_F + 460) / (59 + 460)) * (29.92 / pressure_inHg)

stability = stability * velocity_factor * atmosphere_factor
```

Windage drift contribution:

```text
drift_ft = 1.25 * (stability + 1.2) * time_s^1.83 * twist_sign / 12
```

Twist sign:

- Left twist: `+1`
- Right twist: `-1`

### 9.8 Compatible Integration Requirements

A compatible trajectory integrator may use any numeric scheme, but it shall preserve these observed behaviors:

- Start from `position = (0, -sight_height_ft, 0)` with the launch vector defined by barrel elevation and zero barrel azimuth.
- Use a reduced internal step smaller than the report step, or another scheme that still reproduces the normative scenarios.
- Stop integrating when projectile speed falls below `50 fps` or raw vertical position drops below `-15000 ft`.
- Refresh local density factor and local speed of sound only when projectile altitude has changed by more than `3.28084 ft` since the last refresh.
- Switch wind segments when travelled slant distance first reaches or exceeds the next segment boundary.
- Emit a report sample on the first computed state whose travelled slant distance reaches or exceeds the next report threshold; do not back-interpolate to the exact threshold.
- Apply wind against air-relative velocity.
- Compute travelled slant distance from horizontal position and shot angle, and accumulate time from the integrated motion rather than from nominal report thresholds.
- The normative trajectory outputs are the acceptance scenarios in Section 12; those outputs take precedence over any particular loop structure.

### 9.9 Emitted Sample Fields

Each trajectory sample shall provide, conceptually, at least the following values.

`time`

- Total time since launch, in seconds.

`travel distance`

- Slant range along the line of sight.
- For uninclined shots, this equals flat downrange distance.
- Because samples are emitted on threshold crossing rather than by interpolation, the stored value may be slightly greater than the nominal reporting threshold.

`flat distance`

- Horizontal/downrange `x` distance.

`velocity`

- Projectile speed magnitude.

`mach`

- `velocity / local_speed_of_sound`.

`drop`

- Distance from the current projectile position to the line of sight.
- Positive means above the line of sight.
- Negative means below the line of sight.

`drop flat`

- Raw vertical (Y) position of the projectile, where Y=0 is the sight centerline. At the muzzle, this equals negative sight height (e.g., `-2.50 in` for a `2.5 in` sight height). For level shots, `drop flat` equals `drop`. For inclined shots, `drop flat` is the unrotated Y coordinate while `drop` is rotated into the line-of-sight frame per Section 9.10.

`drop adjustment`

- `atan(drop / reference_distance)`.
- Use flat distance for level shots.
- Use slant distance for inclined/declined shots.
- Adjustment at zero distance is implementation-defined.

`windage`

- Cross-range displacement, including spin drift if active.

`windage adjustment`

- `atan(windage / reference_distance)`.
- Use flat distance for level shots.
- Use slant distance for inclined/declined shots.
- Adjustment at zero distance is implementation-defined.

`line of sight elevation`

- `x * tan(shot_angle)`

`line of departure elevation`

- `x * tan(barrel_elevation) - sight_height`

`energy`

- Foot-pounds computed from:

```text
energy_ft_lb = weight_gr * velocity_fps^2 / 450400
```

`optimal game weight`

- Pounds computed from:

```text
ogw_lb = weight_gr^2 * velocity_fps^3 * 1.5e-12
```

### 9.10 Drop Computation For Inclined Shots

For nonzero shot angle, the current repository rotates the vertical frame relative to the line of sight:

```text
y = raw_y + sight_height
y_rot = -x * sin(shot_angle) + y * cos(shot_angle)
drop = y_rot - sight_height
```

For level shots:

```text
drop = raw_y
```

## 10. Required Behavioral Decisions

These choices should be treated as normative because they materially affect observed behavior.

- The default atmosphere is `59 F`, `29.92 inHg`, `0 ft`, `RH=0.78`.
- Standard drag tables include `RA4` in addition to the common `G*` tables.
- Standard-drag and custom-curve drag values are scaled by `2.08551e-04`.
- Custom drag-function values are also scaled by `2.08551e-04` before use in the solver.
- Wind is applied against air-relative velocity, not directly against ground velocity.
- Wind-direction semantics are the ones implied by Section 9.6's vector math: `0 deg` is tailwind and `180 deg` is headwind.
- Multi-segment wind uses each segment's `until distance` as the handoff boundary to the next segment, and the final segment persists for the remainder of the shot.
- Spin drift only applies when both twist info and projectile dimensions are available.
- Right twist produces negative drift, left twist positive drift.
- Inclined shots report slant range separately from flat distance.
- Atmosphere updates are altitude-sensitive during trajectory integration.
- Cant rotates wind decomposition only; gravity and the drop reference frame are not rotated.
- Early-terminated trajectories retain the full preallocated sample count and leave the unfilled tail at zero values.
- External standard-table identifiers outside the advertised catalog are rejected without panic.

## 11. Non-Normative Or Under-Specified Behavior

These behaviors exist at the current repository head but should not be fossilized unless they matter for acceptance.

- Passive zero metadata fields for alternate ammunition and zero atmosphere.
- Passive click-value metadata.
- Exact helper-package APIs for vectors and unit wrappers.
- The original repository's zero-distance adjustment value at range `0`.
- The helper vector package contains behavior that is not material to the ballistic contract.
- The low-temperature humidity shortcut may be replaced with a more physical model if acceptance scenarios still pass.

## 12. Normative Acceptance Scenarios

Unless otherwise stated:

- Use the default atmosphere from Section 6.1.
- Distances, drops, and windage are measured in the units shown in each table.
- In sample tables, the `Range` column names the reporting threshold. The emitted sample's stored travelled distance may overshoot that threshold slightly because the current repository does not interpolate samples back to exact boundaries.
- `n/a` means the original test suite did not make that value normative at that point.

### 12.1 Zero Solver Scenario A

Inputs:

- Drag model: `G1`
- BC: `0.365`
- Weight: `69 gr`
- Muzzle velocity: `2600 fps`
- Zero distance: `100 yd`
- Sight height: `3.2 in`

Expected result:

- Sight angle: `0.001651 rad +/- 1e-6`

### 12.2 Zero Solver Scenario B

Inputs:

- Drag model: `G7`
- BC: `0.223`
- Weight: `168 gr`
- Muzzle velocity: `2750 fps`
- Zero distance: `100 yd`
- Sight height: `2 in`

Expected result:

- Sight angle: `0.001228 rad +/- 1e-6`

### 12.2a Zero Solver Scenario C

Inputs:

- Drag model: `G1`
- BC: `0.223`
- Weight: `168 gr`
- Muzzle velocity: `2750 fps`
- Zero distance: `100 yd`
- Sight height: `2 in`

Expected result:

- Sight angle: `0.001265 rad +/- 1e-6`

### 12.2b Zero Solver Scenario D

Inputs:

- Drag model: `G1`
- BC: `0.365`
- Weight: `65 gr`
- Muzzle velocity: `2600 fps`
- Zero distance: `100 yd`
- Sight height: `2.5 in`

Expected result:

- Sight angle: `0.001457 rad +/- 1e-6`

### 12.2c Zero Solver Scenario E

Inputs:

- Drag model: `RA4`
- BC: `0.115`
- Weight: `40 gr`
- Muzzle velocity: `2000 fps`
- Zero distance: `100 yd`
- Sight height: `2.5 in`

Expected result:

- Sight angle: `0.002180 rad +/- 1e-6`

### 12.3 G1 Trajectory With Constant Wind

Inputs:

- Drag model: `G1`
- BC: `0.223`
- Weight: `168 gr`
- Muzzle velocity: `2750 fps`
- Zero distance: `100 yd`
- Sight height: `2 in`
- Sight angle: `4.221 moa` (equivalent to `0.001228 rad`)
- Max range: `1000 yd`
- Report step: `100 yd`
- Wind: `5 mph` at `-45 deg`

> **Note:** This is a legacy fixed fixture value reused by the current Go tests. It is not the G1/0.223 zero-solver output, so a clean-room rebuild should treat it as an explicit scenario input rather than re-solving it.

Expected sample count:

- `11`

Expected samples:

| Range yd | Velocity fps | Mach | Energy ft-lb | Drop in | Hold moa | Windage in | Wind adj moa | Time s | OGW lb |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| 0 | 2750.0 | 2.463 | 2820.6 | -2.0 | n/a | 0.0 | n/a | 0.000 | 880 |
| 100 | 2351.2 | 2.106 | 2061.0 | 0.0 | 0.0 | -0.6 | -0.6 | 0.118 | 550 |
| 500 | 1169.1 | 1.047 | 509.8 | -87.9 | -16.8 | -19.5 | -3.7 | 0.857 | 67 |
| 1000 | 776.4 | 0.695 | 224.9 | -823.9 | -78.7 | -87.5 | -8.4 | 2.495 | 20 |

Tolerances:

- Velocity: `+/- 5 fps`
- Mach: `+/- 0.005`
- Energy: `+/- 5 ft-lb`
- Time: `+/- 0.06 s`
- OGW: `+/- 1 lb`
- Drop:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 2 in` from `500 yd` through `799 yd`
  - `+/- 4 in` at `800 yd` and beyond
- Windage:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 1 in` from `500 yd` through `799 yd`
  - `+/- 1.5 in` at `800 yd` and beyond
- Angular adjustments at nonzero range: `+/- 0.5` in the stated angular unit

### 12.4 G7 Trajectory With Wind And Spin Drift

Inputs:

- Drag model: `G7`
- BC: `0.223`
- Diameter: `0.308 in`
- Length: `1.282 in`
- Weight: `168 gr`
- Muzzle velocity: `2750 fps`
- Zero distance: `100 yd`
- Sight height: `2 in`
- Twist: right hand, `11.24 in`
- Sight angle: `4.221 moa` (equivalent to `0.001228 rad` from zero-solver scenario 12.2)
- Max range: `1000 yd`
- Report step: `100 yd`
- Wind: `5 mph` at `-45 deg`

Expected sample count:

- `11`

Expected samples:

| Range yd | Velocity fps | Mach | Energy ft-lb | Drop in | Hold mil | Windage in | Wind adj mil | Time s | OGW lb |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| 0 | 2750.0 | 2.463 | 2820.6 | -2.0 | n/a | 0.0 | n/a | 0.000 | 880 |
| 100 | 2544.3 | 2.279 | 2416.0 | 0.0 | 0.0 | -0.35 | -0.09 | 0.113 | 698 |
| 500 | 1810.7 | 1.622 | 1226.0 | -56.3 | -3.18 | -9.96 | -0.55 | 0.673 | 252 |
| 1000 | 1081.3 | 0.968 | 442.0 | -401.6 | -11.32 | -50.98 | -1.44 | 1.748 | 55 |

Tolerances:

- Velocity: `+/- 5 fps`
- Mach: `+/- 0.005`
- Energy: `+/- 5 ft-lb`
- Time: `+/- 0.06 s`
- OGW: `+/- 1 lb`
- Drop:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 2 in` from `500 yd` through `799 yd`
  - `+/- 4 in` at `800 yd` and beyond
- Windage:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 1 in` from `500 yd` through `799 yd`
  - `+/- 1.5 in` at `800 yd` and beyond
- Angular adjustments at nonzero range: `+/- 0.5` in the stated angular unit

### 12.5 Form-Factor Conversion Scenario

Inputs:

- Drag mode: custom drag function
- Value type: form factor
- Form factor: `1.184`
- Diameter: `0.204 in`
- Length: `1.0 in`
- Weight: `40 gr`

Expected results:

- Stored form factor: `1.184 +/- 0.0005`
- Equivalent BC: `0.116 +/- 0.0005`

### 12.6 Custom-Curve Metric Trajectory

Custom curve points:

| Mach | Drag |
| --- | ---: |
| 0.00 | 0.119 |
| 0.70 | 0.119 |
| 0.85 | 0.120 |
| 0.87 | 0.122 |
| 0.90 | 0.126 |
| 0.93 | 0.148 |
| 0.95 | 0.182 |

Inputs:

- Drag mode: custom drag function (built by precomputing interpolation coefficients from the above curve, then wrapping the corresponding curve evaluator as described in Section 7.4)
- Value type: form factor
- Form factor: `1.0`
- Diameter: `119.56 mm`
- Length: `20 in`
- Weight: `13585 g`
- Muzzle velocity: `555 m/s`
- Zero distance: `100 m`
- Sight height: `40 mm`
- Sight angle: solved from the zero solver for the same ammo/weapon/atmosphere
- Max range: `1500 m`
- Report step: `100 m`

Expected samples:

| Range m | Drop cm | Velocity m/s | Time s |
| --- | ---: | ---: | ---: |
| 100 | 0.0 | 550 | 0.182 |
| 200 | -28.4 | 544 | 0.364 |
| 1500 | -3627.8 | 486 | 2.892 |

Tolerances:

- Drop: `0.3 MOA` expressed at that distance
- Velocity: `+/- 5 m/s`
- Time: `+/- 0.05 s`

### 12.7 RA4 Trajectory

Inputs:

- Drag model: `RA4`
- BC: `0.115`
- Weight: `40 gr`
- Muzzle velocity: `2000 fps`
- Zero distance: `100 yd`
- Sight height: `2.5 in`
- Sight angle: solved from the zero solver
- Max range: `500 yd`
- Report step: `50 yd`
- Wind: none

Expected sample count:

- `11`

Expected samples:

| Range yd | Velocity fps | Mach | Energy ft-lb | Drop in | Hold moa | Windage in | Wind adj moa | Time s | OGW lb |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| 0 | 2000.0 | 1.791 | 355.2 | -2.5 | n/a | 0.0 | n/a | 0.000 | 19.2 |
| 50 | 1726.5 | 1.546 | 264.7 | 0.2 | 0.4 | 0.0 | 0.0 | 0.081 | 12.3 |
| 100 | 1485.5 | 1.331 | 196.0 | 0.0 | 0.0 | 0.0 | 0.0 | 0.175 | 7.8 |
| 150 | 1270.0 | 1.138 | 143.2 | -4.2 | -2.7 | 0.0 | 0.0 | 0.284 | 4.9 |
| 200 | 1108.4 | 0.993 | 109.1 | -13.8 | -6.6 | 0.0 | 0.0 | 0.411 | 3.2 |
| 500 | 740.5 | 0.663 | 48.7 | -265.6 | -50.7 | 0.0 | 0.0 | 1.423 | 0.9 |

Tolerances:

- Velocity: `+/- 5 fps`
- Mach: `+/- 0.005`
- Energy: `+/- 5 ft-lb`
- Time: `+/- 0.06 s`
- OGW: `+/- 1 lb`
- Drop:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 2 in` from `500 yd` through `799 yd`
  - `+/- 4 in` at `800 yd` and beyond
- Windage:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 1 in` from `500 yd` through `799 yd`
  - `+/- 1.5 in` at `800 yd` and beyond
- Angular adjustments at nonzero range: `+/- 0.5` in the stated angular unit

### 12.8 Inclined-Shot G1 Trajectory

Inputs:

- Drag model: `G1`
- BC: `0.365`
- Weight: `65 gr`
- Muzzle velocity: `2600 fps`
- Zero distance: `100 yd`
- Sight height: `2.5 in`
- Sight angle: solved from the zero solver (see scenario 12.2b for expected value)
- Shot angle: `+10 deg`
- Cant: `0 deg`
- Max range: `1000 yd`
- Report step: `50 yd`
- Wind: none

Expected sample count:

- `21`

Expected samples:

| Range yd | Velocity fps | Mach | Energy ft-lb | Drop in | Hold moa | Windage in | Wind adj moa | Time s | OGW lb |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| 0 | 2600.0 | 2.329 | 975.5 | -2.5 | n/a | 0.0 | n/a | 0.000 | 111 |
| 100 | 2360.1 | 2.114 | 803.8 | 0.0 | 0.0 | 0.0 | 0.0 | 0.121 | 83 |
| 500 | 1537.4 | 1.378 | 341.1 | -67.4 | -12.9 | 0.0 | 0.0 | 0.753 | 23 |
| 1000 | 982.8 | 0.882 | 139.4 | -527.4 | -50.4 | 0.0 | 0.0 | 2.019 | 6 |

Tolerances:

- Velocity: `+/- 5 fps`
- Mach: `+/- 0.005`
- Energy: `+/- 5 ft-lb`
- Time: `+/- 0.06 s`
- OGW: `+/- 1 lb`
- Drop:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 2 in` from `500 yd` through `799 yd`
  - `+/- 4 in` at `800 yd` and beyond
- Windage:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 1 in` from `500 yd` through `799 yd`
  - `+/- 1.5 in` at `800 yd` and beyond
- Angular adjustments at nonzero range: `+/- 0.5` in the stated angular unit

Additional field expectations for the same inclined-shot scenario:

| Range yd | Flat dist yd | Drop flat in | LoS elev in | LoD elev in |
| --- | ---: | ---: | ---: | ---: |
| 0 | 0.00 | -2.50 | 0.00 | -2.50 |
| 100 | 98.69 | 626.54 | 626.47 | 629.30 |
| 500 | 492.58 | 3058.13 | 3126.80 | 3150.94 |
| 1000 | 984.93 | 5717.35 | 6252.14 | 6302.92 |

Additional tolerances:

- Flat distance: `+/- 0.5 yd`
- `drop flat`, `line of sight elevation`, and `line of departure elevation`: `+/- 4 in`

### 12.9 Round-Trip Unit Invariant

For each supported unit listed in Section 5:

- Construct a value in that unit.
- Read it back in the same unit.
- The round-tripped value must equal the original within `1e-7`.

This scenario should cover every listed angle, distance, velocity, energy, pressure, temperature, and weight unit, not merely one representative unit per family.

Additional angular invariant:

- `1 in/100yd` must equal `0.954930 moa +/- 1e-5`.
- Converting `1 in/100yd` to `cm/100m` and formatting in the converted unit should yield `2.78cm/100m`.

### 12.10 Standard Drag-Table Catalog And Validation

The rebuild shall include the standard table identifiers:

- `G1`
- `G2`
- `G5`
- `G6`
- `G7`
- `G8`
- `GS`
- `RA4`

Expected behavior:

- A positive-BC drag definition can be created for each listed table.
- An unknown standard-table identifier is rejected without panic.
- If the implementation has an internal custom-table sentinel for non-standard drag functions, that sentinel is not accepted by the standard-table constructor.
- For each listed table, evaluating the underlying raw drag curve at Mach `0.5`, `1.0`, and `2.0` yields the following values before the `PIR` scaling from Section 7.3:

| Table | Mach 0.5 | Mach 1.0 | Mach 2.0 |
| --- | ---: | ---: | ---: |
| G1 | 0.2032 | 0.4805 | 0.5934 |
| G2 | 0.1980 | 0.3983 | 0.2933 |
| G5 | 0.1603 | 0.3379 | 0.3861 |
| G6 | 0.2155 | 0.3597 | 0.3515 |
| G7 | 0.1194 | 0.3803 | 0.2980 |
| G8 | 0.2102 | 0.4068 | 0.3288 |
| GS | 0.4970 | 0.8140 | 1.0010 |
| RA4 | 0.2281 | 0.3975 | 0.5323 |

- Verification assets for this scenario should cite the public/reference drag data used to justify those sampled values rather than transcribing the original repository's full table literals.
- Reusing the same placeholder curve for multiple standard tables is non-compliant.

Tolerances:

- Raw drag values in the table above: `+/- 0.0005`

### 12.11 Atmosphere Construction And Humidity Normalization

Inputs:

- Altitude: `0 ft`
- Pressure: `29.92 inHg`
- Temperature: `59 F`
- Humidity input A: `0.78`
- Humidity input B: `78`

Expected behavior:

- Input A and input B produce the same normalized humidity, density factor, and local speed of sound.
- Humidity is normalized to `0.78`.
- Inputs below `0` or above `100` are rejected.

### 12.12 ICAO Atmosphere Spot Checks

Expected results:

- At `0 ft`: `59.0000 F`, `29.92000 inHg`, humidity `0`
- At `5000 ft`: `41.1692 F`, `24.89488 inHg`, humidity `0`

Tolerances:

- Temperature: `+/- 1e-4 F`
- Pressure: `+/- 1e-5 inHg`

### 12.13 Custom Drag-Function Scaling Spot Checks

Inputs:

- Drag mode: custom drag function
- Value type: ballistic coefficient
- BC: `1.0`
- Custom raw drag function: `raw_drag(mach) = 0.25 + 0.01 * mach`

Expected behavior:

- Through a direct drag-evaluation hook or verification helper, the raw drag function returns:
  - `0.255` at Mach `0.5`
  - `0.260` at Mach `1.0`
  - `0.270` at Mach `2.0`
- After applying the required `PIR` scaling from Section 7.3, the solver-visible drag values are:
  - `0.000053180505` at Mach `0.5`
  - `0.000054223260` at Mach `1.0`
  - `0.000056308770` at Mach `2.0`

Tolerances:

- Raw drag values: exact to within `1e-9`
- Scaled drag values: exact to within `1e-12`

### 12.14 Multi-Segment Wind Boundary On An Inclined Shot

Inputs:

- Drag model: `G1`
- BC: `0.365`
- Weight: `65 gr`
- Muzzle velocity: `2600 fps`
- Zero distance: `100 yd`
- Sight height: `2.5 in`
- Sight angle: solved from the zero solver
- Shot angle: `+10 deg`
- Cant: `0 deg`
- Max range: `1000 yd`
- Report step: `100 yd`
- Wind segments, ordered nearest to farthest:
  - Until `300 yd`: `5 mph` at `+90 deg`
  - Until `600 yd`: `5 mph` at `-90 deg`
  - Until `700 yd`: `5 mph` at `+90 deg`

Expected sample count:

- `11`

Expected selected samples:

| Range yd | Travel yd | Flat yd | Windage in | Time s |
| --- | ---: | ---: | ---: | ---: |
| 300 | 300.04 | 295.48 | 5.03 | 0.403 |
| 600 | 600.05 | 590.93 | 8.17 | 0.960 |
| 700 | 700.04 | 689.41 | 7.40 | 1.191 |
| 1000 | 1000.01 | 984.82 | 16.94 | 2.017 |

Interpretation requirements:

- Segment handoff is keyed to travelled slant distance, not flat distance.
- The final segment remains active beyond its own `until distance`; it does not shut off at `700 yd`.

Tolerances:

- Travel distance: `+/- 0.5 yd`
- Flat distance: `+/- 0.5 yd`
- Windage: `+/- 1 in`
- Time: `+/- 0.06 s`

### 12.15 Nonzero-Cant Wind Decomposition

Inputs:

- Drag model: `G1`
- BC: `0.223`
- Weight: `168 gr`
- Muzzle velocity: `2750 fps`
- Zero distance: `100 yd`
- Sight height: `2 in`
- Sight angle: `0.001228 rad`

> **Note:** This sight angle is intentionally NOT the G1/0.223 zero-solver output (which is `0.001265 rad`). The scenario re-uses the G7 zero-solver value from Section 12.2 to test cant-angle wind decomposition in isolation. The nonzero drop at 100 yards (`-0.69 in`) is expected.

- Shot angle: `0 deg`
- Cant: `+90 deg`
- Max range: `1000 yd`
- Report step: `100 yd`
- Wind: `5 mph` at `-45 deg`

Expected sample count:

- `11`

Expected selected samples:

| Range yd | Drop in | Windage in | Wind adj moa | Time s |
| --- | ---: | ---: | ---: | ---: |
| 100 | -0.69 | 0.00 | 0.00 | 0.118 |
| 500 | -108.07 | 0.02 | 0.00 | 0.857 |
| 1000 | -913.16 | 0.11 | 0.01 | 2.495 |

Interpretation requirement:

- This scenario intentionally verifies HEAD's limited cant model: cant rotates wind decomposition only; it does not rotate gravity or the drop reference frame.

Tolerances:

- Drop:
  - `+/- 0.5 in` below `500 yd`
  - `+/- 2 in` from `500 yd` through `799 yd`
  - `+/- 4 in` at `800 yd` and beyond
- Windage: `+/- 0.2 in`
- Windage adjustment: `+/- 0.05 moa`
- Time: `+/- 0.06 s`

### 12.16 Early-Termination Output Shape

Inputs:

- Drag model: `RA4`
- BC: `0.115`
- Weight: `40 gr`
- Muzzle velocity: `60 fps`
- Zero distance: `100 yd`
- Sight height: `2.5 in`
- Sight angle: `0 rad`
- Max range: `500 yd`
- Report step: `50 yd`
- Wind: none

Expected behavior:

- The returned trajectory slice length is `11`.
- Samples at indices `0` through `5` are populated.
- Samples at indices `6` through `10` remain default-zero values.
- A default-zero sample means all reported numeric fields remain zero-valued, including travelled distance, flat distance, velocity, mach, drop, windage, time, energy, and optimal game weight.

Expected selected samples:

| Index | Travel yd | Velocity fps | Drop in | Time s |
| --- | ---: | ---: | ---: | ---: |
| 0 | 0.00 | 60.0 | -2.50 | 0.000 |
| 5 | 250.00 | 252.9 | -94029.34 | 35.426 |
| 6 | 0.00 | 0.0 | 0.00 | 0.000 |

Tolerances:

- Travel distance at populated entries: `+/- 0.5 yd`
- Velocity at populated entries: `+/- 1 fps`
- Drop at populated entries: `+/- 100 in`
- Time at populated entries: `+/- 0.1 s`
- Zero-tail entries: exact

## 13. Verification Assets To Build

The rebuilt repository should include verification assets that exercise at least:

- The nineteen scenario families in Section 12.
- All advertised standard drag-table identifiers, not just `G1`, `G7`, and `RA4`.
- Standard-table drag spot checks against independently sourced reference data at multiple Mach points for every advertised table.
- Standard drag models `G1`, `G7`, and `RA4`.
- Direct custom drag-function support.
- Custom drag curve support.
- Mixed imperial and metric inputs.
- Zero solving.
- Wind.
- At least one multi-segment wind boundary case.
- Inclined shots.
- At least one inclined-shot windage-adjustment check that distinguishes slant-distance versus flat-distance reference behavior.
- At least one nonzero-cant case.
- Spin drift.
- An early-termination output-shape case that confirms the fixed-length zero-tail behavior.
- Unit round-tripping for every unit listed in Section 5.

The verification suite does not need to preserve the original Go test names.

## 14. Suggested Non-Normative Decomposition

An independent implementation will be easier to reason about if it is split into:

- A measurement/conversion layer.
- An atmosphere module.
- A drag-model module.
- A projectile/ammunition/weapon domain layer.
- A zero solver.
- A trajectory integrator.
- A verification suite with scenario fixtures.

This decomposition is guidance only.

## Appendix A: Observed Conversion Constants

The current repository uses the following exact conversion constants or formulas.

### Angles

- `rad` is canonical
- `deg -> rad = deg / 180 * pi`
- `MOA -> rad = moa / 180 * pi / 60`
- `mil -> rad = mil / 3200 * pi`
- `mrad -> rad = mrad / 1000`
- `thousand -> rad = thousand / 3000 * pi`
- `in/100yd -> rad = atan(value / 3600)`
- `cm/100m -> rad = atan(value / 10000)`

### Distance

- `1 ft = 12 in`
- `1 yd = 36 in`
- `1 mile = 63360 in`
- `1 nautical mile = 72913.3858 in`
- `1 line = 0.1 in`
- `1 mm = 1 / 25.4 in`
- `1 cm = 1 / 2.54 in`
- `1 m = 1000 / 25.4 in`
- `1 km = 1000000 / 25.4 in`

### Velocity

- `1 km/h = 1 / 3.6 m/s`
- `1 fps = 1 / 3.2808399 m/s`
- `1 mph = 1 / 2.23693629 m/s`
- `1 kt = 1 / 1.94384449 m/s`

### Energy

- `1 J = 0.737562149277 ft-lb`

### Pressure

- `1 inHg = 25.4 mmHg`
- `1 bar = 750.061683 mmHg`
- `1 hPa = 750.061683 / 1000 mmHg`
- `1 psi = 51.714924102396 mmHg`

### Temperature

- `F` is canonical
- `R -> F = R - 459.67`
- `C -> F = C * 9 / 5 + 32`
- `K -> F = (K - 273.15) * 9 / 5 + 32`

### Weight

- `grain` is canonical
- `1 g = 15.4323584 gr`
- `1 kg = 15432.3584 gr`
- `1 N = 151339.73750336 gr`

> **Clean-room note:** This Newton-to-grain constant does not match the standard physics derivation (`1 N ≈ 1573.66 gr`). It is the observed value in the current repository and is required if a rebuild wants to preserve HEAD's direct Newton conversions or any cross-unit checks involving Newtons. Section 12.9's same-unit round-trip invariant alone does not determine this constant.

- `1 lb = 1 / 0.000142857143 gr`
- `1 oz = 437.5 gr`
