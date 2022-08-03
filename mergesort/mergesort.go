package mergesort

func sort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	f := sort(a[:len(a)/2])
	s := sort(a[len(a)/2:])

	return concatArrays(f, s)
}

func concatArrays(a, b []int) []int {
	var c []int

	for len(a) != 0 && len(b) != 0 {
		if a[0] > b[0] {
			c = append(c, b[0])
			b = b[1:]
		} else {
			c = append(c, a[0])
			a = a[1:]
		}
	}

	for len(a) != 0 {
		c = append(c, a[0])
		a = a[1:]
	}
	for len(b) != 0 {
		c = append(c, b[0])
		b = b[1:]
	}
	return c
}
