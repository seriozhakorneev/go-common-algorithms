package tests

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func Sort(t *testing.T, sortFunc func([]int)) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		// Generate a random array of length i
		a := rand.Perm(i)

		sortFunc(a)

		if !sort.IntsAreSorted(a) {
			t.Fatalf("%v array not sorted as expected", a)
		}
	}
}
