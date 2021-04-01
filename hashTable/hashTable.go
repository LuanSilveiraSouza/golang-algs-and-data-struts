package hashTable

import (
	"github.com/LuanSilveiraSouza/golang-algs-and-data-struts/linkedList"
)

type HashTable struct {
	Table []*linkedList.LinkedList
	Size  int
}

func NewHashTable(capacity int) *HashTable {
	table := make([]*linkedList.LinkedList, capacity)

	for i := 0; i < capacity; i++ {
		list := linkedList.NewLinkedList()
		table[i] = &list
	}

	return &HashTable{table, 0}
}

func (table *HashTable) Insert(key string) {
	index := table.hash(key)

	table.Size++

	table.Table[index].Insert(key)
}

func (table *HashTable) Search(key string) *linkedList.Node {
	index := table.hash(key)

	return table.Table[index].Search(key)
}

func (table *HashTable) Display() {
	for _, v := range table.Table {
		v.Display()
	}
}

func (table *HashTable) hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % len(table.Table)
}
