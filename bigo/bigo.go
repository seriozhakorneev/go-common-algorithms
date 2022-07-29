package main

import "fmt"

// O(0) - сложность алгоритма не выполняющего ничего
// O(1) - сложность алгоритма выполняющего константное кол-во операций(1,2,3)

func sum(n int) int {
	if n == 1 {
		return 1
	}
	// чем больше будет значение n
	// тем больше раз вызовется sum(n-1)
	// O(N)
	return n + sum(n-1)
}

func pairSumSequence(n int) int {
	var sum int
	// чем больше будет значение n
	// тем больше раз выполнится цикл i < n
	// O(N)
	for i := 0; i < n; i++ {
		sum += pairSum(i, i+1)
	}
	return sum
}

func pairSum(a, b int) int {
	// не имеет ни циклов ни рекурсий
	// выполняет константное количество операций
	// O(1)
	return a + b
}

func main() {
	fmt.Println(sum(3))
	fmt.Println(pairSumSequence(3))
}
