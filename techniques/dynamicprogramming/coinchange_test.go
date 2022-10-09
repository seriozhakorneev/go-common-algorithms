package dynamicprogramming

import (
	"testing"
)

func TestCoinChange(t *testing.T) {

	type testCoinChange struct {
		coins                  []int
		amount, expectedResult int
	}

	tests := []testCoinChange{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}

	for _, test := range tests {
		result := coinChange(test.coins, test.amount)
		if result != test.expectedResult {
			t.Fatalf("expected result: %v, got %v", test.expectedResult, result)
		}
	}

}
