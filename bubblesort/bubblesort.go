package bubblesort

// notMine version
func notMine(a []int) {
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
			}
		}
	}
}

// mine version
func mine(a []int) {
	if len(a) < 1 {
		return
	}

	for {
		var swapped bool
		for i, el := range a[:len(a)-1] {
			if el > a[i+1] || i == len(a)-1 {
				tmp := a[i]
				a[i] = a[i+1]
				a[i+1] = tmp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
