package insertionsort

import (
	"fmt"
	"go-common-algorithms/tests"
	"testing"
)

func TestInsertion(t *testing.T) {
	tests.Sort(t, mine)
	tests.Sort(t, notMine)
}

func TestName(t *testing.T) {
	a := []int{52, 15, 3, 6, 23, 9000, 47, 435, 1, 234}
	notMine(a)

	fmt.Println(a)
}
