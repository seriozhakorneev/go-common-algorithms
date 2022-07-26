package main

import (
	"fmt"
)

// better version
func better(a []int) {
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				swap := a[j]
				a[j] = a[j+1]
				a[j+1] = swap
			}
		}
	}
}

// mine version
func mine(a []int) {
	if len(a) < 1 {
		return
	}

	for {
		var swapped bool
		for i, el := range a[:len(a)-1] {
			if el > a[i+1] || i == len(a)-1 {
				swap := a[i]
				a[i] = a[i+1]
				a[i+1] = swap
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	a := []int{9, 10, 11, 224, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	better(a)
	//mine(a)
	fmt.Println(a)
}
