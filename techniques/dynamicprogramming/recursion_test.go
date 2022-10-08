package dynamicprogramming

import (
	"testing"
)

func TestFindFibNumberRecursion(t *testing.T) {
	seq := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i := 0; i < len(seq); i++ {
		expected, value := seq[i], findFibNumberRecursion(i+1)
		if expected != value {
			t.Fatalf("expected value %v, got %v", expected, value)
		}
	}
}
