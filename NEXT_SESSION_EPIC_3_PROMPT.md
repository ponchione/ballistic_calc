Handle Epic 3 for the `ballistic_calc` repo.

Current repo state
- Epic 1 (measurement foundation) is complete and already pushed to `main`.
- Epic 2 (atmosphere model) is complete and already pushed to `main`.
- Current pushed tip when this prompt was written: `e9d1fa1` (`feat: add atmosphere local condition helper`).
- `go test ./...` was green at handoff.
- The `atmosphere` package now supports:
  - default atmosphere construction
  - explicit atmosphere construction with humidity normalization/validation
  - ICAO-by-altitude construction
  - density factor
  - local speed of sound
  - in-flight altitude-adjusted local-condition helper behavior
- Files added in Epic 2:
  - `atmosphere/atmosphere.go`
  - `atmosphere/icao.go`
  - `atmosphere/local.go`
  - `atmosphere/atmosphere_test.go`

Your job in this session
- Work on Epic 3: Drag subsystem.
- Do NOT start Epic 4, zero-solver work, or trajectory-engine work.
- Do NOT revisit Epic 2 polish unless an Epic 3 test proves a real dependency bug.

User workflow constraints
- Before implementation, make a short plan for Epic 3 and wait for user approval.
- One problem at a time.
- Use strict TDD: write the failing test first, run it and watch it fail, then write minimal code.
- Commit locally after each focused task.
- Do NOT push unless the user explicitly says to push.
- Do not over-polish once core behavior is correct.

Spec areas to use
- `spec.md` Section 7 (Drag Model)
- `spec.md` Section 10 drag-related required behavioral decisions
- `spec.md` Scenario 12.5
- `spec.md` Scenario 12.10
- `spec.md` Scenario 12.13

If these local planning docs exist, use them as secondary guidance
- `docs/plans/2026-03-29-spec-epic-breakdown.md`
- Epic 3 in that file is the current decomposition
- If those files are absent, proceed from `spec.md` alone

Epic 3 target outcome
- Standard tables, BC vs FF handling, PIR scaling, custom raw functions, and curve interpolation all exist behind one coherent drag interface.

Recommended Epic 3 task breakdown
1. Define the external drag modes: standard table, custom function, and custom curve helper path.
2. Add strict validation for supported identifiers: `G1`, `G2`, `G5`, `G6`, `G7`, `G8`, `GS`, `RA4`.
3. Load independently sourced standard drag reference data and document provenance.
4. Implement raw drag lookup and PIR scaling (`2.08551e-04`).
5. Implement BC mode and FF mode, with FF rejecting missing or nonpositive diameter.
6. Implement custom curve precomputation and evaluation with the documented edge behavior.
7. Add spot-check tests for every standard table at Mach `0.5`, `1.0`, and `2.0`.
8. Add tests for custom drag-function scaling and FF-to-BC conversion.

What to do first in this session
- First, read `spec.md` Section 7 and Epic 3 in `docs/plans/2026-03-29-spec-epic-breakdown.md`.
- Then produce a short implementation plan for Epic 3 and stop for user approval.
- After approval, start with Task 1 only.
- Create the smallest `drag` package/files needed for the public drag-mode surface and its tests.
- Use TDD.
- Commit that one task locally.
- Then stop and summarize before moving to Task 2.

Suggested initial file layout
- `drag/types.go`
- `drag/drag_test.go`

Task 1 acceptance target
- There is a minimal public representation for:
  - standard drag table mode
  - custom drag-function mode
  - custom curve helper path
- The initial drag package tests pass.
- `go test ./...` remains green.
- Do not pull in standard-table data, PIR scaling, or interpolation math yet unless a Task 1 test genuinely requires it.

Reminder
- Do not push.
- Keep the session narrowly focused on planning first, then only the first Epic 3 slice after approval.
