package traversal

import (
	"github.com/seriozhakorneev/go-data-structures/queue"
	"github.com/seriozhakorneev/go-data-structures/trees/binarytree"
)

func LevelOrder[T any](root *binarytree.Node[T]) (visited []T) {
	q := queue.New[*binarytree.Node[T]](0)
	q.Enqueue(root)

	for !q.IsEmpty() {
		node, _ := q.Dequeue()
		visited = append(visited, node.Data)

		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}

	return visited
}
