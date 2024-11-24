package iter

import (
	"golang.org/x/exp/constraints"
	"iter"
	"maps"
	"slices"
)

// KeySort returns an iterator that iterates the given map by ordered keys so that iteration is
// repeatable and predictable.
func KeySort[Map ~map[K]V, K constraints.Ordered, V any](m Map) iter.Seq2[K, V] {
	keys := slices.Sorted(maps.Keys(m))
	return func(yield func(K, V) bool) {
		for _, k := range keys {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

// Cast2 converts the iterator K,V1 to an iterator of K,V2.
// It uses interface type casting to do so, which means either V1 or
// V2 should be an interface that is type castable to the other type.
// It will panic if not.
func Cast2[K constraints.Ordered, V1, V2 any](iter1 iter.Seq2[K, V1]) iter.Seq2[K, V2] {
	return func(yield func(K, V2) bool) {
		for k, v := range iter1 {
			v2 := any(v)
			v3 := v2.(V2)
			if !yield(k, v3) {
				return
			}
		}
	}
}

// Clip2 will stop the iteration after n times.
func Clip2[K, V any](it iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var i int
		for k, v := range it {
			if i >= n {
				return
			}
			if !yield(k, v) {
				return
			}
			i++
		}
	}
}

// MatchKeys will pass only the pairs that have the first item in keys.
func MatchKeys[K comparable, V any, S ~[]K](it iter.Seq2[K, V], keys S) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		// make faster by turning the keys into a set?
		// perhaps only if the length of keys exceeds a certian amount
		for k, v := range it {
			if slices.Contains(keys, k) {
				if !yield(k, v) {
					return
				}
			}
		}
	}

}
