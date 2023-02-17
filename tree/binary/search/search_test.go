package binary

import (
	"github.com/seriozhakorneev/go-data-structures/trees/btree"

	"testing"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	for i := 11; i < 1102; i++ {
		tree := btree.GenFromRange(1, i)
		expectedNode := tree.Root.Left.Right

		node := search(tree.Root, expectedNode.Value)

		if expectedNode != node {
			t.Fatalf("Expected node: %v, Got %v", expectedNode, node)
		}
	}
}
