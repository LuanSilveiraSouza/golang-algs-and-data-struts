package binarySearchTree

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{nil}
}

func (tree *BinarySearchTree) Insert(value int) {
	newNode := Node{value, nil, nil}
	if tree.Root == nil {
		tree.Root = &newNode
	} else {
		currentNode := tree.Root
		for {
			if value < currentNode.Value {
				if currentNode.Left == nil {
					currentNode.Left = &newNode
					break
				} else {
					currentNode = currentNode.Left
					continue
				}
			} else if value > currentNode.Value {
				if currentNode.Right == nil {
					currentNode.Right = &newNode
					break
				} else {
					currentNode = currentNode.Right
					continue
				}
			}
			break
		}
	}
}

func (tree *BinarySearchTree) Search(value int) (*Node, error) {
	currentNode := tree.Root

	for {
		if currentNode == nil {
			return nil, errors.New("value dont exists in the tree")
		} else {
			if currentNode.Value == value {
				return currentNode, nil
			} else if value < currentNode.Value {
				currentNode = currentNode.Left
			} else {
				currentNode = currentNode.Right
			}
		}
	}
}

func (tree *BinarySearchTree) InOrderTraversal(root *Node) {
	if root != nil {
		if root.Left != nil {
			tree.InOrderTraversal(root.Left)
		}
		PrintNode(root)
		if root.Right != nil {
			tree.InOrderTraversal(root.Right)
		}
	}
}

func PrintNode(node *Node) {
	fmt.Printf("node %v \tvalue: %v\tleft: %p \tright: %p\n", &node, node.Value, node.Left, node.Right)
}
