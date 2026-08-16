# candy

Go utility packages for common application code.

## Packages

- arrays: generic slice filtering, mapping, reducing, chunking, grouping, and deduplication.
- optional: explicit-present Optional values.
- cache: concurrent-safe LRU cache.
- concurrent: context-aware task groups and batch helpers.
- strs: chainable string builder and generic joiners.
- types: lightweight type helpers such as bit flags.

The repository also contains earlier experimental Java-style collection and stream packages. New utility packages are designed to be independent and practical for normal Go code.

## Requirements

Go 1.18 or newer.

## Test

```sh
go test ./...
```
