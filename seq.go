package iter

import (
	"golang.org/x/exp/constraints"
	"iter"
)

func Keys[K constraints.Ordered, V any](i1 iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range i1 {
			if !yield(k) {
				return
			}
		}
	}
}

func Values[K constraints.Ordered, V any](i1 iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range i1 {
			if !yield(v) {
				return
			}
		}
	}
}

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
