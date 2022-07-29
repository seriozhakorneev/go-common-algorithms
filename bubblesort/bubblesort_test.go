package bubblesort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMine(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		// Generate a random array of length i
		a := rand.Perm(i)

		mine(a)

		if !sort.IntsAreSorted(a) {
			t.Fatalf("%v array not sorted as expected", a)
		}
	}
}

func TestBetter(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		// Generate a random array of length i
		a := rand.Perm(i)

		better(a)

		if !sort.IntsAreSorted(a) {
			t.Fatalf("%v array not sorted as expected", a)
		}
	}
}
