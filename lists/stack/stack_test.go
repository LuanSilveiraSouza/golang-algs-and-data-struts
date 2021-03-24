package lists

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	stack := Stack{}

	stack.Push(2)
	stack.Push(4)
	stack.Push(6)

	if !reflect.DeepEqual(stack.array, []interface{}{2, 4, 6}) {
		t.Fatal("Push method not worked as expected")
	}

	if !reflect.DeepEqual(stack.Top(), 6) {
		t.Fatal("Top method not worked as expected")
	}

	stack.Pop()

	if !reflect.DeepEqual(stack.array, []interface{}{2}) && !reflect.DeepEqual(stack.Top(), 4) {
		t.Fatal("Pop method not worked as expected")
	}
}
