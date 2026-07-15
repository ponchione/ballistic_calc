# Verification strategy

## Goal

Verification builds independent evidence that `ballistic_calc` implements its
documented model within declared engineering tolerances. Passing values copied
from another implementation is not the acceptance model, and exact agreement
with another calculator is not a product goal.

The strategy is layered so a failed end-to-end case can be traced to units,
atmosphere, drag, analytic physics, numerical behavior, or transport mapping.

## Evidence record

Every reference-backed test set records:

- the behavior under test;
- complete normalized inputs and expected outputs;
- the direct source, publication identifier, page/table/equation locator, and
  retrieval information;
- any transcription, digitization, interpolation, or unit conversion;
- source and transformed-artifact checksums where practical;
- absolute and/or relative tolerance with units;
- the engineering reason for that tolerance; and
- the source-record identifier from
  [Provenance and licensing](provenance-and-licensing.md).

A tolerance is selected before comparing implementation output. It accounts for
source precision, input rounding, interpolation, integration error, and known
model differences. It is not widened simply to admit an observed result.

## Layer 1: units and conversion invariants

Coverage includes:

- round trips for every supported unit;
- independently sourced conversion spot checks across unit families;
- exact geometric relationships for angular units;
- imperial/metric representations of the same physical quantity;
- finite-value and invalid-unit behavior; and
- validation of calculation inputs after normalization.

Same-unit round trips alone do not validate a conversion constant because an
incorrect forward and inverse pair can still round-trip. At least one independent
cross-unit reference anchors every conversion family. The existing Newton
weight constant receives explicit review rather than automatic retention.

## Layer 2: atmosphere reference checks

Atmosphere evidence uses independently cited standard-atmosphere and
thermodynamic references. Coverage includes representative combinations of:

- sea-level and elevated altitude;
- imperial and metric inputs;
- temperature, local station pressure, and humidity;
- density or density ratio;
- local speed of sound; and
- any altitude-local update used during flight.

Checks distinguish station pressure from altimeter setting and state the
reference atmosphere and humidity assumptions. Expected values are calculated
or transcribed independently of the production functions.

## Layer 3: drag reference checks

Only provenance-cleared G1 and G7 references participate in v0.1 acceptance.
Evidence covers:

- raw source points across subsonic, transonic, and supersonic regions;
- transformed or embedded data against the source artifact;
- interpolation at source points and between points;
- continuity near segment boundaries;
- documented behavior at the supported Mach limits;
- positive BC scaling and invalid BC rejection; and
- finite results throughout the supported domain.

Tests derived from the same embedded literals can detect accidental code
changes but are not independent validation of those literals. The current JBM
transcriptions remain quarantine/integrity fixtures until replaced or cleared.

## Layer 4: analytic physics checks

The numerical core is exercised in configurations with closed-form or symmetry
results:

- vacuum/no-drag projectile motion against closed-form position and time;
- level and inclined geometry without aerodynamic forces;
- zero-wind lateral symmetry;
- mirrored constant crosswind cases where the documented model is symmetric;
- stationary or no-force invariants where the public numerical design permits
  them; and
- energy calculations from independently checked mass and velocity values.

Internal test seams for zero drag or simplified force models are verification
tools, not additional public v0.1 drag modes.

## Layer 5: property and metamorphic checks

Property tests cover useful behavior across a broad ordinary input domain:

- all successful numeric outputs are finite;
- sample range and time are strictly increasing after the muzzle sample;
- range samples exactly match the requested grid;
- velocity is generally non-increasing under positive drag in still air,
  allowing only explicitly justified numerical tolerance;
- zero wind produces zero lateral displacement;
- opposite crosswind components mirror lateral sign for otherwise identical
  shots within tolerance;
- equivalent unit-system inputs produce equivalent normalized results;
- repeated calculations are deterministic; and
- invalid inputs and unsupported domains fail explicitly without panic or
  zero-padded success.

Property generators stay inside documented v0.1 domains. Separate boundary
tests exercise rejection and convergence limits.

## Layer 6: independent end-to-end references

A small scenario set covers ordinary G1 and G7 use rather than exotic edge
cases. It includes, at minimum:

- one level no-wind G1 trajectory;
- one level constant-crosswind G1 trajectory;
- one level G7 trajectory;
- one inclined trajectory;
- imperial and metric representations of an equivalent scenario; and
- zero-angle results for both G1 and G7.

Each scenario is compared with at least two independently implemented reference
calculators or published worked examples. Reference choices document their drag
model, atmosphere, pressure interpretation, wind convention, zero convention,
and sampling definition. When references implement materially different models,
the evidence explains the difference rather than treating disagreement as an
engine defect or copying one result.

Acceptance uses field-specific tolerances for zero angle, time, velocity, Mach,
drop, windage, adjustment, and energy. Exact output cloning and bit-for-bit
comparison are excluded.

## Layer 7: KickBrass contract checks

This layer is later work in the KickBrass repository:

- engine package tests remain independent of HTTP and JSON;
- backend fixtures verify transport-to-engine field and unit mapping;
- engine errors map to stable transport error categories;
- engine behavior and transport-contract versions survive round trips;
- stateless calculations do not create snapshots;
- saved snapshots retain complete normalized inputs and outputs; and
- historical snapshots remain unchanged after an engine upgrade.

The adapter test suite imports and executes the Go engine. It does not implement
a TypeScript solver to create an artificial parity target.

## Current evidence status

The current tests provide useful foundation coverage:

- unit conversion examples and same-unit round trips;
- atmosphere constructor, humidity, formula, and altitude-update regression
  checks;
- standard table presence, ordering, and distinctness;
- exact-point G1 raw lookup; and
- raw-to-scaled drag checks for standard and custom-function scaffolding.

They do not yet satisfy v0.1 verification. Several expectations came from
superseded planning material rather than independently registered sources.
There is no arbitrary-Mach drag interpolation, analytic integrator suite, zero
solver, trajectory property suite, independent end-to-end scenario, or
KickBrass adapter.

## v0.1 verification gate

Release evidence is complete when:

1. All seven layers that apply to the Go engine are green; KickBrass contract
   checks remain a separate later gate.
2. Every reference-backed fixture links to a complete, cleared source record.
3. G1/G7 source points, transformations, and embedded artifacts are reproducible
   from recorded inputs.
4. The supported Go versions and execution environments are recorded.
5. Determinism, unit equivalence, and exact-range sampling pass across the
   supported ordinary domain.
6. Independent G1/G7 end-to-end scenarios meet predeclared tolerances against
   two references and explain material model differences.
7. Validation, convergence, and unsupported-domain failures are covered.
8. A reviewer can reproduce the evidence without consulting retired plans or
   any other library's source tree.
9. `go test ./...`, static analysis selected for the release, and provenance
   artifact checks pass from a clean checkout.
10. Legal/provenance and product-safety review gates are recorded separately;
    numerical test success is not treated as release permission or a safety
    certification.

The [v0.1 product contract](product-contract.md) remains the authority for
supported features and output semantics.
