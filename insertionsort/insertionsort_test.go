package insertionsort

import (
	"github.com/seriozhakorneev/go-common-algorithms/tests"
	"testing"
)

func TestInsertion(t *testing.T) {
	t.Parallel()

	tests.Sort(t, Sort)
}
