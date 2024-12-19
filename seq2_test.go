package iter

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"maps"
	"reflect"
	"slices"
	"testing"
)

func TestKeySort(t *testing.T) {
	type testCase[K constraints.Ordered, V any] struct {
		name string
		m    map[K]V
		want []V
	}
	tests := []testCase[int, string]{
		{"basic", map[int]string{2: "b", 1: "a", 3: "c"}, []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Collect(Clip(Values(KeySort(tt.m)), 2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleByKey() {
	m := map[int]string{2: "b", 1: "a", 3: "c"}
	keys := slices.Collect(Keys(KeySort(m)))
	fmt.Print(keys)
	// Output: [1 2 3]
}

func TestTypeCast2(t *testing.T) {
	m := map[string]string{"a": "A", "b": "B", "c": "C"}
	w := map[string]any{"a": "A", "b": "B"}

	got := maps.Collect(Clip2(Clip2(Cast2[string, string, any](KeySort(m)), 5), 2))

	if !reflect.DeepEqual(got, w) {
		t.Errorf("Cast2() = %v, want %v", got, w)
	}
}

func TestMatchKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	iter := maps.All(m)

	type testCase struct {
		name string
		keys []string
		want map[string]int
	}
	tests := []testCase{
		{"filter 0", []string{"e", "h"}, map[string]int{}},
		{"filter 2", []string{"a", "c"}, map[string]int{"a": 1, "c": 3}},
		{"filter 1", []string{"d"}, map[string]int{"d": 4}},
		{"filter all", []string{"a", "b", "c", "d"}, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maps.Collect(MatchKeys(iter, tt.keys))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	iter := maps.All(m)

	type testCase struct {
		name   string
		values []int
		want   map[string]int
	}
	tests := []testCase{
		{"filter 0", []int{8, 9}, map[string]int{}},
		{"filter 2", []int{1, 3}, map[string]int{"a": 1, "c": 3}},
		{"filter 1", []int{4}, map[string]int{"d": 4}},
		{"filter all", []int{1, 2, 3, 4}, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maps.Collect(MatchValues(iter, tt.values))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
