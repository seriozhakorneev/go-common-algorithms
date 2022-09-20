package binarysearch

import "math"

func search(a []int, target int) int {
	left := 0
	right := len(a)

	for left+1 < right {
		mid := int(math.Ceil(float64(right+left) / 2))

		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if a[left] == target {
		return left
	}

	return -1
}
