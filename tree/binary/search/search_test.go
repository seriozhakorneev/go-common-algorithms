package binary

import (
	"github.com/seriozhakorneev/go-data-structures/trees/binarytree"
	"testing"
)

func TestSearch(t *testing.T) {

	for i := 11; i < 1102; i++ {
		tree := binarytree.GenerateFromRange(1, i)
		expectedData := tree.Root.Left.Right.Data

		node := search(tree.Root, expectedData)
		if expectedData != node.Data {
			t.Fatalf("Expected data: %v, Got %v", expectedData, node.Data)
		}
	}
}
