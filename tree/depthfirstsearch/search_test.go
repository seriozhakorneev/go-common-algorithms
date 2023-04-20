package depthfirstsearch

import (
	"reflect"
	"testing"

	"github.com/seriozhakorneev/go-data-structures/trees/simpletree"
)

func createTree() *simpletree.Tree[any] {
	tree := simpletree.NewTree[any]("A")

	b := tree.Root.AddNode("B")
	b.AddNode("C")
	b.AddNode("D")
	b.AddNode("E").AddNode("F")

	tree.Root.AddNode("G").AddNode("H").AddNode("I")

	tree.FindDepth()

	return &tree
}

func TestSearch(t *testing.T) {
	exp := []any{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	tree := createTree()

	result := search(tree.Root)

	if !reflect.DeepEqual(exp, result) {
		t.Fatalf("Expected sequence: %v\nGot: %v", exp, result)
	}
}
