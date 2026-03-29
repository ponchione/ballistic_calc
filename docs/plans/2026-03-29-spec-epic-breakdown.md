# External Ballistics Library Reconstruction Implementation Plan

> For Hermes: Use subagent-driven-development skill to implement this plan task-by-task.

Goal: Turn `spec.md` into a staged implementation roadmap for a fresh clean-room Go ballistics library with verification assets.

Architecture: Build the library in small packages with hard boundaries: `units`, `atmosphere`, `drag`, `ballistic`, `zero`, `trajectory`, and `testdata`/verification helpers. Treat Section 12 acceptance scenarios as the source of truth and use them as milestone gates, not just final regression tests.

Tech Stack: Go 1.25, standard library, table-driven tests, embedded fixture data where helpful.

---

## Proposed package layout

- `units/`
  - `angle.go`
  - `distance.go`
  - `velocity.go`
  - `energy.go`
  - `pressure.go`
  - `temperature.go`
  - `weight.go`
  - `units_test.go`
- `atmosphere/`
  - `atmosphere.go`
  - `icao.go`
  - `local.go`
  - `atmosphere_test.go`
- `drag/`
  - `types.go`
  - `standard_tables.go`
  - `standard_data.go`
  - `bc.go`
  - `curve.go`
  - `custom.go`
  - `drag_test.go`
- `ballistic/`
  - `projectile.go`
  - `ammunition.go`
  - `weapon.go`
  - `shot.go`
  - `wind.go`
  - `domain_test.go`
- `zero/`
  - `solver.go`
  - `solver_test.go`
- `trajectory/`
  - `integrator.go`
  - `sample.go`
  - `wind.go`
  - `spin.go`
  - `trajectory_test.go`
- `testdata/`
  - `drag_reference.md`
  - `scenarios/`
- `docs/`
  - `plans/`
  - `drag-data-sources.md`

Notes:
- Keep the public API thin until the numerical core is stable.
- Do not optimize for original Go API compatibility; optimize for behavioral coverage.
- Make tests consume scenario fixtures instead of duplicating numbers inline across many files.

---

## Epic 1: Measurement foundation

Outcome: every measurement family in Section 5 exists, converts correctly, and round-trips within tolerance.

Primary files:
- Create `units/*.go`
- Create `units/units_test.go`

Covers:
- Section 5
- Appendix A
- Scenario 12.9

Tasks:
1. Create one value type per measurement family with a canonical internal unit.
2. Add constructors and getters for every required unit in Section 5.
3. Implement the geometric angle conversions for `in/100yd` and `cm/100m` using `atan`, not linear approximation.
4. Encode Appendix A constants exactly, including the observed Newton conversion quirk.
5. Add a round-trip test table for every supported unit.
6. Add explicit invariant tests for `1 in/100yd == 0.954930 moa +/- 1e-5` and `1 in/100yd -> 2.78 cm/100m` formatting/representation.

Exit criteria:
- Scenario family 12.9 passes.
- No other package defines ad hoc conversion logic.

Suggested commit slices:
- `feat: add core measurement types`
- `test: add round-trip unit invariants`

---

## Epic 2: Atmosphere model

Outcome: atmosphere creation, ICAO construction, density factor, speed of sound, and in-flight altitude updates are stable and testable.

Primary files:
- Create `atmosphere/atmosphere.go`
- Create `atmosphere/icao.go`
- Create `atmosphere/local.go`
- Create `atmosphere/atmosphere_test.go`

Covers:
- Section 6
- Scenario 12.11
- Scenario 12.12

Tasks:
1. Implement default-atmosphere construction with the normative `RH=0.78` default.
2. Implement explicit atmosphere construction with humidity normalization and input validation.
3. Implement ICAO-by-altitude construction from the specified formulas.
4. Implement density factor and speed-of-sound formulas from Section 6.4.
5. Implement a helper that recomputes local density/mach for a projectile altitude relative to a base atmosphere, including the `30 ft` reuse rule and `3.28084 ft` refresh threshold behavior.
6. Add tests for default values, humidity normalization (`0.78` vs `78`), invalid humidity rejection, and ICAO spot checks.

Exit criteria:
- Atmosphere tests prove both construction-time and in-flight update behavior.
- Later packages can depend on `atmosphere.Atmosphere` without re-deriving formulas.

Suggested commit slices:
- `feat: add atmosphere constructors and formulas`
- `test: add humidity and ICAO atmosphere coverage`

---

## Epic 3: Drag subsystem

Outcome: standard tables, BC vs FF handling, PIR scaling, custom raw functions, and curve interpolation all exist behind one coherent drag interface.

Primary files:
- Create `drag/types.go`
- Create `drag/standard_tables.go`
- Create `drag/standard_data.go`
- Create `drag/bc.go`
- Create `drag/curve.go`
- Create `drag/custom.go`
- Create `drag/drag_test.go`
- Create `docs/drag-data-sources.md`
- Create `testdata/drag_reference.md`

Covers:
- Section 7
- Scenario 12.5
- Scenario 12.10
- Scenario 12.13

Tasks:
1. Define the external drag modes: standard table, custom function, and custom curve helper path.
2. Add strict validation for supported identifiers: `G1`, `G2`, `G5`, `G6`, `G7`, `G8`, `GS`, `RA4`.
3. Load independently sourced standard drag reference data and document provenance.
4. Implement raw drag lookup and PIR scaling (`2.08551e-04`).
5. Implement BC mode and FF mode, with FF rejecting missing or nonpositive diameter.
6. Implement custom curve precomputation and evaluation with the documented edge behavior.
7. Add spot-check tests for every standard table at Mach `0.5`, `1.0`, and `2.0`.
8. Add tests for custom drag-function scaling and FF-to-BC conversion.

Exit criteria:
- Every advertised table is independently distinguishable.
- Custom drag behavior is testable without the trajectory engine.

Suggested commit slices:
- `feat: add standard drag table catalog`
- `feat: add custom drag function and curve support`
- `test: add drag scaling and catalog validation`

---

## Epic 4: Ballistic domain objects

Outcome: projectiles, ammunition, weapons, shots, and winds can be constructed and validated without any solver logic leaking into them.

Primary files:
- Create `ballistic/projectile.go`
- Create `ballistic/ammunition.go`
- Create `ballistic/weapon.go`
- Create `ballistic/shot.go`
- Create `ballistic/wind.go`
- Create `ballistic/domain_test.go`

Covers:
- Section 4
- Required defaults and validation rules consumed later by Sections 8 and 9

Tasks:
1. Implement `Projectile` with drag description, weight, and optional dimensions.
2. Implement `Ammunition` with projectile plus muzzle velocity.
3. Implement `Weapon` with sight height, zero distance, and optional twist metadata.
4. Implement `Shot` with sight angle, max range, report step, optional shot angle, and optional cant defaulting to zero.
5. Implement piecewise-constant wind segments with normalized “no wind” handling.
6. Add validation tests for required inputs and defaults.

Exit criteria:
- Zero solver and trajectory engine accept fully validated objects.
- Passive metadata not required by the spec is omitted unless intentionally added later.

Suggested commit slices:
- `feat: add projectile and ammunition models`
- `feat: add weapon shot and wind models`

---

## Epic 5: Zero solver

Outcome: sight angle solving works independently of the full reporting engine and matches the normative zero scenarios.

Primary files:
- Create `zero/solver.go`
- Create `zero/solver_test.go`

Covers:
- Section 8
- Scenarios 12.1, 12.2, 12.2a, 12.2b, 12.2c

Tasks:
1. Implement muzzle-state initialization at `(0, -sight_height, 0)`.
2. Implement a solver that uses muzzle atmosphere values for the entire solve.
3. Match the behavioral contract around iterative correction and convergence tolerance, even if the internal numeric method differs.
4. Add one focused test per zero scenario from Section 12.
5. Add a regression test proving RA4 zero solving works, not just G1/G7.

Exit criteria:
- All five zero-solver scenario families pass.
- The solver package is callable directly by higher-level trajectory code.

Suggested commit slices:
- `feat: add zero solver`
- `test: add zero solver scenario coverage`

---

## Epic 6: Core trajectory engine

Outcome: a level-shot trajectory with report thresholds, sample emission, core outputs, and early-stop shape works before advanced geometry is layered on.

Primary files:
- Create `trajectory/integrator.go`
- Create `trajectory/sample.go`
- Create `trajectory/trajectory_test.go`

Covers:
- Section 9.1 through 9.5
- Section 9.8 through 9.9
- Scenario 12.3 (partially, once basic wind is present)
- Scenario 12.7
- Scenario 12.16

Tasks:
1. Implement the state vector, gravity, and launch vector setup.
2. Implement report-threshold generation and preallocated output sizing.
3. Implement the internal reduced integration step behavior or an equivalent method that still matches fixtures.
4. Implement sample emission on threshold crossing without back-interpolation.
5. Populate sample fields: time, travel distance, flat distance, velocity, mach, drop, adjustments, energy, and optimal game weight.
6. Implement early termination on speed `< 50 fps` or raw `y < -15000 ft` while preserving the fixed-length zero tail.
7. Add trajectory tests for RA4 no-wind behavior and the early-termination output shape.

Exit criteria:
- Basic trajectory output shape is stable.
- No advanced wind/cant/spin logic is required to pass RA4 and early-termination coverage.

Suggested commit slices:
- `feat: add core trajectory integrator`
- `test: add RA4 and early-termination trajectory coverage`

---

## Epic 7: Wind, inclined shots, cant, and spin drift

Outcome: the advanced geometry and lateral behavior match HEAD-compatible semantics.

Primary files:
- Create `trajectory/wind.go`
- Create `trajectory/spin.go`
- Modify `trajectory/integrator.go`
- Modify `trajectory/sample.go`
- Modify `trajectory/trajectory_test.go`

Covers:
- Section 9.6
- Section 9.7
- Section 9.10
- Scenario 12.3
- Scenario 12.4
- Scenario 12.8
- Scenario 12.14
- Scenario 12.15

Tasks:
1. Implement wind-vector decomposition with the documented direction semantics.
2. Apply wind against air-relative velocity, not ground-relative velocity.
3. Implement multi-segment wind switching on travelled slant distance, with the final segment persisting.
4. Implement inclined-shot drop rotation and distinct slant-vs-flat reference behavior for adjustments.
5. Implement cant as a wind-decomposition rotation only, preserving the documented limitation.
6. Implement spin drift gated on both twist metadata and projectile dimensions.
7. Add scenario tests for constant wind G1, G7 with spin drift, inclined G1, multi-segment wind boundary behavior, and nonzero-cant decomposition.

Exit criteria:
- Advanced flight behavior is proven by the exact Section 12 scenario families.
- No ambiguity remains around slant distance vs flat distance.

Suggested commit slices:
- `feat: add wind vector and segment switching`
- `feat: add inclined-shot and cant handling`
- `feat: add spin drift`
- `test: add advanced trajectory scenario coverage`

---

## Epic 8: Verification assets and repository finishing pass

Outcome: the repo contains a maintainable verification suite and supporting documentation, not just passing ad hoc tests.

Primary files:
- Create `testdata/scenarios/*`
- Create package-specific `*_test.go` fixture loaders/helpers as needed
- Update `docs/drag-data-sources.md`
- Update `README.md` if/when introduced

Covers:
- Section 12 in full
- Section 13

Tasks:
1. Move hard-coded scenario numbers into fixture tables grouped by scenario family.
2. Add a coverage matrix mapping every Section 12 scenario to one or more tests.
3. Add provenance notes for standard drag reference data.
4. Add a smoke test that verifies all advertised standard drag tables are constructible.
5. Add a final regression test run command list for local development.

Exit criteria:
- Every scenario family in Section 12 is represented.
- The repo explains why the standard drag values are trustworthy in a clean-room context.

Suggested commit slices:
- `test: consolidate scenario fixtures`
- `docs: add drag data provenance and verification notes`

---

## Priority order

Recommended execution order:
1. Epic 1 — Measurements
2. Epic 2 — Atmosphere
3. Epic 3 — Drag subsystem
4. Epic 4 — Domain objects
5. Epic 5 — Zero solver
6. Epic 6 — Core trajectory engine
7. Epic 7 — Advanced trajectory behavior
8. Epic 8 — Verification/documentation consolidation

Reasoning:
- Units, atmosphere, and drag are hard prerequisites for every numerical path.
- Zero solving is simpler than the full trajectory reporter and gives an early correctness checkpoint.
- Advanced wind/cant/spin behavior should land only after the core integrator is numerically stable.

## Smallest sensible first milestone

If we want to stay disciplined and attack one problem at a time, the cleanest first milestone is:
- Epic 1 only
- specifically: measurement types + round-trip invariants + angular geometric conversions

That gives us one isolated subsystem, one clean local commit chain, and a base the rest of the library can trust.

## Acceptance coverage map

- Epic 1: 12.9
- Epic 2: 12.11, 12.12
- Epic 3: 12.5, 12.10, 12.13
- Epic 5: 12.1, 12.2, 12.2a, 12.2b, 12.2c
- Epic 6: 12.7, 12.16
- Epic 7: 12.3, 12.4, 12.8, 12.14, 12.15
- Epic 8: closes the gaps and ensures all Section 12 families are represented consistently
