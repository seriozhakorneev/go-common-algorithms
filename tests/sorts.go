package tests

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func Sort(t *testing.T, sortFunc interface{}) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		// Generate a random array of length i
		a := rand.Perm(i)

		if f, ok := sortFunc.(func([]int) []int); ok {
			a = f(a)
		} else if f, ok := sortFunc.(func([]int)); ok {
			f(a)
		}

		if !sort.IntsAreSorted(a) {
			t.Fatalf("%v array not sorted as expected", a)
		}
	}
}
