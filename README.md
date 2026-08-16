# candy

Go utility packages for common application code.

## Packages

- arrays: generic slice filtering, mapping, reducing, chunking, grouping, and deduplication.
- optional: explicit-present Optional values.
- cache: concurrent-safe LRU cache.
- concurrent: context-aware task groups and batch helpers.
- strs: chainable string builder and generic joiners.
- dates: date/time boundaries, parsing, and formatting.
- encoding/jsontypes: adaptive JSON bool and number types.
- types: lightweight type helpers such as bit flags and bitmaps.

The repository also contains earlier experimental Java-style collection and stream packages. New utility packages are designed to be independent and practical for normal Go code.

## Examples

```go
start := dates.BeginOfDay(time.Date(2024, 5, 6, 15, 4, 5, 0, time.UTC))
fmt.Println(dates.FormatDate(start)) // 2024-05-06
```

```go
var payload struct {
	Active jsontypes.AdaptiveBool   `json:"active"`
	Count  jsontypes.AdaptiveNumber `json:"count"`
}
_ = json.Unmarshal([]byte(`{"active":"1","count":"42"}`), &payload)
fmt.Println(payload.Active.Bool(), payload.Count.Int64()) // true 42
```

```go
var bitmap types.Bitmap
bitmap.Set(1)
bitmap.Set(3)
fmt.Println(bitmap.Get(1), bitmap.String()) // true b:0101
```

## Requirements

Go 1.18 or newer.

## Test

```sh
go test ./...
```
