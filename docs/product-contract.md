# v0.1 product contract

- Status: Normative target for v0.1
- Audience: Engine implementers, reviewers, and transport-adapter authors

The key words **MUST**, **MUST NOT**, **REQUIRED**, **SHOULD**, **SHOULD NOT**,
and **MAY** in this document express requirement strength for v0.1. They do not
describe the incomplete API currently in the repository.

## Product

`ballistic_calc` is a reusable Go engine for external-ballistics calculations.
It is independently designed from public physics references, cleared reference
data, and the requirements in this document.

The engine MUST provide:

- validated, unit-aware typed inputs;
- atmospheric conditions used by the flight model;
- G1 and G7 standard-drag evaluation in ballistic-coefficient mode;
- zero-angle solving;
- deterministic point-mass trajectory calculation; and
- structured, unit-aware samples suitable for APIs, tables, charts, DOPE
  views, and exports.

The engine MUST NOT own:

- HTTP routes, JSON DTOs, or public API versioning;
- authentication, authorization, rate limiting, accounts, or entitlements;
- databases, persistence, or saved-calculation lifecycle;
- KickBrass or Bench domain records;
- frontend state or components;
- charge weights, load recommendations, component substitutions, pressure
  diagnosis, internal-ballistics calculations, or safety certification.

## v0.1 calculation inputs

A v0.1 calculation MUST be expressible with typed Go values covering the
following information. Exact exported type and function names will be chosen
during implementation and are not specified here.

### Projectile and launch

- Standard drag model: `G1` or `G7`.
- Positive, finite ballistic coefficient appropriate to the selected model.
- Positive, finite projectile weight.
- Positive, finite muzzle velocity.

Ballistic coefficient is the only coefficient mode in v0.1. Projectile
diameter, length, twist, and form factor are not required inputs.

### Firearm and zero

- Finite, non-negative sight height above bore.
- Positive, finite zero range.

The engine MUST solve and report the bore-to-sight angle that zeros the modeled
trajectory at the requested range. A failure to converge within documented
limits MUST return a specific error rather than an unverified angle.

The v0.1 zero reference is a level, no-wind calculation in the supplied
atmosphere. Shot inclination and constant wind apply to trajectory generation;
they do not silently redefine the firearm's zero condition.

### Atmosphere

- Finite altitude above mean sea level.
- Positive, finite local station pressure at the firing location.
- Finite temperature above absolute zero.
- Relative humidity normalized to a fraction in `[0, 1]`; a percent input in
  `[0, 100]` MAY be accepted by a convenience constructor and MUST normalize
  unambiguously.

The atmosphere model and the interpretation of each field MUST be documented
with independently sourced references. Altimeter setting and station pressure
MUST NOT be silently interchanged.

### Shot and sampling

- A finite shot inclination relative to horizontal, greater than `-90` degrees
  and less than `+90` degrees. Zero is a level shot.
- A positive, finite maximum line-of-sight range.
- A positive, finite report interval no greater than the maximum range.
- Either no wind or one constant horizontal wind condition for the whole
  trajectory.

The constant wind SHOULD be represented inside the engine as downrange and
cross-range air-velocity components in shot coordinates. If a convenience API
accepts speed and direction, its direction convention MUST be explicit and its
conversion to components MUST be tested. Wind variation with range is not part
of v0.1.

The wind vector describes air velocity relative to the ground. Positive
downrange velocity is a tailwind, and positive cross-range velocity moves air
toward the right side of the shot coordinates.

The shot coordinate system MUST be documented. The v0.1 convention is:

- positive `x`: along the initial line of sight toward the target;
- positive `y`: vertical, opposite gravity; and
- positive `z`: right of the line of sight from the shooter's perspective.

## Units and normalization

Physical inputs and outputs MUST use the repository's typed unit layer rather
than bare values whose units are implied. The v0.1 surface MUST support common
imperial and metric representations for:

- angle;
- distance;
- velocity;
- energy;
- pressure;
- temperature; and
- projectile weight.

The existing unit implementation MUST be audited against independent standards
before release. Existing conversion constants or edge behavior are not
normative merely because they are already present.

The engine MUST normalize equivalent unit-system inputs into one calculation
state. Equivalent imperial and metric inputs MUST produce equivalent results
within the published unit-conversion and numerical tolerances.

## Drag and atmosphere behavior

G1 and G7 drag references MUST come from provenance-cleared direct sources or
from sources with permission sufficient for the intended distribution. The
engine MUST evaluate the supported Mach domain continuously enough for
trajectory integration; exact-point-only lookup is insufficient.

Interpolation, boundary behavior, and the supported Mach domain MUST be
documented and verified. Unsupported Mach values MUST produce documented
behavior or a calculation error, not an out-of-bounds access, silent table
substitution, or non-finite value.

Atmospheric density and local speed of sound MUST be calculated consistently
from the normalized atmospheric input. Any change in atmosphere with projectile
altitude MUST use one documented model in both zero solving and trajectory
calculation unless a verified design explicitly justifies a difference.

## Trajectory and zero behavior

The v0.1 engine MUST use one documented deterministic point-mass numerical
method. Zero solving MUST use the same force model and numerical core as the
trajectory calculation.

Requested samples MUST form a stable range grid:

1. a muzzle sample at range zero;
2. samples at successive report intervals that do not exceed maximum range;
3. a final sample exactly at maximum range when it is not already on the grid.

Reported sample range MUST identify the requested range, not an integration
overshoot. If the integrator does not land on a report range, the sampling
method MUST use a documented interpolation or dense-output method. A trajectory
that cannot reach a requested range MUST return an explicit error or an
explicitly incomplete result; it MUST NOT silently pad missing samples with
zero-valued records.

For identical normalized inputs, engine behavior version, and supported
execution environment, repeated calculations MUST return samples in the same
order with the same values. Bit-for-bit agreement with other calculators or
across every floating-point platform is not required.

## Result contract

A successful calculation result MUST contain:

- the solved zero angle;
- an engine behavior version;
- the normalized input values needed to reproduce the calculation; and
- ordered trajectory samples.

Each trajectory sample MUST contain these unit-aware values:

| Field | Meaning and sign |
| --- | --- |
| Range | Requested slant range along the initial line of sight |
| Time of flight | Elapsed time since the muzzle sample |
| Velocity | Projectile speed relative to the ground |
| Mach | Projectile air-relative speed divided by local speed of sound |
| Drop | Vertical displacement from the initial line of sight; positive is above |
| Elevation adjustment | Sight correction that opposes drop; positive is an upward correction |
| Windage | Cross-range displacement; positive is right |
| Wind adjustment | Sight correction that opposes windage; positive is a rightward correction |
| Energy | Kinetic energy derived from projectile mass and velocity |

At range zero, angular adjustments MUST be reported as zero rather than a
division-by-zero result. Every value in a successful result MUST be finite.

The engine behavior version MUST identify calculation semantics, not only the
source commit. A release MAY use the module semantic version when every
calculation-affecting change increments that version. Otherwise the result MUST
carry a separate behavior/schema identifier. Consumers MUST be able to persist
the identifier without parsing human-readable text.

## Validation and errors

The calculation entry point MUST validate all inputs before numerical work. It
MUST reject:

- `NaN`, positive infinity, and negative infinity in any numeric field;
- unsupported drag models or coefficient modes;
- non-positive ballistic coefficient, projectile weight, muzzle velocity,
  station pressure, zero range, maximum range, or report interval;
- negative sight height;
- temperature at or below absolute zero;
- humidity outside the accepted fractional or percent range;
- report intervals larger than maximum range;
- shot inclination outside the open interval `(-90 deg, +90 deg)`; and
- negative wind speed when a speed/direction convenience form is used.

Invalid caller data MUST return a structured or inspectable error. It MUST NOT
cause a panic or a successful result containing non-finite values. Convergence,
unsupported-domain, and integration failures MUST be distinguishable from
input-validation errors.

## Explicitly deferred

The following are outside the v0.1 promise:

- form-factor-to-BC workflows;
- arbitrary custom drag functions or custom curve APIs;
- G2, G5, G6, G8, GS, and RA4 release support;
- multiple wind segments;
- cant, spin drift, Coriolis, aerodynamic jump, and moving targets;
- internal ballistics or any reloading recommendation;
- compatibility with another library's API, output fixtures, convergence
  details, bugs, padding, names, or file layout;
- HTTP servers, accounts, persistence, and saved calculations in this module;
- a TypeScript solver, WASM, native bindings, mobile app, desktop UI, or CLI.

These topics can be reconsidered after v0.1 with new requirements, provenance,
verification, and safety review. Existing incomplete code for a deferred area
does not make that area supported.

## Definition of done

v0.1 is complete only when all of the following are true:

1. The inputs, validation, zero solving, trajectory behavior, outputs, and
   behavior version in this contract are implemented and documented.
2. G1 and G7 data have passed the release-clearance gate in
   [Provenance and licensing](provenance-and-licensing.md); uncleared and
   deferred tables are absent from the release artifact.
3. The layered evidence in [Verification strategy](verification.md) passes with
   published engineering tolerances and independent references.
4. Public Go API documentation and runnable examples describe only implemented
   behavior.
5. Repeated and unit-equivalent calculations meet determinism and equivalence
   requirements.
6. Legal/provenance and product-safety reviews required for public release are
   recorded; no unsupported claim of unrestricted use or real-world accuracy
   is made.
7. The module is tagged with a stable semantic version and its result reports a
   version identifier sufficient to preserve calculation history.

The safety and intended-use boundary in
[Safety and limitations](safety-and-limitations.md) applies to every v0.1 use.
