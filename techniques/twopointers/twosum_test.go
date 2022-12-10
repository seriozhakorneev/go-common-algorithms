package twopointers

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums, expectedResult []int
		target               int
	}{
		{
			nums:           []int{2, 7, 11, 15},
			expectedResult: []int{1, 2},
			target:         9,
		},
		{
			nums:           []int{2, 3, 4},
			expectedResult: []int{1, 3},
			target:         6,
		},
		{
			nums:           []int{-1, 0},
			expectedResult: []int{1, 2},
			target:         -1,
		},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Fatalf("Expected result: %v, Got: %v", test.expectedResult, result)
		}
	}

}
