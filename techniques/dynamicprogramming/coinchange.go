package dynamicprogramming

// problem https://leetcode.com/problems/coin-change/submissions/
// coinChange bottom up dynamic programming
func coinChange(coins []int, amount int) int {
	dpArray := make([]int, amount+1)
	dpArray[0] = 0
	for i := 1; i < len(dpArray); i++ {
		dpArray[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, n := range coins {
			if i-n >= 0 {
				dpArray[i] = min(dpArray[i], dpArray[i-n]+1)
			}
		}
	}

	if dpArray[amount] != amount+1 {
		return dpArray[amount]
	} else {
		return -1
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
