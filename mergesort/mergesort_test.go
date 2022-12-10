package mergesort

import (
	"github.com/seriozhakorneev/go-common-algorithms/tests"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	tests.Sort(t, sort)
}
