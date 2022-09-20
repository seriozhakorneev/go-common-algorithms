package binarysearch

import (
	"testing"
)

func TestSearch(t *testing.T) {

	tests := [][]int{
		{1, 2, 4, 6, 8, 9, 11, 13, 15, 17, 19, 100, 211, 304, 10964},
		{1, 2, 4, 6, 8, 9, 11, 13, 15, 17, 19},
		{1, 7},
		{1},
	}

	for _, test := range tests {
		for expectedIndex, target := range test {

			index := search(test, target)

			if expectedIndex != index {
				t.Fatalf(
					"expected index in test %v with target %v: %v.Got index: %v",
					test,
					target,
					expectedIndex,
					index,
				)
			}
		}
	}
}
