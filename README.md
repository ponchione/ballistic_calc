# ballistic_calc

`ballistic_calc` is a small Go module for deterministic external-ballistics
calculations. It is intended to power both a public standalone calculator and
ballistics features reached through KickBrass Bench, while remaining useful to
other Go programs without any KickBrass dependency.

The project is an independent design based on publicly documented physics,
independently sourced reference material, and its own product requirements. It
is not a drop-in replacement for another library, and reproducing another
calculator's API, quirks, or exact outputs is not a goal.

It exists so the owner controls the engine's design and can use one Go module
across a hosted backend, a standalone calculator, and possible future CLI,
desktop, mobile, self-hosted, or offline products. The intended permissive
license supports that distribution flexibility. This direction does not imply
that LGPL software cannot be used privately or on a hosted backend; it is a
product ownership and downstream-distribution choice, not legal advice.

## Product boundary

The engine's v0.1 target is deliberately narrow:

- unit-aware, validated inputs;
- atmosphere and standard G1/G7 drag modeling;
- ballistic-coefficient inputs;
- zero-angle solving;
- level and inclined point-mass trajectories;
- no wind or one constant wind condition;
- deterministic samples for range, time, velocity, Mach, drop, adjustments,
  windage, and energy; and
- imperial and metric input/output through the Go unit layer.

The module is a calculation engine, not an end-user application. It does not
own HTTP or JSON contracts, accounts, databases, saved calculations, frontend
components, Bench records, load recommendations, pressure diagnosis, or safety
certification. The normative requirements and deferred features are in the
[v0.1 product contract](docs/product-contract.md).

## Current status

The engine is not yet usable for trajectory calculations.

Implemented today:

- `units`: angle, distance, velocity, energy, pressure, temperature, and weight
  conversions, including imperial and metric units;
- `atmosphere`: constructors, humidity normalization, density and speed-of-sound
  calculations, ICAO-style altitude conditions, and altitude-local updates;
- `drag`: drag-definition scaffolding, vendored reference tables, exact-point
  raw lookup, and PIR scaling; and
- tests for those implemented slices.

Not yet implemented:

- provenance-cleared G1/G7 release data and arbitrary-Mach interpolation;
- the v0.1 projectile, ammunition, firearm/zero, shot, and wind input model;
- comprehensive validation and version-identifiable results;
- a zero-angle solver;
- a trajectory integrator; and
- the v0.1 trajectory sample/result model.

Exact-point drag lookup is only scaffolding; it is not sufficient for a
working solver. The current drag-table literals are also quarantined from a
public release until their provenance and distribution rights are resolved.
See [Provenance and licensing](docs/provenance-and-licensing.md).

No stable calculation API or usage example exists yet. Current exported Go
symbols describe incomplete foundations and are not a v0.1 compatibility
promise. Development starts from the [implementation roadmap](docs/roadmap.md).

## KickBrass direction

The Go engine remains transport- and storage-independent. A future KickBrass
backend adapter will map versioned transport inputs into typed engine inputs
and can serve both Bench and a distinct standalone ballistic-calculator web
app. The existing KickBrass hit-factor calculator is a different product.

The dependency boundary, stateless calculation flow, optional immutable
snapshots, and Bench prefill rules are described in the
[KickBrass integration guide](docs/kickbrass-integration.md).

## Safety

External ballistics is an estimate of a complex physical process. Results
depend on input quality, model limits, equipment and ammunition variation,
weather, and numerical choices. This library does not guarantee impacts,
backstops, safe trajectories, or environmental conditions, and it is not for
safety-critical firing decisions. Verify results through appropriate real-world
procedures and follow firearm-safety practices and applicable law.

The project does not provide reloading data, recommend loads, validate
pressure, certify ammunition, or replace published manuals. See
[Safety and limitations](docs/safety-and-limitations.md) for the full product
boundary.

## Licensing status

Original code created for this repository is intended for eventual release
under the `0BSD` license because it places minimal conditions on downstream
use. The repository is not currently declared wholly `0BSD`, and no
repository-wide license has been applied. Imported data and other non-original
material require separate provenance, rights, and attribution review.

In particular, the current standard-drag tables were transcribed from
JBM-hosted files. Their direct provenance and reuse status have not been
cleared for the intended distribution. G1 and G7 need cleared replacement data
or sufficient permission before v0.1 can ship; deferred tables cannot ship
unless separately cleared. This is a release gate, not a legal conclusion.

## Documentation

- [Documentation index](docs/README.md)
- [v0.1 product contract](docs/product-contract.md)
- [Architecture](docs/architecture.md)
- [KickBrass integration](docs/kickbrass-integration.md)
- [Verification strategy](docs/verification.md)
- [Provenance and licensing](docs/provenance-and-licensing.md)
- [Safety and limitations](docs/safety-and-limitations.md)
- [Implementation roadmap](docs/roadmap.md)
