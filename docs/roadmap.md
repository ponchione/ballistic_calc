# Implementation roadmap

## Current position

The repository has a unit foundation, atmosphere calculations, and partial drag
scaffolding. It does not have a usable external-ballistics engine. The product
reset uses five focused engine milestones and one separate KickBrass integration
milestone.

Milestone 0 is active. The authoritative documentation and data quarantine are
in place; direct-source and rights work for G1/G7 remains open. The next code
milestone is Milestone 1 after that source decision is recorded.

## Milestone 0: Scope and provenance reset

Outcome: implementation can proceed from current product requirements without
carrying obsolete scope assumptions or unclear release claims.

Work:

- establish the README and focused authoritative documents;
- retire superseded plans and operational handoffs;
- inventory current tables, formulas, constants, and fixtures;
- quarantine current JBM-transcribed drag literals;
- identify candidate direct sources for G1/G7 and record authorship,
  publication, retrieval, transformation, and rights analysis;
- define the intended original-code `0BSD` posture and separate data terms;
- record verification, safety, and legal/provenance release gates; and
- leave Go behavior unchanged during the documentation reset.

Exit criteria:

- every active document is reachable from the documentation index;
- current implementation status and v0.1 scope are unambiguous;
- no superseded plan appears current;
- all known non-original material has an inventory entry and status; and
- G1/G7 direct-source candidates and the decision path for replacement or
  permission are documented.

## Milestone 1: G1/G7 drag foundation

Outcome: the engine can evaluate cleared G1 and G7 drag references throughout a
documented Mach domain in BC mode.

Work:

- replace current G1/G7 literals with reproducibly generated, provenance-cleared
  data or record sufficient distribution permission;
- remove deferred and uncleared tables from the release build path;
- define a positive finite BC-only drag input;
- implement arbitrary-Mach interpolation and explicit domain boundaries;
- document raw-reference meaning and scaling/force derivation;
- reject invalid models, BC values, and unsupported Mach conditions; and
- build independent source-point, interpolation, continuity, and boundary
  tests.

Exit criteria:

- G1/G7 artifacts reproduce from recorded direct sources and transformations;
- no v0.1 runtime path selects G2, G5, G6, G8, GS, RA4, or custom drag;
- ordinary Mach values no longer depend on exact table-point matches;
- drag outputs remain finite and continuous within the supported domain; and
- verification layers 1 and 3 pass for the completed drag surface.

## Milestone 2: Domain and calculation contract

Outcome: callers can construct and validate a complete v0.1 calculation without
depending on solver internals.

Work:

- add typed projectile/launch, firearm/zero, atmosphere, shot, sampling, and
  constant-wind inputs;
- audit the existing unit and atmosphere layers against independent sources;
- define normalization and finite-value validation;
- define zero and trajectory result/sample types;
- establish coordinate, sign, pressure, wind, and range conventions;
- add inspectable validation error categories; and
- add the engine behavior-version field and change policy.

Exit criteria:

- every input and output in the product contract has one unambiguous typed
  representation;
- invalid and nonsensical inputs fail before numerical work and without panic;
- equivalent imperial and metric inputs normalize equivalently;
- domain packages contain no transport, persistence, UI, or KickBrass fields;
  and
- public API documentation describes only the validated contract, clearly
  marking calculation entry points unfinished until Milestone 3.

## Milestone 3: Core trajectory and zero solver

Outcome: one deterministic numerical core produces solved zero angles and
requested-range samples for level/inclined shots with no wind or constant wind.

Work:

- choose and document a point-mass integration and error-control method;
- implement gravity, atmospheric conditions, air-relative velocity, and G1/G7
  drag in shot coordinates;
- implement zero-angle root finding over the same integration path;
- implement level and inclined geometry plus one constant wind vector;
- sample exactly at the requested range grid;
- calculate time, velocity, Mach, drop, elevation adjustment, windage, wind
  adjustment, and energy;
- return explicit convergence, domain, and incomplete-trajectory errors; and
- keep each calculation isolated from global mutable state, time, and
  randomness.

Exit criteria:

- all v0.1 input paths produce a versioned result or an inspectable error;
- zero and trajectory calculations use the same force and atmosphere model;
- repeated runs are deterministic on supported environments;
- outputs are finite, ordered, and located at requested ranges;
- vacuum, symmetry, property, and unit-equivalence tests pass; and
- no missing result tail is represented as successful zero samples.

## Milestone 4: Independent verification and v0.1 release

Outcome: the module has enough reproducible evidence, documentation, and review
to tag a stable v0.1 release.

Work:

- complete the layered verification suite;
- compare ordinary G1/G7 scenarios with at least two independent reference
  implementations or published examples using predeclared tolerances;
- add runnable Go examples and public API documentation;
- run performance, allocation, convergence, and broad sanity checks;
- record supported Go versions and environments;
- complete provenance/legal and product-safety review;
- remove every uncleared or out-of-scope release asset;
- apply the intended license only to original material that can be licensed
  that way and add any required third-party notices; and
- tag a stable Go module release with version-identifiable calculation
  behavior.

Exit criteria:

- the [v0.1 definition of done](product-contract.md#definition-of-done) is
  complete;
- every release fixture and reference artifact links to a cleared source record;
- all required tests and static checks pass from a clean checkout;
- examples reproduce using the tagged module API; and
- release notes state model limits and evidence without claiming guaranteed
  real-world accuracy or unrestricted legal status.

## Milestone 5: KickBrass adapter and apps

Repository: KickBrass, not `ballistic_calc`.

Outcome: a single Go engine serves a stateless standalone calculator and
optional private Bench workflows through a versioned backend contract.

Work:

- adopt a KickBrass ADR reopening the deferred calculator/loadbook boundary;
- import a pinned engine release into the Go/Chi backend;
- define the versioned stateless transport adapter and OpenAPI contract;
- build a distinct Nuxt/Vue ballistic-calculator app using shared KickBrass web
  packages;
- add reviewable Bench prefill from selected private records;
- add optional immutable input/output snapshots with engine and contract
  versions;
- preserve predicted-versus-measured separation; and
- complete KickBrass privacy, rate-limit, telemetry, legal, provenance, and
  product-safety release gates.

Exit criteria:

- transport mapping has contract tests without a duplicated TypeScript solver;
- stateless calculations work without persistence;
- saved results remain historical snapshots across engine upgrades;
- Bench prefill exposes its sources and requires user review; and
- the standalone calculator is distinct from the existing hit-factor app.

Milestone 5 does not block the Go engine's standalone v0.1 release.

## Post-v0.1 possibilities

Later product work can evaluate additional drag models, custom curves,
range-varying wind, cant, spin drift, Coriolis, aerodynamic jump, moving
targets, CLI tools, WASM, native bindings, mobile/offline use, desktop UI, or
self-hosted application packaging.

Each extension gets its own product requirement, provenance record, numerical
model, verification evidence, versioning decision, and safety review. None is
part of the v0.1 completion path by implication.
