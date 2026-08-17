// Package stream provides sequential, eager, slice-backed streams.
//
// Stream operations preserve encounter order and materialize their results in
// memory. Because Go 1.18 does not support methods with their own type
// parameters, type-changing operations such as Map and FlatMap are
// package-level functions. SortedBy is also provided as a package-level helper
// for sorting values with a less function.
package stream
