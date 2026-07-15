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

## G1/G7 direct-source investigation

Investigation date: 2026-07-15.

### Engineering decision

No candidate examined in this investigation satisfies both parts of the
Milestone 1 gate: an inspectable direct source containing the exact G1 and G7
reference values, and a documented reuse basis covering the intended source
and binary distribution. The current G1/G7 literals therefore remain
quarantined, and they are not promoted as an authority for replacement data.
No drag implementation or fixture is changed on the strength of this review.

The strongest engineering provenance lead is the Sporting Arms and Ammunition
Manufacturers' Institute (SAAMI) publication *Tables of Siacci Functions and
Drag Coefficients*, dated June 1976. Secondary implementation manuals identify
that publication as their source for G1 and G7 and explicitly say their use was
by permission. An official copy, its exact table locations, and permission
applicable to this project were not found. That lead is not release clearance.

The records below distinguish a source candidate from a rights conclusion.
Government publication, government hosting, scientific content, public access,
or a statement that distribution of a report is unlimited does not by itself
settle copyright or all other reuse questions. Final evaluation remains an
owner/legal decision, not an engineering declaration or legal opinion.

### `CAND-DRAG-SAAMI-1976`: paired G1/G7 lead, blocked

| Field | Evidence |
| --- | --- |
| Author or responsible body | Sporting Arms and Ammunition Manufacturers' Institute (SAAMI); no individual compiler was identified. SAAMI is a private standards organization, so the requested government-employee/contractor classification is **unknown/not governmental** for this artifact. |
| Publication | *Tables of Siacci Functions and Drag Coefficients*, June 1976. No report number, edition statement, or later revision identifier was found. A related title reported by the same secondary records is *SAAMI Exterior Ballistics Centerfire and Rimfire Ammunition*, June 1976. |
| Exact locator | Not established. No official or otherwise rights-documented copy of the 1976 tables was located, so page, table, column definitions, normalization, and source precision remain unverified. |
| URL and retrieval | No direct official digital artifact was found. Current official rights context: <https://saami.org/terms-of-use/> and <https://saami.org/technical-information/recoil-formulae/>, retrieved 2026-07-15. Secondary provenance leads only: Oehler *System 89 User Manual*, PDF page 51 of 72, <https://oehler-research.com/wp-content/uploads/2021/06/Sys89UserManual.pdf>; Dexadine *Ballistic Explorer* help, “Drag Functions,” <https://www.dexadine.com/bexhelp/bexhelp140.htm>, both retrieved 2026-07-15. |
| Rights or reuse statement | No reuse statement for the 1976 publication was found. Current SAAMI terms say site content is owned, licensed, or used by permission and that no license is implied; a current data-sheet page reserves publication and reproduction rights. Those current statements are context, not proof of the 1976 terms. The secondary manuals' “with permission” statements show that other publishers sought permission; they do not transfer permission to this repository. |
| Required values | Reported by the two secondary provenance leads to contain the standard G1 and G7 material, but the actual source values were not inspected. This is therefore a promising engineering candidate, not verified numerical authority. |
| Reproducible transformation | If a source and permission are obtained: preserve and checksum the source artifact; transcribe only G1/G7 rows and all column definitions; determine whether the source contains Mach/drag pairs or Siacci functions; record the atmosphere, velocity-to-Mach convention, reference area, normalization, and rounding; generate Go data from an intermediate machine-readable file; and independently compare source points and interpolated values. The transformation cannot be specified completely before the source is inspected. |
| Checksum or archive ID | None available for the unavailable primary artifact. The secondary Oehler and Dexadine locations are not substitutes for a checksum of the SAAMI source. |
| Remaining uncertainty | Obtain a complete copy and written SAAMI response covering machine-readable transcription or derivation, modification/interpolation, public source and binary redistribution, attribution/notice, and any duration or territory limits. Owner/legal review must then decide whether the response and artifact are compatible with the intended release. |

### `CAND-DRAG-INGALLS-1918`: historical G1 lineage only

| Field | Evidence |
| --- | --- |
| Author or responsible body | Colonel James M. Ingalls, U.S. Army; revised under the direction of the U.S. Army Ordnance Board and published by the U.S. Government Printing Office. This is **mixed/uncertain** for rights analysis: an identified Army officer, government revision and publication, and no investigation here of Ingalls' employment capacity for every original or revised element. |
| Publication | James M. Ingalls, *Ingalls' Ballistic Tables*, computed 1893, revised edition, Artillery Circular “M” (Revised), Washington: Government Printing Office, 1918; xxviii + 233 pages. Stable catalog identifiers include HathiTrust record `001621998`, University of Michigan item `mdp.39015076691438`, LCCN `18026814`, and OCLC `782794`. |
| Exact locator | The volume's primary tabulations are integrated Siacci/ballistic functions rather than a modern paired G1 Mach/drag-coefficient dataset. No G7 table is present. Exact rows suitable for a modern G1 reconstruction were therefore not selected. |
| URL and retrieval | HathiTrust catalog and full-view item: <https://catalog.hathitrust.org/Record/001621998> and <https://babel.hathitrust.org/cgi/pt?id=mdp.39015076691438>, retrieved 2026-07-15. Google Books catalog: <https://books.google.com/books?id=_HgDAAAAYAAJ>, retrieved 2026-07-15. |
| Rights or reuse statement | Full-view access and government publication metadata were found, but no artifact-specific license or reviewed government-work determination was found. This review does not infer public-domain status from the publisher, age, or host. |
| Required values | Does not provide the required paired G1 and G7 Mach/drag reference values. It is a historical G1-lineage source only. |
| Reproducible transformation | A G1-only reconstruction would require selecting the correct historical function, differentiating or otherwise inverting integrated functions, establishing atmosphere and speed-of-sound conventions, converting velocity to Mach, mapping the historical projectile/normalization to modern G1, and quantifying numerical error. That is not an acceptable transformation for the paired v0.1 source gate without a direct modern specification and independent validation. |
| Checksum or archive ID | Stable catalog/item IDs above; no source-file checksum recorded. |
| Remaining uncertainty | Even if rights were reviewed, absence of G7 and absence of direct Mach/drag pairs make this artifact insufficient for the requested implementation. |

### `CAND-DRAG-LOWRY-1965`: private G1-only lead

| Field | Evidence |
| --- | --- |
| Author or responsible body | E. D. Lowry; Olin Mathieson Chemical Corporation, Research Department, Winchester-Western Division. The identified organization is private; individual employment and any government sponsorship are **unknown**, and no government authorship was identified. |
| Publication | E. D. Lowry, *Exterior Ballistics of Small Arms Projectiles*, Olin Mathieson Chemical Corporation, Winchester-Western Division, 1965, 620 pages. No report number or edition statement was established. |
| Exact locator | Not established from an inspectable direct copy. Secondary bibliographic descriptions point to G1 material but not G7; no exact table/page was adopted. |
| URL and retrieval | Google Books catalog record: <https://books.google.com/books?id=iEEaHAAACAAJ>, retrieved 2026-07-15. No publisher-hosted digital artifact was found. |
| Rights or reuse statement | No explicit license or reuse permission was found. Corporate authorship/publication is not a reuse basis. |
| Required values | Reported as a G1 source, not a paired G1/G7 source. The actual numeric table was not inspected. |
| Reproducible transformation | Unknown until a direct copy, column definitions, and rights basis are obtained. It cannot supply G7 in any event. |
| Checksum or archive ID | Google Books stable volume ID `iEEaHAAACAAJ`; no source checksum. |
| Remaining uncertainty | Exact contents and rights would require publisher/rightsholder inquiry, but the known G7 omission already disqualifies it as the paired v0.1 source. |

### `CAND-DRAG-MCCOY-1999`: integrated functions, no reuse grant

| Field | Evidence |
| --- | --- |
| Author or responsible body | Robert L. McCoy; privately published by Schiffer Publishing. McCoy's government career does not establish that this book was written as a government work; the relevant authorship classification is **unknown/private publication**. |
| Publication | Robert L. McCoy, *Modern Exterior Ballistics: The Launch and Flight Dynamics of Symmetric Projectiles*, Schiffer Publishing, first edition 1999, ISBN `978-0-7643-0720-1`. A second edition was published in 2012, ISBN `978-0-7643-3825-0`; page evidence below applies to the 1999 edition. |
| Exact locator | “Tables of the Primary Siacci Functions,” pages 114–156; G1 pages 119–124 and G7 pages 140–145. The order is velocity, space function, altitude function, trajectory-inclination function, and time-of-flight function. These are integrated Siacci functions, not an established direct table of Mach/drag-coefficient pairs. The locators are corroborated by the 2010 errata, not by a licensed source copy inspected for this project. |
| URL and retrieval | Google Books catalog: <https://books.google.com/books?id=ay5UPQAACAAJ>. Donald G. Miller et al., *Typographical Errors in Robert L. McCoy, “Modern Exterior Ballistics”*, revision dated 2010-04-13, page 7 of the 20-page PDF, <https://www.dexadine.com/download/mccoy10.pdf>. Retrieved 2026-07-15. |
| Rights or reuse statement | No license or permission allowing republication of the book's tables was found. A commercial publication and accessible errata do not supply a reuse basis for the underlying tables. |
| Required values | Contains G1/G7 integrated functions according to the errata, but no direct raw Mach/drag reference dataset was established. |
| Reproducible transformation | Reconstruction would require accurately recovering the tabular functions, resolving units and atmospheric conventions, numerically differentiating/inverting them into drag, converting velocity to Mach, and proving equivalence and error bounds. That introduces avoidable derivation and transcription risk and remains blocked by rights. |
| Checksum or archive ID | Errata SHA-256 `8093227109d785b7322ace94e90079742a633c81eb02f93ac4af9a6eb466f1bf`; Google Books volume ID `ay5UPQAACAAJ`. No checksum of the book was recorded. |
| Remaining uncertainty | Obtain publisher/rightsholder permission and an authorized edition before using any book table. Even with permission, prefer a direct paired Mach/drag source over reconstruction from integrated functions. |

### `CAND-DRAG-MCCOY-1981`: official report, wrong dataset

| Field | Evidence |
| --- | --- |
| Author or responsible body | Robert L. McCoy; U.S. Army Armament Research and Development Command, U.S. Army Ballistic Research Laboratory. The report identifies the government laboratory as performing organization and leaves the contract/grant field blank, but does not expressly state the author's employment status; classification is **government organization, individual status not expressly stated**. |
| Publication | Robert L. McCoy, *“MC DRAG” — A Computer Program for Estimating the Drag Coefficients of Projectiles*, Technical Report `ARBRL-TR-02293`, February 1981; DTIC accession `ADA098110`. |
| Exact locator | Report-documentation page and abstract; user's guide, printed pages 24–31; example projectile curves, Figures 14–24 on printed pages 45–55; program listing in the appendix beginning on printed page 63. A full-text check found no G1 or G7 table. |
| URL and retrieval | DTIC citation/PDF: <https://apps.dtic.mil/sti/citations/ADA098110> and <https://apps.dtic.mil/sti/tr/pdf/ADA098110.pdf>. Stable archival item: <https://archive.org/details/DTIC_ADA098110>. Retrieved 2026-07-15; the DTIC PDF endpoint was rate-limited during retrieval, so the checksum is of the archival scan. |
| Rights or reuse statement | Title page and report-documentation block 16 state “Approved for public release; distribution unlimited.” That is a report-distribution statement, not an explicit software/data license or a complete copyright analysis. No further reuse license was found. |
| Required values | Does not contain G1/G7 reference values. It estimates zero-yaw drag for user-supplied projectile geometry over Mach 0.5–5 and presents curves for example projectiles. |
| Reproducible transformation | Running or reimplementing the included estimator for assumed G1/G7 geometry would create modeled curves, not reproduce a documented standard reference dataset. No such transformation is accepted. |
| Checksum or archive ID | Internet Archive item `DTIC_ADA098110`; retrieved PDF SHA-256 `98e124f005fd9dfe18463e1acfb94799030619fd2961bf1da704767a31c0ae12`. |
| Remaining uncertainty | Rights review would still be needed for reuse of report material, but numerical unsuitability is decisive for this gate. |

### `CAND-DRAG-DICKINSON-1967`: official false lead

| Field | Evidence |
| --- | --- |
| Author or responsible body | Elizabeth R. Dickinson; U.S. Army Ballistic Research Laboratories, Computing Laboratory. No contractor is identified; classification is **government laboratory, individual status not expressly stated**. |
| Publication | Elizabeth R. Dickinson, *The Production of Firing Tables for Cannon Artillery*, Ballistic Research Laboratories Report `1371`, November 1967; DTIC accession `AD0826735`. |
| Exact locator | Contents, printed page 5, lists “Table G1 Supplementary Data” at printed page 62. In context this is firing-table section G1, between Tables F and H, not the G1 standard drag model. The report's drag-coefficient discussion begins on printed page 23. A full-text check found no G7 reference table. |
| URL and retrieval | DTIC citation/PDF: <https://apps.dtic.mil/sti/citations/AD0826735> and <https://apps.dtic.mil/sti/tr/pdf/AD0826735.pdf>. Stable archival item: <https://archive.org/details/DTIC_AD0826735>. Retrieved 2026-07-15; the DTIC PDF endpoint was rate-limited during retrieval, so the checksum is of the archival scan. |
| Rights or reuse statement | The scan preserves the report's original 1967 export-control/distribution markings and a later limitation-change page dated 1981-02-20 changing distribution to “Approved for public release; distribution unlimited.” This review records both and does not treat the later distribution statement as a general reuse license. |
| Required values | Does not contain paired G1/G7 standard drag reference values. The apparent “G1” hit is a firing-table label. |
| Reproducible transformation | None; the report explains firing-table production and cannot be transformed into the missing standard curves. |
| Checksum or archive ID | Internet Archive item `DTIC_AD0826735`; retrieved PDF SHA-256 `daaf731d4299603178ea988085d8f5ec77618bd8636f2eddef539bda253de36c`. |
| Remaining uncertainty | None material to the numerical decision. Any future reuse of report content would still require an artifact-specific rights review. |

### Required next source and permission action

The shortest credible path is for the repository owner to request from SAAMI:

1. an authoritative complete copy of the June 1976 *Tables of Siacci Functions
   and Drag Coefficients*, including title/edition pages and G1/G7 column notes;
2. written permission, or a specific published rights statement, covering a
   machine-readable transcription or derived G1/G7 dataset, modification and
   interpolation, and redistribution in public source and compiled binaries;
3. required attribution, notices, source links, and any scope, duration, or
   territory limits; and
4. confirmation of who owns or can authorize the underlying compilation.

The owner and legal reviewer must decide whether that response supports the
intended release. If it does, engineering can preserve and checksum the source,
document a deterministic transformation, generate only the G1/G7 data, and
start the focused drag implementation. An alternative is an official direct
artifact that contains the exact paired numeric curves and comes with equally
specific authorship and reuse evidence. Until one of those paths succeeds,
Milestone 0 remains active and Milestone 1 does not start.

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
