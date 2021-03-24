package lists

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	queue := Queue{}

	queue.Push(2)
	queue.Push(4)

	if !reflect.DeepEqual(queue.Size, 2) {
		t.Fatal("Size not worked as expected")
	}

	queue.Push(6)

	if !reflect.DeepEqual(queue.array, []interface{}{6, 4, 2}) {
		t.Fatal("Push method not worked as expected")
	}

	if !reflect.DeepEqual(queue.Front(), 2) {
		t.Fatal("Front method not worked as expected")
	}

	queue.Pop()

	if !reflect.DeepEqual(queue.array, []interface{}{6}) && !reflect.DeepEqual(queue.Front(), 4) {
		t.Fatal("Pop method not worked as expected")
	}
}
