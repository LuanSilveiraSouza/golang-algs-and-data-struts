package linkedList

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	link  *Node
}

type LinkedList struct {
	head *Node
	Size int
}

func NewLinkedList() LinkedList {
	return LinkedList{head: nil, Size: 0}
}

func (list *LinkedList) Insert(value interface{}) bool {
	node := Node{value: value, link: nil}

	if list.head == nil {
		list.head = &node
		list.Size++
	} else {
		lastNode, err := list.Get(list.Size - 1)

		if err == nil {
			lastNode.link = &node
			list.Size++
			return true
		}
	}
	return false
}

func (list *LinkedList) Get(position int) (resultNode *Node, err error) {
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
		node = node.link
	}
}

func (list *LinkedList) Search(key interface{}) (resultNode *Node) {
	node := list.head

	for {
		if node == nil {
			return nil
		} else if node.value == key {
			return node
		}
		node = node.link
	}
}

func (list *LinkedList) Delete(position int) (resultNode *Node, err error) {
	targetNode, err := list.Get(position)

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
}

func (list *LinkedList) Display() {
	fmt.Printf("head: %p\n", list.head)

	node := list.head

	for {
		fmt.Printf("value: %v\tlink: %p\n", node.value, node.link)

		if node.link == nil {
			break
		}
		node = node.link
	}
}
