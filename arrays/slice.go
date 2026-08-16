package arrays

// Filter returns a new slice containing items for which keep returns true.
// A nil input slice is preserved as nil.
func Filter[T any](items []T, keep func(T) bool) []T {
	if items == nil {
		return nil
	}

	filtered := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// FilterInPlace compacts items into its existing backing array, keeping items
// for which keep returns true.
func FilterInPlace[T any](items []T, keep func(T) bool) []T {
	filtered := items[:0]
	for _, item := range items {
		if keep(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// Reject returns a new slice containing items for which reject returns false.
// A nil input slice is preserved as nil.
func Reject[T any](items []T, reject func(T) bool) []T {
	return Filter(items, func(item T) bool { return !reject(item) })
}

// RejectInPlace compacts items into its existing backing array, keeping items
// for which reject returns false.
func RejectInPlace[T any](items []T, reject func(T) bool) []T {
	return FilterInPlace(items, func(item T) bool { return !reject(item) })
}

// RemoveZero returns a new slice containing all non-zero values from items.
// A nil input slice is preserved as nil.
func RemoveZero[T comparable](items []T) []T {
	var zero T
	return Filter(items, func(item T) bool { return item != zero })
}

// RemoveZeroInPlace compacts items into its existing backing array, keeping all
// non-zero values.
func RemoveZeroInPlace[T comparable](items []T) []T {
	var zero T
	return FilterInPlace(items, func(item T) bool { return item != zero })
}

// Map returns a new slice containing mapper applied to each item.
func Map[T any, U any](items []T, mapper func(T) U) []U {
	if items == nil {
		return nil
	}

	mapped := make([]U, len(items))
	for i, item := range items {
		mapped[i] = mapper(item)
	}
	return mapped
}

// FlatMap returns a new slice containing the concatenated mapper results for
// each item.
func FlatMap[T any, U any](items []T, mapper func(T) []U) []U {
	if items == nil {
		return nil
	}

	mapped := make([]U, 0, len(items))
	for _, item := range items {
		mapped = append(mapped, mapper(item)...)
	}
	return mapped
}

// Reduce folds items from left to right, starting with initial.
func Reduce[T any, U any](items []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, item := range items {
		result = reducer(result, item)
	}
	return result
}

// Contains reports whether target is present in items.
func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

// Unique returns a new slice with duplicate items removed, preserving first-seen
// order. A nil input slice is preserved as nil.
func Unique[T comparable](items []T) []T {
	if items == nil {
		return nil
	}

	seen := make(map[T]struct{}, len(items))
	unique := make([]T, 0, len(items))
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		unique = append(unique, item)
	}
	return unique
}

// Chunk splits items into chunks with at most size items each. It panics when
// size is less than or equal to zero. A nil input slice is preserved as nil.
func Chunk[T any](items []T, size int) [][]T {
	if size <= 0 {
		panic("arrays: chunk size must be positive")
	}
	if items == nil {
		return nil
	}

	chunks := make([][]T, 0, (len(items)+size-1)/size)
	for start := 0; start < len(items); start += size {
		end := start + size
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[start:end])
	}
	return chunks
}

// GroupBy groups items by the key returned from keyFunc.
func GroupBy[T any, K comparable](items []T, keyFunc func(T) K) map[K][]T {
	groups := make(map[K][]T)
	for _, item := range items {
		key := keyFunc(item)
		groups[key] = append(groups[key], item)
	}
	return groups
}

// ToMap builds a map from items using keyFunc for keys and valueFunc for values.
// Later items with the same key overwrite earlier items.
func ToMap[T any, K comparable, V any](items []T, keyFunc func(T) K, valueFunc func(T) V) map[K]V {
	mapped := make(map[K]V, len(items))
	for _, item := range items {
		mapped[keyFunc(item)] = valueFunc(item)
	}
	return mapped
}
