Continue Epic 3 for the `ballistic_calc` repo.

Where we are stopping
- Epic 1 (measurement foundation) is complete and already pushed.
- Epic 2 (atmosphere model) is complete and already pushed.
- Epic 3 is in progress.
- In this session, Epic 3 Tasks 1-3 were completed locally and tested.
- `go test ./...` was green at stop time.

Epic 3 work completed so far
1. Task 1: public drag-mode surface
   - Added `drag/types.go`
   - Added `drag/drag_test.go`
   - Public surface now includes:
     - standard-table drag mode
     - custom raw-function drag mode
     - custom curve helper wrapper path

2. Task 2: supported table validation
   - `drag.NewStandardTable` now validates supported identifiers
   - Supported IDs:
     - `G1`, `G2`, `G5`, `G6`, `G7`, `G8`, `GS`, `RA4`
   - Unknown standard-table identifiers are rejected without panic

3. Task 3: independently sourced standard drag reference data
   - Added `drag/standard_data.go`
   - Added `drag/standard_data_test.go`
   - Added `docs/drag-data-sources.md`
   - Loaded public JBM-hosted reference tables for:
     - `G1`, `G2`, `G5`, `G6`, `G7`, `G8`, `GS`, `RA4`
   - Provenance doc explicitly notes:
     - JBM says the G-functions and sphere table were obtained from BRL
     - JBM hosts `RA4` but says the original source is unknown

Local commits created before this handoff file
- `6a08333` — `feat: add drag package mode surface`
- `981068f` — `feat: validate supported drag tables`
- `ce22451` — `feat: add standard drag reference data`

Current drag-package status
- We have type definitions, constructor surface, identifier validation, and raw reference tables.
- We do NOT yet have:
  - raw drag lookup API from standard tables
  - PIR scaling (`2.08551e-04`)
  - BC vs FF conversion logic
  - custom curve precompute/eval implementation
  - standard-table Mach spot checks
  - custom drag-function scaling tests
  - FF conversion tests

What to read first next time
- `spec.md` Section 7 (Drag Model)
- `spec.md` Section 10 required drag-related behavior
- `spec.md` Scenario 12.5
- `spec.md` Scenario 12.10
- `spec.md` Scenario 12.13
- `docs/plans/2026-03-29-spec-epic-breakdown.md` (Epic 3 section)
- `docs/drag-data-sources.md`

Exactly where to start next time
- Start with Epic 3 Task 4 only.
- Do NOT start Task 5+ yet.
- Task 4 target: implement raw drag lookup for standard tables and the required PIR scaling hook (`2.08551e-04`).
- Keep it test-first and minimal.

Required workflow next time
- One problem at a time.
- Use strict TDD:
  - write the failing test first
  - run it and watch it fail
  - write minimal code
  - rerun tests
- Commit locally after the focused Task 4 slice.
- Do NOT push unless the user explicitly says to push.
- Do not over-polish once the core behavior is correct.

Suggested Task 4 approach
1. Add the smallest failing tests that prove:
   - a standard table can expose raw drag at a requested Mach from loaded reference data
   - scaling by `PIR = 2.08551e-04` is applied only by the drag-evaluation path, not by stored raw table data
2. Add the smallest implementation needed for:
   - standard-table raw lookup at exact Mach values already present in the loaded tables
   - scaled drag evaluation hook/path
3. Run:
   - `go test ./drag -count=1`
   - `go test ./...`
4. Commit that Task 4 slice locally.
5. Stop and summarize before moving to Task 5.

Important scope guard
- Stay inside Epic 3.
- Do NOT start Epic 4, zero-solver work, or trajectory-engine work.
- Do NOT revisit Epic 2 unless a real Epic 3 test proves a dependency bug.
