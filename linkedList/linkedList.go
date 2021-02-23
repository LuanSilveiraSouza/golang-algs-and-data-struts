package linkedList

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	link *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	Size int
}

func NewLinkedList() LinkedList {
	return LinkedList{head: nil, tail: nil, Size: 0} 
}

func (list *LinkedList) Insert(value interface{}) {
	node := Node{value: value, link: nil}

	if list.head == nil {
		list.head = &node
		list.tail = &node
	} else {
		list.tail.link = &node
		list.tail = &node
	}

	list.Size++
}

func (list *LinkedList) Search(position int) (resultNode Node, err error) {
	node := list.head

	for i := 0; i < list.Size; i++ {
		if i == position {
			return *node, nil
		}
		node = node.link
	}
	return *list.tail, errors.New("Not found")
}

func (list *LinkedList) Delete(position int) (resultNode Node, err error) {
	targetNode, err := list.Search(position)

	if err == nil {
		node := list.head
		fmt.Println(*node.link, targetNode)
		for {
			if *node.link == targetNode {
				node.link = targetNode.link
				list.Size--
				return targetNode, nil
			}
			node = node.link
		}
	}
	return *list.tail, err
}

func (list *LinkedList) Display() {
	fmt.Printf("head: %p\n", list.head)

	node := list.head
	
	for {
		fmt.Printf("value: %v\tnext: %p\n", node.value, node.link)

		if node.link == nil {
			break
		}	
		node = node.link
	}
}