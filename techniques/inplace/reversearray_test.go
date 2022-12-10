package inplace

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	t.Parallel()

	type testData struct{ input, expectedResult []int }
	testsData := []testData{
		{
			input:          []int{1, 1, 2, 3, 4, 5},
			expectedResult: []int{5, 4, 3, 2, 1, 1},
		},
		{
			input:          []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedResult: []int{4, 3, 3, 2, 2, 1, 1, 1, 0, 0},
		},
		{
			input:          []int{0, 1, 1, 0},
			expectedResult: []int{0, 1, 1, 0},
		},
		{
			input:          []int{7, 2, 99, 0},
			expectedResult: []int{0, 99, 2, 7},
		},
	}

	for _, test := range testsData {
		reverseArray(test.input)
		if !reflect.DeepEqual(test.input, test.expectedResult) {
			t.Fatalf("expected array: %v\n got: %v", test.expectedResult, test.input)
		}
	}

}
