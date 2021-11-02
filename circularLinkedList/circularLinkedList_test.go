package circularlinkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	list := NewCircularLinkedList()

	list.Insert(2)
	list.Insert(4)

	if !reflect.DeepEqual(list.Size, 2) {
		t.Fatal("Size attribute not worked as expected")
	}

	result, err := list.Search(1)
	if !reflect.DeepEqual(result.value, 4) && reflect.DeepEqual(err, nil) {
		t.Fatal("Search method not worked as expected")
	}
	fmt.Println(result.next)
	if !reflect.DeepEqual(result.next, list.head) {
		t.Fatal("Circular link not worked as expected")
	}
}
