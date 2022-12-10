package countingsort

func SortRunes(s []rune, max int) {
	// generate slice of zeros with length = max element of ascii slice
	max++
	cnt := make([]int, max, max)

	for _, r := range s {
		// we mark value with 1 at ascii code index [int(r)] in cnt
		cnt[int(r)]++
	}

	for i, j := 0, 0; i < max; i++ {
		// setting cnt slice index with 1 value in ascii slice with rune convertion
		// decrement mark by 1
		// set index of i to beginning, index of j to a next value
		if cnt[i] != 0 {
			s[j] = rune(i)
			cnt[i]--
			i = 0
			j++
		}
	}
}

func maxRune(runes []rune) int {
	max := 0
	for _, r := range runes {
		code := int(r)
		if code > max {
			max = code
		}
	}

	return max
}

//TODO make ints version
