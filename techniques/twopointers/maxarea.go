package twopointers

// https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {

	a, b, max := 0, len(height)-1, 0

	for a < b {
		area := min(height[a], height[b]) * (b - a)
		if area > max {
			max = area
		}

		if height[a] < height[b] {
			a++
		} else {
			b--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
