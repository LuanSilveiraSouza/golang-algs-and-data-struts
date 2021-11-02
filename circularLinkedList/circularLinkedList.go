package circularlinkedlist

import (
	"errors"
)

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

type CircularLinkedList struct {
	head *Node
	Size int
}

func NewCircularLinkedList() CircularLinkedList {
	return CircularLinkedList{head: nil, Size: 0}
}

func (list *CircularLinkedList) Insert(value interface{}) bool {
	node := Node{value: value, prev: nil, next: nil}

	if list.head == nil {
		list.head = &node
		node.next = list.head
		list.Size++
	} else {
		lastNode, err := list.Search(list.Size - 1)

		if err == nil {
			lastNode.next = &node
			node.prev = lastNode
			node.next = list.head
			list.Size++
			return true
		}
	}
	return false
}

func (list *CircularLinkedList) Search(position int) (resultNode *Node, err error) {
	node := list.head

	if position >= list.Size || position < 0 {
		return node, errors.New("Not found")
	}

	i := 0
	for {
		if i == position {
			return node, nil
		}
		i++
		node = node.next
	}
}
