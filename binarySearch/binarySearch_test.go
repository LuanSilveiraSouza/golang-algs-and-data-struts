package binarySearch

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/LuanSilveiraSouza/golang-algs-and-data-struts/utils"
)

func TestBinarySearch(t *testing.T) {
	tests, err := utils.GetArray()

	if (err != nil) {
		fmt.Println(err)
	}

	array := tests[0].Sorted

	if !reflect.DeepEqual(BinarySearch(array, 94, 0, len(array)), 2) {
		t.Fatal("Binary Search failed")
	}

	if !reflect.DeepEqual(BinarySearch(array, 9991, 0, len(array)), 99) {
		t.Fatal("Binary Search failed")
	}
}
