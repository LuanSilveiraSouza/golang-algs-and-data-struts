package search

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/LuanSilveiraSouza/golang-algs-and-data-struts/utils"
)

func TestInterpolationSearch(t *testing.T) {
	tests, err := utils.GetArray()

	if err != nil {
		fmt.Println(err)
	}

	array := tests[0].Sorted

	if !reflect.DeepEqual(InterpolationSearch(array, 94, 0, len(array)-1), 2) {
		t.Fatal("interpolation Search failed")
	}

	if !reflect.DeepEqual(InterpolationSearch(array, 9991, 0, len(array)-1), 99) {
		t.Fatal("interpolation Search failed")
	}
}
