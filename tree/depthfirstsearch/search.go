package depthfirstsearch

import (
	"github.com/seriozhakorneev/go-data-structures/stack"
	"github.com/seriozhakorneev/go-data-structures/trees/simpletree"
)

func search[T comparable](node *simpletree.Node[T]) (visited []T) {
	set := make(map[T]bool)
	next := stack.New[*simpletree.Node[T]](0)

	set[node.Data] = true
	next.Push(node)

	for !next.IsEmpty() {
		s, _ := next.Pop()

		visited = append(visited, s.Data)

		for i := len(s.Childrens) - 1; i >= 0; i-- {
			if !set[s.Childrens[i].Data] {
				set[s.Childrens[i].Data] = true
				next.Push(s.Childrens[i])
			}
		}
	}

	return visited
}

// TODO: add more simple implementation, recursive probably
