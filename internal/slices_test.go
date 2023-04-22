package internal

import (
	"testing"
)

func TestSliceAt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		idx      int
		fallback int
		expected int
	}{
		{name: "index within range", slice: []int{1, 2, 3, 4, 5}, idx: 2, fallback: -1, expected: 3},
		{name: "index out of range", slice: []int{1, 2, 3, 4, 5}, idx: 7, fallback: -1, expected: -1},
		{name: "index negative", slice: []int{1, 2, 3, 4, 5}, idx: -1, fallback: -1, expected: -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val := sliceAt(tc.slice, tc.idx, tc.fallback)
			if val != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, val)
			}
		})
	}
}
