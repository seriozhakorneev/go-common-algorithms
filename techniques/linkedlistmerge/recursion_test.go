package linkedlistmerge

import (
	sll "github.com/seriozhakorneev/go-data-structures/linkedlist/singlylinkedlist"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	t.Parallel()

	type testData struct{ list1, list2, expectedResult []int }
	testsData := []testData{
		{
			list1:          []int{1, 2, 4},
			list2:          []int{1, 3, 4},
			expectedResult: []int{1, 1, 2, 3, 4, 4},
		},
		{
			list1:          nil,
			list2:          nil,
			expectedResult: nil,
		},
		{
			list1:          nil,
			list2:          []int{0},
			expectedResult: []int{0},
		},
		{
			list1:          []int{1},
			list2:          nil,
			expectedResult: []int{1},
		},
		{
			list1:          []int{2},
			list2:          []int{1},
			expectedResult: []int{1, 2},
		},
	}

	var result []int
	for _, test := range testsData {
		result = llToArr(
			mergeTwoLists(
				fillWithIntegers(test.list1...),
				fillWithIntegers(test.list2...),
			))

		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Fatalf("expected result: %v, but got: %v", test.expectedResult, result)
		}
	}
}

func fillWithIntegers(a ...int) *sll.Node[int] {
	if len(a) == 0 {
		return nil
	}
	node := &sll.Node[int]{Value: a[0]}
	node.Next = fillWithIntegers(a[1:]...)
	return node
}

func llToArr(node *sll.Node[int]) (arr []int) {
	if node == nil {
		return nil
	}
	arr = append(arr, node.Value)
	return append(arr, llToArr(node.Next)...)
}
