package bubblesort

import (
	"github.com/seriozhakorneev/go-common-algorithms/tests"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	tests.Sort(t, v1)
	tests.Sort(t, v2)
}
