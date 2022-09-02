package linkedlistmerge

import (
	sll "github.com/seriozhakorneev/go-data-structures/linkedlist/singlylinkedlist"
)

func mergeTwoLists(
	list1, list2 *sll.Node[int]) *sll.Node[int] {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
