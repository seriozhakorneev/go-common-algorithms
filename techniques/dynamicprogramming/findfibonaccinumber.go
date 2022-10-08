package dynamicprogramming

// findFibNumber returns fibonacci value under the number of n
func findFibNumber(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return findFibNumber(n-1) + findFibNumber(n-2)
	}
}
