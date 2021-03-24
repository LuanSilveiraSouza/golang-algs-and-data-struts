package binarySearchTree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
	Size int
}

func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{nil, 0}
}

func (tree *BinarySearchTree) Insert(value int) {
	newNode := Node{value, nil, nil}
	if tree.Root == nil {
		tree.Root = &newNode
		tree.Size++
	} else {
		currentNode := tree.Root
		for {
			if value < currentNode.Value {
				if currentNode.Left == nil {
					currentNode.Left = &newNode
					tree.Size++
					break
				} else {
					currentNode = currentNode.Left
					continue
				}
			} else if value > currentNode.Value {
				if currentNode.Right == nil {
					currentNode.Right = &newNode
					tree.Size++
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

func (tree *BinarySearchTree) Display() {
	loopNode(tree.Root)
}

func loopNode(node *Node) {
	if node != nil {
		fmt.Printf("node %v \tvalue: %v\tleft: %p \tright: %p\n", &node, node.Value, node.Left, node.Right)

		if node.Left != nil {
			loopNode(node.Left)
		}
		if node.Right != nil {
			loopNode(node.Right)
		}
	}
}
