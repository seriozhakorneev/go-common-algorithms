package binary

import (
	"github.com/seriozhakorneev/go-data-structures/trees/btree"
)

func search(root *btree.Node, target int) *btree.Node {
	var rec func(node *btree.Node) *btree.Node

	rec = func(node *btree.Node) *btree.Node {
		if node == nil {
			return nil
		} else if target == node.Value {
			return node
		} else if target > node.Value {
			return rec(node.Right)
		} else {
			return rec(node.Left)
		}
	}

	return rec(root)
}
