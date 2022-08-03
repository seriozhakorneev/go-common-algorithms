package mergesort

import (
	"go-common-algorithms/tests"
	"testing"
)

func TestMerge(t *testing.T) {
	tests.Sort(t, sort)
}
