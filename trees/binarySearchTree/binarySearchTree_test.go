package binarySearchTree

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := NewBinarySearchTree()

	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(8)

	tree.LevelOrderTraversal()

	searched, error := tree.Search(6)

	if !reflect.DeepEqual(error, nil) {
		t.Fatal("Search method not worked as expected")
	}

	if !reflect.DeepEqual(searched.Value, 6) {
		t.Fatal("Search method not worked as expected")
	}

	searched, error = tree.Search(5)

	if reflect.DeepEqual(error, "value dont exists in the tree") {
		t.Fatal("Search method not worked as expected")
	}

	if reflect.DeepEqual(searched, nil) {
		t.Fatal("Search method not worked as expected")
	}
}
