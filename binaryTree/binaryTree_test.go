package binaryTree

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := NewBinaryTree()

	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	tree.Display()

	if !reflect.DeepEqual(tree.Size, 4) {
		t.Fatal("Size attribute not worked as expected")
	}
}