# Architecture

## Outcome

`ballistic_calc` is a pure Go calculation module. It turns typed physical inputs
into a solved zero angle and ordered trajectory samples. The numerical engine
stays independent of KickBrass, transport formats, storage, authentication, and
user-interface concerns.

The [v0.1 product contract](product-contract.md) owns supported behavior. This
document owns the internal boundaries used to deliver it.

## Design principles

- One numerical implementation serves every consumer.
- Physical dimensions remain explicit through typed units.
- Validation happens before numerical integration.
- Zero solving and trajectory generation share the same physical model and
  integrator.
- Reference data is replaceable and traceable independently of solver code.
- Calculation-affecting changes are identifiable in results.
- Public APIs expose domain concepts, not transport or persistence concerns.

## Current implementation

The module currently targets Go 1.25.5 and uses the standard library. It has
three packages:

| Package | Present behavior | Work remaining for v0.1 |
| --- | --- | --- |
| `units` | Typed angle, distance, velocity, energy, pressure, temperature, and weight conversions | Independent standards audit, invalid-input policy, and v0.1 API review |
| `atmosphere` | Default and explicit construction, humidity normalization, density factor, speed of sound, ICAO-style altitude construction, and local altitude updates | Reference provenance, validation beyond humidity, model review, and integration with the solver |
| `drag` | Definition scaffolding, eight vendored table sets, exact-point raw lookup, and PIR scaling | G1/G7 provenance clearance or replacement, arbitrary-Mach interpolation, boundary behavior, BC validation, and solver integration |

The drag package also contains placeholders for custom functions and curves.
Those are incomplete and outside the v0.1 product promise. The currently
vendored G2, G5, G6, G8, GS, and RA4 tables are likewise not v0.1 release scope.

There is no projectile/ammunition/firearm/shot domain contract, zero solver,
trajectory integrator, constant-wind model, result sampler, or engine behavior
version yet. Exact-point drag lookup cannot support ordinary integration
between reference points.

Current tests verify implemented unit conversions, atmosphere calculations,
table shape, exact drag lookups, and scaling. They do not establish end-to-end
ballistic accuracy or public-release readiness.

## Target layers

Exact future package names remain an implementation decision. The ownership
layers are stable:

```text
public engine facade
  typed calculation input, validation orchestration,
  zero/trajectory entry points, versioned result
               |
               v
calculation domain and numerical core
  projectile/launch state, shot coordinates, constant wind,
  force evaluation, integration, zero solving, range sampling
          |                         |
          v                         v
atmosphere model               drag model
  density, speed of sound,       G1/G7 reference curves,
  altitude-local conditions      interpolation and BC scaling
          \                         /
           v                       v
                    units
          typed physical quantities and conversion
```

The public facade is small: normalized typed inputs in, a versioned result or
an inspectable error out. It hides integrator state, reference-table storage,
interpolation coefficients, root-finding iteration state, and temporary
allocations.

## Responsibility boundaries

### Unit layer

The unit layer owns physical quantity representation and conversion. It has no
ballistic-domain, atmosphere, drag, solver, or consumer dependencies.

Calculations normalize values at the engine boundary and convert for display
only at consumer boundaries. The numerical core does not mix implicit unit
systems or pass unitless values across public boundaries.

### Atmosphere layer

The atmosphere layer owns the meaning and normalization of altitude,
temperature, local station pressure, and humidity. It supplies density and
local speed of sound to the numerical core from one documented atmosphere
model.

It has no knowledge of projectiles, firearms, range samples, HTTP requests, or
Bench records.

### Drag layer

The drag layer owns the cleared G1/G7 reference curves, interpolation, supported
Mach domain, and conversion from ballistic coefficient plus local flight state
to the drag term needed by the force model.

Reference data and its metadata are isolated from interpolation and solver
logic. Replacing a source table does not require a transport or application
change, while a calculation-affecting replacement changes the engine behavior
version.

### Calculation domain

The domain layer owns the typed meaning of projectile/launch, firearm/zero,
atmosphere, shot, sampling, and constant-wind inputs. It also owns normalized
results and sample semantics.

It contains no JSON tags added solely for a public API, database annotations,
account IDs, recipe IDs, or frontend presentation fields. A consumer can keep
its own correlation identifier outside the engine call.

### Numerical core

The numerical core owns:

- shot-coordinate transforms;
- gravity and air-relative velocity;
- atmospheric and drag force evaluation;
- the documented point-mass integration method;
- zero-angle root finding over that same integration path;
- exact-range sampling or documented dense output; and
- finite-value and convergence checks.

The core operates on an immutable normalized input for a calculation. Mutable
integration state remains local to the call. No global mutable solver state,
clock reads, randomness, network access, filesystem access, or database access
participates in a result.

### Public engine facade

The facade coordinates validation, normalization, zero solving, trajectory
calculation, and result construction. It reports the engine behavior version
and errors that distinguish validation, unsupported-domain, convergence, and
integration failures.

Current exported symbols are implementation scaffolding. The facade becomes a
stable public API only when its contract, examples, and versioning policy are
reviewed for the v0.1 tag.

## Dependency rules

- Lower layers do not import the public facade or consumer applications.
- `units` remains dependency-free within this module.
- Atmosphere and drag remain independently testable without a trajectory.
- The numerical core consumes atmosphere and drag through narrow internal
  behavior, not through consumer-specific DTOs.
- Reference fixtures do not become runtime authority merely because tests use
  them.
- This module contains no imports from KickBrass repositories.
- KickBrass depends on a released version of this module; the dependency never
  points back toward KickBrass.
- HTTP, JSON, auth, storage, and public API versioning live outside this module.

## Determinism and sampling

Determinism means the same normalized input and engine behavior version produce
the same ordered sample values on a supported execution environment. It does
not claim bit-identical output from unrelated calculators or every processor
and compiler combination.

The implementation records all calculation choices that affect this promise:

- drag reference revision and interpolation behavior;
- atmosphere equations and constants;
- integration and error-control method;
- zero-solver convergence policy;
- range-sampling behavior; and
- output formulas and sign conventions.

Sampling is based on requested ranges, not on whatever integration step happens
to cross a threshold. The engine uses interpolation, dense output, or controlled
step placement so the reported range grid follows the product contract.
Integration failure is explicit; missing tail samples are never disguised as
zero-valued successful samples.

## Versioning

Two version surfaces are separate:

1. The Go module semantic version describes the public code release.
2. The engine behavior version stored in a calculation identifies numerical
   semantics and reference-data revisions.

A module release can use one identifier for both when every behavior change
forces an appropriate semantic-version change. Otherwise the engine exposes a
separate opaque behavior identifier. Consumers store it verbatim.

Changes to validation, constants, drag data, interpolation, atmosphere,
integration, convergence, sampling, or output formulas count as behavior
changes. Documentation corrections that do not alter calculations do not.

## Integration boundary

The engine produces typed Go values. A future KickBrass backend adapter owns
transport DTOs, API versioning, authentication, rate limits, and optional
snapshot persistence. Bench and a standalone web application consume the
adapter rather than embedding solver math in TypeScript.

The full boundary is in [KickBrass integration](kickbrass-integration.md).
WASM, native bindings, and offline execution remain later architecture choices,
not reasons to distort the v0.1 Go design.

## Data and release boundary

Build-time embedding does not establish permission to distribute data. The
current table literals remain quarantined under
[Provenance and licensing](provenance-and-licensing.md). Only cleared G1/G7
assets are part of the intended v0.1 release, and their source records travel
with release evidence.
