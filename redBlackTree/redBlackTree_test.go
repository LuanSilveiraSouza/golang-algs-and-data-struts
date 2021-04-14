package redBlackTree

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := NewRBTree()

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(4)

	tree.InOrderTraversal(tree.root)

	searched, error := tree.Search(1)

	if !reflect.DeepEqual(error, nil) {
		t.Fatal("Search method not worked as expected")
	}

	if !reflect.DeepEqual(searched.value, 3) {
		t.Fatal("Search method not worked as expected")
	}

	searched, error = tree.Search(7)

	if reflect.DeepEqual(error, "value dont exists in the tree") {
		t.Fatal("Search method not worked as expected")
	}

	if reflect.DeepEqual(searched, nil) {
		t.Fatal("Search method not worked as expected")
	}
}
