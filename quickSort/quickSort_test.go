package quickSort

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/LuanSilveiraSouza/golang-algs-and-data-struts/utils"
)

func TestQuickSort(t *testing.T) {
	tests, err := utils.GetArray()

	if err != nil {
		fmt.Println(err)
	}

	for _, testCase := range tests {
		testSetup(
			t,
			testCase.Unsorted,
			testCase.Sorted,
		)
	}
}

func testSetup(t *testing.T, array []int, sortedArray []int) {
	t0 := time.Now()

	result := Sort(array, 0, len(array)-1)

	t1 := time.Now()

	if !reflect.DeepEqual(result, sortedArray) {
		t.Fatal("Result not expected")
	}

	fmt.Printf("Array length: %d\n Executed in: %s\n", len(array), t1.Sub(t0).String())
}
