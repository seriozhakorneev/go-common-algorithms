package inplace

func reverseArray(array []int) {
	length := len(array)
	for i, j := 0, length-1; i < length/2; i++ {
		array[i], array[j] = array[j], array[i]
		j--
	}

	// or
	//for i := 0; i < length/2; i++ {
	//	array[i], array[length-i-1] = array[length-i-1], array[i]
	//}
}
