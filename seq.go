package iter

import (
	"golang.org/x/exp/constraints"
	"iter"
)

// Keys converts a Seq2 to a Seq, iterating on just the first item in the sequence.
func Keys[K constraints.Ordered, V any](i1 iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range i1 {
			if !yield(k) {
				return
			}
		}
	}
}

// Values converts a Seq2 to a Seq, iterating on just the second item in the sequence.
func Values[K constraints.Ordered, V any](i1 iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range i1 {
			if !yield(v) {
				return
			}
		}
	}
}

// Clip stops the iteration after n times.
func Clip[K any](i1 iter.Seq[K], n int) iter.Seq[K] {
	return func(yield func(K) bool) {
		var i int
		for v := range i1 {
			if i >= n {
				return
			}
			if !yield(v) {
				return
			}
			i++
		}
	}
}
