package linkedList

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	list := NewLinkedList()

	list.Insert(2)
	list.Insert(4)
	list.Insert(6)

	if !reflect.DeepEqual(list.Size, 3) {
		t.Fatal("Size attribute not worked as expected")
	}

	result, err := list.Get(list.Size - 2)
	if !reflect.DeepEqual(result.value, 4) && reflect.DeepEqual(err, nil) {
		t.Fatal("Get method not worked as expected")
	}

	searched := list.Search(7)
	if reflect.DeepEqual(searched, nil) {
		t.Fatal("Search method not worked as expected")
	}

	list.Reverse()

	result, err = list.Get(list.Size - 1)
	if !reflect.DeepEqual(result.value, 2) && reflect.DeepEqual(err, nil) {
		t.Fatal("Reverse method not worked as expected")
	}

	result, err = list.Delete(1)
	if !reflect.DeepEqual(result.value, 4) && reflect.DeepEqual(err, nil) {
		t.Fatal("Delete method not worked as expected")
	}

	if !reflect.DeepEqual(list.Size, 2) {
		t.Fatal("Size attribute not worked as expected")
	}

	result, err = list.Get(0)
	if !reflect.DeepEqual(result.value, 6) && reflect.DeepEqual(err, nil) {
		t.Fatal("Search method not worked as expected")
	}

}
