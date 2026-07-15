# KickBrass integration

## Boundary

`ballistic_calc` remains a separate pure Go module. KickBrass integrates it
through the Go/Chi backend, which adapts a versioned transport contract to the
engine's typed inputs and results.

```text
ballistic_calc
  pure Go engine
       ^
       |
kickbrass-backend
  versioned transport adapter and optional snapshot persistence
       ^                         ^
       |                         |
KickBrass Bench           standalone ballistic-calculator web app
```

The dependency only points toward the engine. This repository has no KickBrass,
HTTP, database, authentication, UI, or persistence dependency.

## KickBrass context

KickBrass uses a Go/Chi/Postgres backend and a Nuxt/Vue web monorepo. Bench is
planned as a separate, private-by-default application for reloading records,
recipe revisions, batches or ammo lots, chronograph data, and range-test
history. Bench records what the user entered; it is not a load-data
recommendation or safety-certification product.

The current `apps/hf-calculator` surface is a USPSA hit-factor calculator. It
is not an external-ballistics calculator. A future standalone ballistic
calculator is a distinct web app that can reuse KickBrass theme, UI, app-shell,
auth-client, and API-client packages without reusing the hit-factor product
identity or calculation logic.

KickBrass currently defers calculator and loadbook architecture. No existing
route, payload, table, or frontend behavior is an integration commitment for
this engine.

## Responsibility split

| Owner | Responsibilities |
| --- | --- |
| `ballistic_calc` | Physical input types, calculation validation, atmosphere, G1/G7 drag, zero solving, point-mass integration, deterministic samples, engine behavior version |
| `kickbrass-backend` | Transport DTOs, JSON and HTTP validation, engine mapping, contract version, auth and entitlements where applicable, rate limits, observability, idempotency policy, optional snapshot storage |
| Standalone web app | Calculator forms, unit preferences, charts and tables, transient client state, explanation of predicted results |
| Bench | Private record selection, reviewable prefill, links to saved snapshots, predicted-versus-measured presentation, record access control |

The backend adapter calls the engine; it does not reproduce formulas or solver
logic. The web monorepo receives generated or reviewed TypeScript transport
types through the shared API client, not a second TypeScript numerical engine.

## Stateless hosted calculation

A standalone hosted calculator can use a stateless backend calculation without
saving calculator data. Authentication, abuse prevention, and rate-limit policy
are transport/product decisions recorded in the future KickBrass ADR.

Conceptual flow:

```text
web form
  -> versioned transport request
  -> backend transport validation and unit mapping
  -> typed Go engine calculation
  -> versioned transport response
  -> table/chart rendering
```

The first transport contract describes normalized physical fields, requested
output units, calculation result, engine behavior version, and transport
contract version. Exact URLs and JSON field names are defined later in a
KickBrass ADR and OpenAPI change; this repository does not invent them in
advance.

A successful stateless response is not persistence. Backend logs and telemetry
avoid recording sensitive Bench context or complete payloads unless an explicit
operational decision establishes a need, retention period, and access policy.

## Bench prefill

Bench can offer an explicit “Open in Ballistic Calculator” action from private
context such as:

- a firearm or test platform;
- a recipe revision;
- a batch or ammo lot;
- chronograph readings or summary; or
- a range test.

Prefill is a draft calculation assembled from fields the user can inspect.
Unknown or ambiguous values remain visibly unset. The calculator identifies the
source record and field where useful, but it never silently treats a catalog
match, recipe status, record completeness, or past verdict as proof of a valid
ballistic input.

The review screen separates:

- values copied directly from private records;
- values calculated from selected measured data, such as a user-approved
  muzzle-velocity summary;
- defaults supplied by the calculator; and
- fields still requiring user entry.

Charge weight, pressure observations, substitutions, and other reloading data
do not become solver inputs. Starting a calculation does not modify a recipe,
batch, test, or firearm record.

## Optional saved snapshots

Bench can persist a calculation only through a backend-owned snapshot model.
The saved record preserves:

- a complete normalized engine-input snapshot;
- the complete output snapshot as shown to the user;
- calculation timestamp;
- engine behavior version;
- transport-contract version;
- display-unit preferences needed to reproduce the view; and
- explicit links to source Bench records, where the user chose that context.

Snapshots are immutable calculation history. Editing source Bench records or
upgrading the engine does not rewrite old inputs or results. Recalculation
creates a new snapshot with its own versions and timestamp; the product can
compare the two without replacing history.

Persistence is optional. A user can calculate without saving, and a standalone
calculator does not require a Bench record. Snapshot retention, export,
deletion, privacy, and authorization belong to KickBrass rather than this Go
module.

## Versioning and failure mapping

The engine behavior version and transport-contract version identify different
boundaries:

- engine version: physics, reference data, validation, and numerical behavior;
- contract version: JSON field meanings, encoding, HTTP error mapping, and
  public API evolution.

The adapter preserves both. Engine validation, convergence, and unsupported
domain failures map to stable transport error categories without exposing
internal solver state. The transport layer can add request IDs and operational
metadata without changing engine results.

Backend contract tests verify DTO mapping, version preservation, and error
mapping. They call the Go engine as a dependency and contain no parallel solver
fixtures presented as another implementation.

## Predicted and measured data

Calculator output is always labeled as predicted or modeled. Chronograph
strings, velocities, group measurements, impacts, environmental observations,
and range notes remain measured or user-entered evidence.

Views that combine them use separate fields and visual treatment. A predicted
sample never overwrites a measured observation, and agreement between a model
and an observation does not certify ammunition, pressure, a backstop, or a safe
trajectory.

## Safety posture

Calculation results are estimates and are not load advice, pressure validation,
ammunition certification, or proof of safe firing conditions. KickBrass keeps
the complete boundary from [Safety and limitations](safety-and-limitations.md)
in the standalone calculator and Bench integration, with formal legal and
product-safety review before public release.

## KickBrass decision required before implementation

Integration begins with a small KickBrass ADR that reopens the currently
deferred calculator/loadbook boundary. That ADR records:

- the backend module owner and import/version policy for `ballistic_calc`;
- the transport contract and OpenAPI versioning approach;
- authentication/access and rate-limit policy for stateless calculations;
- optional Bench snapshot ownership, privacy, retention, and export behavior;
- prefill provenance and predicted/measured presentation rules;
- logging and telemetry data handling;
- standalone app registration, route, and shared-package usage; and
- legal, provenance, and product-safety release gates.

That ADR and all adapter, database, and web work belong in the KickBrass
repository. They are Milestone 5 in the [implementation roadmap](roadmap.md)
and do not block the engine's standalone v0.1 module release.

WASM, offline web execution, native mobile bindings, and desktop integration
remain later options. They do not change the first-release backend-adapter
boundary.
