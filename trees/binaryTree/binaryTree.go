package binaryTree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
	Size int
}

func NewBinaryTree() BinaryTree {
	return BinaryTree{nil, 0}
}

func (tree *BinaryTree) Insert(value int) bool {
	newNode := Node{value, nil, nil}

	if tree.Root == nil {
		tree.Root = &newNode
		tree.Size++
		return true
	} else {
		var nodes = []*Node{}
		nodes = append(nodes, tree.Root)

		for {
			newNodeCollection := []*Node{}
			for _, node := range nodes {
				if node.Left == nil {
					node.Left = &newNode
					tree.Size++
					return true
				} else {
					newNodeCollection = append(newNodeCollection, node.Left)
				}
				if node.Right == nil {
					node.Right = &newNode
					tree.Size++
					return true
				} else {
					newNodeCollection = append(newNodeCollection, node.Left)
				}
			}
			nodes = newNodeCollection
		}
	}
}

func (tree *BinaryTree) Display() {
	fmt.Printf("root: %p\n", tree.Root)

	loopNode(tree.Root, 0)
}

func loopNode(node *Node, level int) {
	fmt.Printf("value: %v\tleft: %p \tright: %p\n", node.Value, node.Left, node.Right)

	if node.Left != nil {
		loopNode(node.Left, level+1)
	}
	if node.Right != nil {
		loopNode(node.Right, level+1)
	}
}
