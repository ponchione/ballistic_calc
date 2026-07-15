# Documentation index

This directory contains the current documentation for `ballistic_calc`. A new
contributor can start with the [root README](../README.md), then use this index
to find the controlling document for a decision.

## Authoritative documents

| Document | Owns |
| --- | --- |
| [v0.1 product contract](product-contract.md) | Normative product requirements, inputs, outputs, validation, scope, and definition of done |
| [Architecture](architecture.md) | Package responsibilities, dependency rules, numerical boundaries, determinism, and versioning |
| [KickBrass integration](kickbrass-integration.md) | Engine/backend/web dependency direction, adapter ownership, stateless use, snapshots, and Bench prefill |
| [Verification strategy](verification.md) | Independent evidence layers, tolerances, and v0.1 verification gates |
| [Provenance and licensing](provenance-and-licensing.md) | Code-license intent, non-original material inventory, data quarantine, and release-clearance process |
| [Safety and limitations](safety-and-limitations.md) | Technical limitations, use boundaries, and reusable safety language |
| [Implementation roadmap](roadmap.md) | Current-to-v0.1 sequence, milestone completion criteria, and post-v0.1 separation |

The [root README](../README.md) is the concise project orientation and current
status. When its summary differs from a focused document, the focused document
owns that subject. Source code and tests remain the authority for behavior that
is already implemented; unfinished exported symbols are not a product promise.

## Supporting and historical material

There are no active supporting plans or archived handoffs in this tree.
Superseded planning files and operational prompts were retired during the
product reset. Git history preserves tracked files for historical
investigation, but they do not guide current implementation.

No document in this repository establishes legal clearance, guarantees
accuracy, or supersedes formal legal and product-safety review.
