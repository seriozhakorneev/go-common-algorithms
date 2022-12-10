package prefixsum

import (
	"reflect"
	"testing"
)

func TestCountOfPositive(t *testing.T) {
	t.Parallel()

	nums := []int{2, -1, 2, -2, 3}
	indexes := [][]int{{1, 1}, {1, 3}, {2, 4}, {1, 5}}
	expectedCounts := []int{1, 2, 1, 3}

	counts := countOfPositive(nums, indexes)

	if !reflect.DeepEqual(counts, expectedCounts) {
		t.Fatalf("expected counts: %v, got: %v", expectedCounts, counts)
	}

}
