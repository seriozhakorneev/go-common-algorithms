package insertionsort

import (
	"github.com/seriozhakorneev/go-common-algorithms/tests"
	"testing"
)

func TestInsertion(t *testing.T) {
	tests.Sort(t, sort)
}
