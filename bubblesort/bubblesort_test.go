package bubblesort

import (
	"go-common-algorithms/tests"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests.Sort(t, mine)
	tests.Sort(t, notMine)
}
