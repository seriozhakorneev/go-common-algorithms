package windowsliding

// Window Sliding Technique is a computational technique
// which aims to reduce the use of nested loop and replace it with a single loop,
// thereby reducing the time complexity.
// window can have fixed range or dynamic like example:
func slidingWindowFindMaxSubstring(s string) (a string) {
	hash := make(map[uint8]int)
	for l, r := 0, 0; r < len(s)-1; r++ {

		// looking for already visited symbol in hash
		// and switching left index to it
		if repeat, ok := hash[s[r]]; ok && repeat > l {
			l = repeat
		}

		// saving every symbol
		hash[s[r]] = r + 1

		// computating difference between right and left indexes
		// if its max, saving it
		if hash[s[r]]-l > len(a) {
			a = s[l : r+1]
		}
	}
	return
}
