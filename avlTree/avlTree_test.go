package avlTree

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := NewAvlTree()

	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(45)
	tree.Insert(50)
	tree.Insert(48)
	tree.Insert(55)

	tree.InOrderTraversal(tree.Root)

	searched, error := tree.Search(55)

	if !reflect.DeepEqual(error, nil) {
		t.Fatal("Search method not worked as expected")
	}

	if !reflect.DeepEqual(searched.Value, 55) {
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
