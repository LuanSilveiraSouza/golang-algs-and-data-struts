package bubbleSort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int {5, 4, 3, 2, 1}

	result := Sort(array)

	if !reflect.DeepEqual(result, []int{1, 2, 3, 4, 5}) {
		t.Fatal("Result not expected")
	}
}
