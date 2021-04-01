package hashTable

import (
	"reflect"
	"testing"
)

func TestHashTable(t *testing.T) {
	table := NewHashTable(4)

	table.Insert("JOHN")
	table.Insert("MARY")
	table.Insert("CARLTON")
	table.Insert("TROY")

	table.Display()

	if !reflect.DeepEqual(table.Size, 3) {
		t.Fatal("Size attribute not worked as expected")
	}

	result := table.Search("JOHN")
	if reflect.DeepEqual(reflect.TypeOf(result), "*linkedList.Node") {
		t.Fatal("Search method not worked as expected")
	}

	result = table.Search("TEST")
	if reflect.DeepEqual(result, nil) {
		t.Fatal("Search method not worked as expected")
	}
}
