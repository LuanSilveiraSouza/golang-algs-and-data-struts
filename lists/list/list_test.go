package lists

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	list := List{}

	list.Push(2)
	list.Push(4)
	list.Push(6)

	if !reflect.DeepEqual(list.array, []interface{}{2, 4, 6}) {
		t.Fatal("Push method not worked as expected")
	}

	if !reflect.DeepEqual(list.Get(1), 4) {
		t.Fatal("Get method not worked as expected")
	}

	list.Pop()
	list.Pop()

	if !reflect.DeepEqual(list.array, []interface{}{2}) {
		t.Fatal("Pop method not worked as expected")
	}
}
