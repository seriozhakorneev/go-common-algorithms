package dynamicprogramming

// any of them returns fibonacci value under the number of n in fibonacci 1,1....

// recursion solution
func recursion(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return recursion(n-1) + recursion(n-2)
	}
}

// memoized solution with store for already counted values
func memoized(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n == 1 || n == 2 {
		return 1
	} else {
		memo[n] = memoized(n-1, memo) + memoized(n-2, memo)
		return memo[n]
	}
}

// bottomUp solution with dynamically generated store
func bottomUp(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	store := make([]int, n)
	store[0], store[1] = 1, 1
	for i := 2; i < n; i++ {
		store[i] = store[i-1] + store[i-2]
	}
	return store[n-1]
}
