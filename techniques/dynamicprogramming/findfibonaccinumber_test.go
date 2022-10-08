package dynamicprogramming

import (
	"testing"
)

func TestRecursion(t *testing.T) {
	seq := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i := 0; i < len(seq); i++ {
		expected, value := seq[i], recursion(i+1)
		if expected != value {
			t.Fatalf("expected value %v, got %v", expected, value)
		}
	}
}

func TestMemoised(t *testing.T) {
	seq := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	memo := make(map[int]int)
	for i := 0; i < len(seq); i++ {
		expected, value := seq[i], memoized(i+1, memo)
		if expected != value {
			t.Fatalf("expected value %v, got %v", expected, value)
		}
	}

	expectedMemo := map[int]int{3: 2, 4: 3, 5: 5, 6: 8, 7: 13, 8: 21, 9: 34, 10: 55}
	for k, v := range memo {
		if expectedMemo[k] != v {
			t.Fatalf("expected memo[%v]: %v, got: %v", k, v, expectedMemo[k])
		}
	}
}

func TestBottomUp(t *testing.T) {
	seq := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i := 0; i < len(seq); i++ {
		expected, value := seq[i], bottomUp(i+1)
		if expected != value {
			t.Fatalf("expected value %v, got %v", expected, value)
		}
	}
}
