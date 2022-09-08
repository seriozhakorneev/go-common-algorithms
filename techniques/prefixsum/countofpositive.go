package prefixsum

// count positive digits in range of two indexes
// nums := []int{2, -1, 2, -2, 3}
// indexes := [][]int{{1, 1},{1, 3},{2, 4},{1, 5}}
// return  []int{1, 2, 1, 3}

func countOfPositive(nums []int, indexes [][]int) []int {

	// making a slice with count of positive digits on i position
	//    [2, -1, 2, -2, 3] -- nums
	// [0, 1,  1, 2,  2, 3] -- how many positive elements on i position(inclusive)
	// posCount[0] = 0 element required for [:1] situation
	posCount := make([]int, len(nums)+1)
	posCount[0] = 0
	for i, num := range nums {
		posCount[i+1] = posCount[i]
		if num >= 1 {
			posCount[i+1]++
		}
	}

	result := make([]int, len(indexes))
	for i, index := range indexes {
		result[i] = posCount[index[1]] - posCount[index[0]-1]
	}

	return result
}
