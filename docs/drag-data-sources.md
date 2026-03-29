# Drag data sources

This repo vendors the standard drag reference curves from the public JBM Ballistics downloads catalog, not from the original `ballistic_calc` source tree.

Scope of the vendored data:
- Raw `Cd` versus Mach reference points only.
- No PIR scaling is baked into these tables.
- Runtime interpolation/evaluation behavior is implemented separately in the `drag` package.

Primary catalog page:
- https://www.jbmballistics.com/ballistics/downloads/downloads.shtml

Why this source set:
- The JBM downloads page exposes a single public catalog covering all advertised standard tables needed by this repo: `G1`, `G2`, `G5`, `G6`, `G7`, `G8`, `GS`, and `RA4`.
- On that page, under "More Drag Functions", JBM states that the G-functions and the sphere table were obtained from BRL.
- The same page notes that the original source of the hosted `RA4` function is unknown. We keep that limitation explicit instead of overstating provenance.
- The page also links a separate Winchester/Olin text for `G1`, `G5`, `G6`, and `GL`, but the `mcg*`/`mcgs`/`ra4` set was chosen because it consistently covers the full catalog required by the spec.

Loaded files:

| Table | Direct source URL | Points | Provenance note |
| --- | --- | ---: | --- |
| `G1` | https://www.jbmballistics.com/ballistics/downloads/text/mcg1.txt | 79 | JBM Ballistics downloads page ('More Drag Functions') says the G functions were obtained from BRL. |
| `G2` | https://www.jbmballistics.com/ballistics/downloads/text/mcg2.txt | 85 | JBM Ballistics downloads page ('More Drag Functions') says the G functions were obtained from BRL. |
| `G5` | https://www.jbmballistics.com/ballistics/downloads/text/mcg5.txt | 76 | JBM Ballistics downloads page ('More Drag Functions') says the G functions were obtained from BRL. |
| `G6` | https://www.jbmballistics.com/ballistics/downloads/text/mcg6.txt | 79 | JBM Ballistics downloads page ('More Drag Functions') says the G functions were obtained from BRL. |
| `G7` | https://www.jbmballistics.com/ballistics/downloads/text/mcg7.txt | 84 | JBM Ballistics downloads page ('More Drag Functions') says the G functions were obtained from BRL. |
| `G8` | https://www.jbmballistics.com/ballistics/downloads/text/mcg8.txt | 78 | JBM Ballistics downloads page ('More Drag Functions') says the G functions were obtained from BRL. |
| `GS` | https://www.jbmballistics.com/ballistics/downloads/text/mcgs.txt | 81 | JBM Ballistics downloads page ('More Drag Functions') says the 9/16" sphere table was obtained from BRL. |
| `RA4` | https://www.jbmballistics.com/ballistics/downloads/text/ra4.txt | 87 | JBM Ballistics downloads page ('More Drag Functions') hosts the RA4 table but explicitly says the original source is unknown. |

Implementation note:
- `drag/standard_data.go` is a literal transcription of the public JBM-hosted text tables into Go `CurvePoint` slices for deterministic builds and offline tests.
- Future validation/spot-check tests should cite this document and the JBM URLs above.
