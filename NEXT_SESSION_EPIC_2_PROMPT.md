Handle Epic 2 for the `ballistic_calc` repo.

Current repo state
- Epic 1 (measurement foundation) is complete and already pushed to `main`.
- Current pushed tip when this prompt was written: `b956571` (`test: lock down angular spec invariants`).
- `go test ./...` was green at handoff.
- The `units` package now supports all required measurement families and has:
  - angle conversions, including geometric `in/100yd` and `cm/100m`
  - distance conversions
  - velocity conversions
  - energy conversions
  - pressure conversions
  - temperature conversions
  - weight conversions, including the observed Newton-to-grain quirk
  - a full same-unit round-trip matrix test
  - explicit angular invariant tests from Section 12.9

Your job in this session
- Work on Epic 2: Atmosphere model.
- Do NOT start Epic 3 or any trajectory/drag/zero-solver work.

User workflow constraints
- One problem at a time.
- Use strict TDD: write the failing test first, run it and watch it fail, then write minimal code.
- Commit locally after each focused task.
- Do NOT push unless the user explicitly says to push.
- Do not over-polish once core behavior is correct.

Spec areas to use
- `spec.md` Section 6 (Atmosphere Model)
- `spec.md` Section 12.11 (Atmosphere Construction And Humidity Normalization)
- `spec.md` Section 12.12 (ICAO Atmosphere Spot Checks)
- `spec.md` Section 10 required behavioral decisions related to atmosphere updates during flight

If these local planning docs exist, use them as secondary guidance
- `docs/plans/2026-03-29-spec-epic-breakdown.md`
- Epic 2 in that file is the current decomposition
- If those files are absent, proceed from `spec.md` alone

Epic 2 target outcome
- Atmosphere creation, ICAO construction, density factor, speed of sound, and in-flight altitude update behavior are implemented and testable.

Recommended Epic 2 task breakdown
1. Add default atmosphere construction with:
   - altitude `0 ft`
   - pressure `29.92 inHg`
   - temperature `59 F`
   - relative humidity `0.78`
2. Add explicit atmosphere construction with humidity normalization:
   - `[0,1]` means fraction
   - `(1,100]` means percent divided by 100
   - values outside those ranges are rejected
3. Add ICAO-by-altitude atmosphere creation using the formulas from Section 6.3
4. Add density factor and local speed-of-sound formulas from Section 6.4
5. Add altitude-adjusted local-condition helper behavior from Section 6.5
6. Add tests covering default values, normalization, invalid humidity, and ICAO spot checks

What to do first in this session
- Start with Task 1 only.
- Create the smallest atmosphere package/files needed for a default atmosphere constructor and its tests.
- Use TDD.
- Commit that one task locally.
- Then stop and summarize before moving to Task 2.

Suggested initial file layout
- `atmosphere/atmosphere.go`
- `atmosphere/atmosphere_test.go`

Task 1 acceptance target
- There is a default constructor that yields exactly:
  - altitude `0 ft`
  - pressure `29.92 inHg`
  - temperature `59 F`
  - humidity `0.78`
- Tests for that constructor pass.
- `go test ./...` remains green.

Reminder
- Do not push.
- Keep the session narrowly focused on the first atmosphere task.
