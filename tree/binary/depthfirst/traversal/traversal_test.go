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

func TestPreOrder(t *testing.T) {
	t.Parallel()

	expected := []string{"F", "B", "A", "D", "C", "E", "G", "I", "H"}
	stack := PreOrder[string](testBinaryTree())

	if !reflect.DeepEqual(expected, stack) {
		t.Fatalf("Expected stack: %v\n Got: %v", expected, stack)
	}

}
func TestInOrder(t *testing.T) {
	t.Parallel()

	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	stack := InOrder[string](testBinaryTree())

	if !reflect.DeepEqual(expected, stack) {
		t.Fatalf("Expected stack: %v\n Got: %v", expected, stack)
	}
}
func TestPostOrder(t *testing.T) {
	t.Parallel()

	expected := []string{"A", "C", "E", "D", "B", "H", "I", "G", "F"}
	stack := PostOrder[string](testBinaryTree())

	if !reflect.DeepEqual(expected, stack) {
		t.Fatalf("Expected stack: %v\n Got: %v", expected, stack)
	}
}
