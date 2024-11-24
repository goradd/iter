package iter

import (
	"reflect"
	"slices"
	"testing"
)

func TestClip(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	s := slices.Collect(Clip(Keys(KeySort(m)), 2))
	want := []string{"a", "b"}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("Clip() = %v, want %v", s, want)
	}

	s2 := []int{1, 2, 3, 4, 5}
	want2 := []int{1, 2}
	s3 := slices.Collect(Clip(Clip(slices.Values(s2), 5), 2))
	if !reflect.DeepEqual(s3, want2) {
		t.Errorf("Clip() = %v, want %v", s, want)
	}

}
