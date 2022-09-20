package binary

import (
	"github.com/seriozhakorneev/go-data-structures/trees/binarytree"
)

func search(root *binarytree.Node[int], target int) *binarytree.Node[int] {

	var rec func(node *binarytree.Node[int]) *binarytree.Node[int]
	rec = func(node *binarytree.Node[int]) *binarytree.Node[int] {
		if node == nil {
			return nil
		} else if target == node.Data {
			return node
		} else if target > node.Data {
			return rec(node.Right)
		} else {
			return rec(node.Left)
		}
	}
	return rec(root)
}
