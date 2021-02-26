package selectionSort

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/LuanSilveiraSouza/golang-algs-and-data-struts/utils"
)

func TestInsertionSort(t *testing.T) {
	tests, err := utils.GetArray()

	if (err != nil) {
		fmt.Println(err)
	}

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
		fmt.Println(sortedArray)
		fmt.Println(result)
		t.Fatal("Result not expected on array")
	}

	fmt.Printf("Array length: %d\n Executed in: %s\n", len(array), t1.Sub(t0).String())
}
