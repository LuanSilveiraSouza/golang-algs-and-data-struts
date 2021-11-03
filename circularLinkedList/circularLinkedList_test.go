package circularlinkedlist

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	list := NewCircularLinkedList()

	list.Insert(2)
	list.Insert(4)
	list.Insert(6)

	if !reflect.DeepEqual(list.Size, 3) {
		t.Fatal("Size attribute not worked as expected")
	}

	result, err := list.Search(2)
	if !reflect.DeepEqual(result.value, 6) && reflect.DeepEqual(err, nil) {
		t.Fatal("Search method not worked as expected")
	}
	if !reflect.DeepEqual(result.next, list.head) {
		t.Fatal("Circular link not worked as expected")
	}

	list.Reverse()

	result, err = list.Search(0)
	if !reflect.DeepEqual(result.value, 6) && reflect.DeepEqual(err, nil) {
		t.Fatal("Reverse method not worked as expected")
	}

	list.Delete(2)

	if !reflect.DeepEqual(list.Size, 2) {
		t.Fatal("Delete attribute not worked as expected")
	}
}
