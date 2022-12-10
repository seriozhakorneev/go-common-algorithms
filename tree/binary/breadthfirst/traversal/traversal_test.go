package traversal

import (
	"github.com/seriozhakorneev/go-data-structures/trees/binarytree"
	"reflect"
	"testing"
)

func testBinaryTree() *binarytree.Node[string] {
	t := binarytree.NewTree("F")

	t.Root.Left = &binarytree.Node[string]{Data: "B"}
	t.Root.Left.Left = &binarytree.Node[string]{Data: "A"}
	t.Root.Left.Right = &binarytree.Node[string]{Data: "D"}
	t.Root.Left.Right.Left = &binarytree.Node[string]{Data: "C"}
	t.Root.Left.Right.Right = &binarytree.Node[string]{Data: "E"}

	t.Root.Right = &binarytree.Node[string]{Data: "G"}
	t.Root.Right.Right = &binarytree.Node[string]{Data: "I"}
	t.Root.Right.Right.Left = &binarytree.Node[string]{Data: "H"}

	return t.Root
}

func TestLevelOrder(t *testing.T) {
	t.Parallel()

	expected := []string{"F", "B", "G", "A", "D", "I", "C", "E", "H"}
	visited := LevelOrder(testBinaryTree())

	if !reflect.DeepEqual(expected, visited) {
		t.Fatalf("Expected visited: %v\n Got: %v", expected, visited)
	}
}
