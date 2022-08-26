package traversal

import "github.com/seriozhakorneev/go-data-structures/trees/binarytree"

func PreOrder[T any](root *binarytree.Node[T]) (stack []T) {
	var traverse func(node *binarytree.Node[T])
	traverse = func(node *binarytree.Node[T]) {
		if node == nil {
			return
		}

		// VISIT // LEFT // RIGHT
		stack = append(stack, node.Data)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return stack
}

func InOrder[T any](root *binarytree.Node[T]) (stack []T) {
	var traverse func(node *binarytree.Node[T])
	traverse = func(node *binarytree.Node[T]) {
		if node == nil {
			return
		}

		// LEFT // VISIT // RIGHT
		traverse(node.Left)
		stack = append(stack, node.Data)
		traverse(node.Right)
	}

	traverse(root)
	return stack
}

func PostOrder[T any](root *binarytree.Node[T]) (stack []T) {
	var traverse func(node *binarytree.Node[T])
	traverse = func(node *binarytree.Node[T]) {
		if node == nil {
			return
		}

		// LEFT // RIGHT // VISIT
		traverse(node.Left)
		traverse(node.Right)
		stack = append(stack, node.Data)
	}

	traverse(root)
	return stack
}
