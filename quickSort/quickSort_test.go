package quickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int {34, 7, 2, 562, 6, 3}

	result := Sort(array, 0, len(array) - 1)

	if !reflect.DeepEqual(result, []int{2, 3, 6, 7, 34, 562}) {
		t.Fatal("Result not expected")
	}
}
