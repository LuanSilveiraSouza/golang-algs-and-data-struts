package linearSearch

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/LuanSilveiraSouza/golang-algs-and-data-struts/utils"
)

func TestLinearSearch(t *testing.T) {
	tests, err := utils.GetArray()

	if err != nil {
		fmt.Println(err)
	}

	array := tests[0].Sorted

	if !reflect.DeepEqual(LinearSearch(array, 94), 2) {
		t.Fatal("linear Search failed")
	}

	if !reflect.DeepEqual(LinearSearch(array, 9991), 99) {
		t.Fatal("linear Search failed")
	}
}
