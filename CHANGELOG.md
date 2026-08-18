# Changelog

All notable changes to this project will be documented in this file.

## v0.1.0 - Unreleased

Initial pre-1.0 release of candy as a Go utility library.

### Added

- Generic slice helpers in `arrays` for filtering, mapping, reducing, chunking, grouping, and deduplication.
- Explicit-presence Optional values in `optional`.
- Concurrent-safe LRU cache in `cache`.
- Context-aware task groups and batch helpers in `concurrent`.
- String builder and generic joiners in `strs`.
- Date/time boundary, parsing, and formatting helpers in `dates`.
- Adaptive JSON bool and number types in `encoding/jsontypes`.
- Bit flag and bitmap helpers in `types`.
- Eager, slice-backed sequential stream helpers in `stream`, including package-level `Map`, `FlatMap`, and `SortedBy`.
- GoDoc package documentation and testable examples for stable utility packages.

### Changed

- Stabilized selected `collection` and `iter` behavior, including HashSet iteration and ArrayList mutation handling.
- Clarified unsupported or abstract collection/stream APIs with explicit panic messages instead of placeholder `implement me` markers.
- GitHub Actions now runs `go vet ./...` and race-enabled tests across the configured Go/platform matrix.

### Notes

- This is a pre-1.0 release. APIs may evolve before v1.0.0.
- The earlier Java-style `collection` package remains available but should be considered experimental compared with the newer utility packages.
