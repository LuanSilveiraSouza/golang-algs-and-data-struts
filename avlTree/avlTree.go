package avlTree

import (
	"errors"
	"fmt"
)

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

type avlTree struct {
	Root *Node
}

func NewAvlTree() avlTree {
	return avlTree{nil}
}

func getMaxHeight(node *Node) int {
	if node.Left.Height > node.Right.Height {
		return node.Left.Height
	}
	return node.Right.Height
}

func balanceTree(currentNode *Node, value int) {
	currentNode.Height = 1 + getMaxHeight(currentNode)

	balance := currentNode.Left.Height - currentNode.Right.Height

	if balance > 1 && value < currentNode.Left.Value {
		// Left rotation
	} else if balance < -1 && value > currentNode.Right.Value {
		// Right rotation
	} else if balance > 1 && value > currentNode.Left.Value {
		// Left Right rotation
	} else if balance < -1 && value < currentNode.Right.Value {
		// Right Left rotation
	}
}

func (tree *avlTree) Insert(value int) {
	newNode := Node{value, nil, nil, 0}

	if tree.Root == nil {
		tree.Root = &newNode
	} else {
		currentNode := tree.Root
		for {
			if value < currentNode.Value {
				if currentNode.Left == nil {
					currentNode.Left = &newNode

					balanceTree(currentNode, value)

					break
				}
				currentNode = currentNode.Left
				continue
			} else if value > currentNode.Value {
				if currentNode.Right == nil {
					currentNode.Right = &newNode

					balanceTree(currentNode, value)

					break
				}
				currentNode = currentNode.Right
				continue
			}
			break
		}
	}
}

func (tree *avlTree) Search(value int) (*Node, error) {
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

func (tree *avlTree) InOrderTraversal(root *Node) {
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
