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

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rotateLeft(node *Node) *Node {
	rightNode := node.Right
	rightLeftNode := rightNode.Left

	rightNode.Left = node
	node.Right = rightLeftNode

	node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))
	rightNode.Height = 1 + max(getHeight(rightNode.Left), getHeight(rightNode.Right))

	return rightNode
}

func rotateRight(node *Node) *Node {
	leftNode := node.Left
	leftRightNode := leftNode.Right

	leftNode.Right = node
	node.Left = leftRightNode

	node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))
	leftNode.Height = 1 + max(getHeight(leftNode.Left), getHeight(leftNode.Right))

	return leftNode
}

func insert(node *Node, value int) *Node {
	if node == nil {
		node = &Node{Value: value, Height: 1}
		return node
	} else if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	} else {
		return node
	}

	node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))

	balance := getHeight(node.Left) - getHeight(node.Right)

	if balance > 1 && value < node.Left.Value {
		return rotateRight(node)
	} else if balance < -1 && value > node.Right.Value {
		return rotateLeft(node)
	} else if balance > 1 && value > node.Left.Value {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	} else if balance < -1 && value < node.Right.Value {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}

func (tree *avlTree) Insert(value int) {
	tree.Root = insert(tree.Root, value)
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
	fmt.Printf("node %p \tvalue: %v \theight: %v \tleft: %p \tright: %p\n", node, node.Value, node.Height, node.Left, node.Right)
}
