# Provenance and licensing

## Current posture

Original code created specifically for `ballistic_calc` is intended for
eventual release under `0BSD`. That is an intent, not the current license state
of the repository. There is no repository-wide `LICENSE`, and the complete tree
has not been cleared as `0BSD`.

Code licensing and non-original material are separate questions. A permissive
license applied to original code does not grant rights in imported tables,
fixtures, publications, diagrams, or documentation. Attribution and data terms
remain attached to their own source records.

The current standard-drag literals are quarantined for public release. They
were transcribed from JBM-hosted text files. The JBM catalog describes most
curves as obtained from the U.S. Army Ballistics Research Laboratory and says
the original source of RA4 is unknown, but a secondary site's statement and
public download do not by themselves establish a distribution right for this
project. Retrieval dates, direct publication identifiers, authorship, and a
rights basis were not recorded when the literals were added.

As a result:

- G1 and G7 need replacement from cleared direct sources or permission
  sufficient for the intended distribution before v0.1 ships.
- G2, G5, G6, G8, GS, and RA4 are deferred product scope and are excluded from
  a v0.1 release artifact unless separately cleared and deliberately added to a
  later product contract.
- The presence of data in Git or a passing test does not clear it for release.

This document records engineering release gates. It is not legal advice and
does not conclude that the repository is free of every copyright, patent,
trademark, contract, export, or other restriction.

## Current non-original material inventory

### Standard drag tables

All rows below are embedded in `drag/standard_data.go`. The transformation was
a literal transcription of Mach/drag text pairs into Go `CurvePoint` slices;
PIR scaling is applied separately at runtime. The recorded source is secondary,
the original publication and retrieval date are missing, and release status is
**quarantined**.

| ID | Table | Recorded JBM-hosted source | Points | Direct-source status |
| --- | --- | --- | ---: | --- |
| `DATA-DRAG-G1` | G1 | <https://www.jbmballistics.com/ballistics/downloads/text/mcg1.txt> | 79 | BRL origin asserted by catalog; publication and rights basis not recorded |
| `DATA-DRAG-G2` | G2 | <https://www.jbmballistics.com/ballistics/downloads/text/mcg2.txt> | 85 | BRL origin asserted by catalog; publication and rights basis not recorded |
| `DATA-DRAG-G5` | G5 | <https://www.jbmballistics.com/ballistics/downloads/text/mcg5.txt> | 76 | BRL origin asserted by catalog; publication and rights basis not recorded |
| `DATA-DRAG-G6` | G6 | <https://www.jbmballistics.com/ballistics/downloads/text/mcg6.txt> | 79 | BRL origin asserted by catalog; publication and rights basis not recorded |
| `DATA-DRAG-G7` | G7 | <https://www.jbmballistics.com/ballistics/downloads/text/mcg7.txt> | 84 | BRL origin asserted by catalog; publication and rights basis not recorded |
| `DATA-DRAG-G8` | G8 | <https://www.jbmballistics.com/ballistics/downloads/text/mcg8.txt> | 78 | BRL origin asserted by catalog; publication and rights basis not recorded |
| `DATA-DRAG-GS` | GS | <https://www.jbmballistics.com/ballistics/downloads/text/mcgs.txt> | 81 | BRL sphere-table origin asserted by catalog; publication and rights basis not recorded |
| `DATA-DRAG-RA4` | RA4 | <https://www.jbmballistics.com/ballistics/downloads/text/ra4.txt> | 87 | Catalog says original source is unknown |

Recorded catalog context:
<https://www.jbmballistics.com/ballistics/downloads/downloads.shtml>.

The source URLs and descriptive notes embedded in Go are provenance clues, not
clearance records.

### Formula and constant groups

| ID | Material | Current location | Status before release |
| --- | --- | --- | --- |
| `FORM-ATM-ICAO` | Temperature lapse and pressure exponent used by ICAO-style constructors and altitude updates | `atmosphere/icao.go`, `atmosphere/local.go` | Direct standard/report and equation mapping not recorded; independently source and verify |
| `FORM-ATM-DENSITY` | Humidity correction, density factor, and speed-of-sound polynomial/constants | `atmosphere/atmosphere.go` | Direct scientific source not recorded; independently source, review model limits, and verify |
| `FORM-DRAG-PIR` | `2.08551e-04` raw-drag scaling constant | `drag/standard_tables.go` | Meaning and direct technical source not recorded; establish derivation or replace model formulation |
| `CONST-UNITS` | Unit conversion constants and geometric angle formulas | `units/*.go` | Independently check against authoritative standards; specifically resolve the anomalous Newton-to-grain value |

Mathematical methods and facts can still require accurate sourcing, careful
expression, and patent or contract review. A formula being scientific or
widely known does not make every transcription, table, explanation, or software
implementation automatically unrestricted.

### Fixtures and expected values

| ID | Material | Current location | Status before release |
| --- | --- | --- | --- |
| `FIX-UNITS` | Conversion examples, round trips, and angle expectations | `units/*_test.go` | Re-anchor cross-unit values to recorded independent standards; round trips remain internal invariants |
| `FIX-ATM` | Default/explicit atmosphere, ICAO spots, density, sound speed, and altitude-update expectations | `atmosphere/atmosphere_test.go` | Regression-only until each group has an independent reference or derivation record |
| `FIX-DRAG-DATA` | Table ordering, presence, and distinctness | `drag/standard_data_test.go` | Integrity-only; inherits the table quarantine |
| `FIX-DRAG-EVAL` | Exact G1 point and PIR scaling expectations, plus custom-function scaffolding | `drag/drag_test.go` | Regression-only; G1 inherits quarantine and PIR needs a direct source |

There are currently no zero-solver or trajectory fixtures. New reference
fixtures follow the evidence record in [Verification strategy](verification.md)
from their first commit.

### Product and documentation context

| ID | Material | Use | Status |
| --- | --- | --- | --- |
| `REQ-PRODUCT-2026-07-15` | Product direction supplied by the repository owner for the documentation reset | Defines the independent engine, v0.1 scope, integration, safety, and license intent | Internal product requirement; no scientific or rights conclusion |
| `CTX-KICKBRASS-2026-07-15` | KickBrass Bench exploration and architecture documents | Establishes backend, web, privacy, ownership, and deferred-ADR context | Summarized for integration; no KickBrass code or substantial document text copied |

The active documentation is original project documentation assembled from
those requirements and repository inspection. Scientific claims and reference
data still require the separate source records described below.

## Source record schema

Each non-original or reference-backed asset gets a durable record containing:

| Field | Recorded information |
| --- | --- |
| Record ID | Stable identifier used by code, data, tests, and release evidence |
| Asset | Table, fixture, formula, excerpt, diagram, or other material covered |
| Direct source | Original artifact URL or archive location, not only a secondary transcription |
| Creator | Author, agency, standards body, contractor, or publisher |
| Publication | Full title, report/standard number, edition, and publication date |
| Locator | Page, table, figure, equation, appendix, or dataset row range |
| Retrieval | Retrieval date, archive URL if available, and source checksum |
| Rights basis | License, permission, government-work analysis, or other reviewed basis, with scope and reviewer |
| Transformation | Transcription, digitization, unit conversion, rounding, interpolation, filtering, or code generation |
| Derived artifact | Repository path, generation command, checksum, and reproducibility notes |
| Verification | Independent checker, tests, tolerances, and review date |
| Attribution | Required credit, notice text, link, or citation placement |
| Status | Proposed, quarantined, cleared for a stated use, replaced, or excluded |

“Cleared” always names the material and distribution use it covers. It is not a
blanket statement about the repository.

## Release-clearance process

1. Inventory every non-original table, fixture, formula source, substantial
   documentation source, and generated artifact.
2. Complete the source record with a direct source, authorship, publication
   identifier, retrieval details, transformation, and reviewed rights basis.
3. Prefer traceable U.S. government technical reports or other sources with
   explicit reuse permission over a secondary website transcription.
4. Verify authorship and scope rather than assuming every government-hosted or
   scientific artifact is public domain. Contractor works, foreign rights,
   secondary compilations, patents, trademarks, and contractual terms receive
   separate consideration.
5. Replace the current G1/G7 literals from cleared sources, or record permission
   sufficient for the intended source and binary distribution.
6. Rebuild and checksum derived Go data through a documented, reproducible
   transformation; independently compare source points and interpolation.
7. Remove deferred table data from the v0.1 release artifact unless each table
   is separately cleared and brought into product scope.
8. Keep original-code licensing, third-party data terms, and attribution in
   separate records and notices.
9. Complete legal/provenance review before applying a repository-wide license
   or describing a public release as unrestricted or independently cleared.

Private use or deployment on a hosted backend does not by itself make LGPL
software unusable. The independent implementation and intended `0BSD` posture
are product-ownership and downstream-distribution choices, not a claim about
what the LGPL permits in every circumstance.

## v0.1 license gate

The intended `0BSD` license is applied only after an authorship inventory
identifies what original code can be licensed that way and all included
non-original assets have a compatible, recorded status. A third-party notices
or data-attribution file is added when cleared material requires it.

Until then, the repository remains development work with unresolved licensing
and provenance. Passing numerical tests, replacing a URL, or citing a source
does not complete this gate.
