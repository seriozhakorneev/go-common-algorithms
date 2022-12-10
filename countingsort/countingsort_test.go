package countingsort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/seriozhakorneev/go-common-algorithms/insertionsort"
)

func TestCountingSortRunes(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		runes, ints := make([]rune, 0), make([]int, 0)

		for j := 0; j < 100; j++ {
			rand.Seed(time.Now().UnixNano())

			random := rand.Intn(26)
			runes = append(runes, rune('a'+random))
			ints = append(ints, 'a'+random)
		}

		SortRunes(runes, maxRune(runes))
		insertionsort.Sort(ints)

		for k := 0; k < len(ints); k++ {
			if rune(ints[k]) != runes[k] {
				t.Fatalf("elements are not equal at index(%d)\n"+
					"elements: ints(%d) runes(%d)\n "+
					"in compare of integers\n(%v)\nrunes\n()%v",
					k, ints[k], runes[k], ints, runes)
			}
		}
	}
}
