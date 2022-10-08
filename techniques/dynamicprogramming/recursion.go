package dynamicprogramming

// findFibNumberRecursion returns fibonacci value under the number of n
func findFibNumberRecursion(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return findFibNumberRecursion(n-1) +
			findFibNumberRecursion(n-2)
	}
}
