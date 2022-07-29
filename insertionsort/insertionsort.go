package insertionsort

// mine and not really correct one but result is same
func mine(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		tmp := slice[i]
		if tmp > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], tmp
			// this part not correct
			//----------------------
			for j := range slice[:i] {
				if slice[j] > slice[i] {
					tmp = slice[i]
					slice[i] = slice[j]
					slice[j] = tmp
				}
			}
			//----------------------
		}
	}
}

// TODO: разобрать
func notMine(slice []int) {
	for i := 1; i < len(slice); i++ {

		j := i
		for j > 0 && slice[j-1] > slice[j] {
			tmp := slice[j]
			slice[j] = slice[j-1]
			slice[j-1] = tmp
			j--
		}
	}
}
