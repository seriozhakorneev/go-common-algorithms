package mergesort

import (
	"github.com/seriozhakorneev/go-common-algorithms/tests"
	"testing"
)

func TestMerge(t *testing.T) {
	tests.Sort(t, sort)
}
