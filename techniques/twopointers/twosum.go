package twopointers

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}
	return []int{}
}
