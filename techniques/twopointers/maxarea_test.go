package twopointers

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	expectedMax := []int{49, 1}

	for i := range height {
		max := maxArea(height[i])

		if max != expectedMax[i] {
			t.Fatalf("expected max: %v, got: %v", expectedMax[i], max)
		}
	}

}
