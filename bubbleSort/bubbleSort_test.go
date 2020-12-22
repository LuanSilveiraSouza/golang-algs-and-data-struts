package bubbleSort

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"github.com/LuanSilveiraSouza/golang-algs-and-data-struts/utils"
)

func TestBubbleSort(t *testing.T) {
	tests := utils.GetArray()

	for _, testCase	:= range tests {
		testSetup(
			t,
			testCase.Unsorted,
			testCase.Sorted,
		)
	} 
}

func testSetup(t *testing.T, array []int, sortedArray []int) {
	t0 := time.Now()

	result := Sort(array)

	t1 := time.Now()

	if !reflect.DeepEqual(result, sortedArray) {
		t.Fatal("Result not expected")
	}

	fmt.Printf("Array length: %d\n Executed in: %dms\n", len(array), t1.Sub(t0))
}