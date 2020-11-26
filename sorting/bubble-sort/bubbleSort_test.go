package sorting

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := [5]int{5, 4, 3, 2, 1}

	result := BubbleSort(array)

	if result != [5]int{1, 2, 3, 4, 5} {
		t.Fatal("Result not expected")
	}
}
