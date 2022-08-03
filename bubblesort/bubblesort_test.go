package bubblesort

import (
	"go-common-algorithms/tests"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests.Sort(t, v1)
	tests.Sort(t, v2)
}
