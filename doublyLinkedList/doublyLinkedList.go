package doublyLinkedList

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head *Node
	Size int
}

func NewDoublyLinkedList() DoublyLinkedList {
	return DoublyLinkedList{head: nil, Size: 0}
}

func (list *DoublyLinkedList) Insert(value interface{}) bool {
	node := Node{value: value, prev: nil, next: nil}	

	if list.head == nil {
		list.head = &node
		list.Size++
	} else {
		lastNode, err := list.Search(list.Size - 1)

		if err == nil {
			lastNode.next = &node
			node.prev = lastNode
			list.Size++
			return true
		}
	}
	return false
}

func (list *DoublyLinkedList) Search(position int) (resultNode *Node, err error) {
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

/*func (list *LinkedList) Delete(position int) (resultNode *Node, err error) {
	targetNode, err := list.Search(position)

	if err == nil {
		node := list.head

		for {
			if node.link == targetNode {
				node.link = targetNode.link
				list.Size--
				return targetNode, nil
			}
			node = node.link
		}
	}
	return list.head, err
}

func (list *LinkedList) Reverse() {
	node := list.head

	var previousNode *Node
	previousNode = nil

	var link *Node

	for {
		link = node.link
		node.link = previousNode
		previousNode = node
		node = link

		if node.link == nil {
			node.link = previousNode
			list.head = node
			break
		}
	}
}*/

func (list *DoublyLinkedList) Display() {
	fmt.Printf("head: %p\n", list.head)

	node := list.head

	for {
		fmt.Printf("value: %v\tprev: %p\tnext: %p\n", node.value, node.prev, node.next)

		if node.next == nil {
			break
		}
		node = node.next
	}
}
