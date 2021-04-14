package redBlackTree

import (
	"errors"
	"fmt"
)

type Color bool

const (
	black, red Color = true, false
)

func GetColor(color Color) string {
	if color {
		return "black"
	}
	return "red"
}

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
	color  Color
}

type RBTree struct {
	root *Node
}

func NewRBTree() RBTree {
	return RBTree{nil}
}

func (tree *RBTree) rotateLeft(node *Node) {
	rightNode := node.right

	node.right = rightNode.left

	if rightNode.left != nil {
		rightNode.left.parent = node
	}

	rightNode.parent = node.parent

	if node.parent == nil {
		tree.root = rightNode
	} else {
		if node == node.parent.left {
			node.parent.left = rightNode
		} else {
			node.parent.right = rightNode
		}
	}

	rightNode.left = node
	node.parent = rightNode
}

func (tree *RBTree) rotateRight(node *Node) {
	leftNode := node.left

	node.left = leftNode.right

	if leftNode.right != nil {
		leftNode.right.parent = node
	}

	leftNode.parent = node.parent

	if node.parent == nil {
		tree.root = leftNode
	} else {
		if node == node.parent.right {
			node.parent.right = leftNode
		} else {
			node.parent.left = leftNode
		}
	}

	leftNode.right = node
	node.parent = leftNode
}

func (tree *RBTree) rebalance(node *Node) {
	var uncle *Node

	for node.parent != nil && node.parent.color == red {
		if node.parent == node.parent.parent.left {
			uncle = node.parent.parent.right

			if uncle != nil && uncle.color == red {
				node.parent.color = black
				uncle.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					tree.rotateLeft(node)
				}
				node.parent.color = black
				node.parent.parent.color = red
				tree.rotateRight(node.parent.parent)
			}
		} else {
			uncle = node.parent.parent.left

			if uncle != nil && uncle.color == red {
				node.parent.color = black
				uncle.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					tree.rotateRight(node)
				}
				node.parent.color = black
				node.parent.parent.color = red
				tree.rotateLeft(node.parent.parent)

			}
		}
	}

	tree.root.color = black
}

func (tree *RBTree) Insert(value int) {

	if tree.root == nil {
		tree.root = &Node{value: value, color: black}
	} else {
		node := tree.root

		for {
			if value < node.value {
				if node.left == nil {
					node.left = &Node{value: value, color: red, parent: node}
					tree.rebalance(node.left)
					break
				} else {
					node = node.left
					continue
				}
			} else if value > node.value {
				if node.right == nil {
					node.right = &Node{value: value, color: red, parent: node}
					tree.rebalance(node.right)
					break
				} else {
					node = node.right
					continue
				}
			}
			break
		}
	}
}

func (tree *RBTree) Search(value int) (*Node, error) {
	currentNode := tree.root

	for {
		if currentNode == nil {
			return nil, errors.New("value dont exists in the tree")
		} else {
			if currentNode.value == value {
				return currentNode, nil
			} else if value < currentNode.value {
				currentNode = currentNode.left
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

func (tree *RBTree) InOrderTraversal(root *Node) {
	if root != nil {
		if root.left != nil {
			tree.InOrderTraversal(root.left)
		}
		PrintNode(root)
		if root.right != nil {
			tree.InOrderTraversal(root.right)
		}
	}
}

func PrintNode(node *Node) {
	fmt.Printf("node %p \tvalue: %v \tleft: %p \tright: %p \tparent: %p \tcolor: %v \n", node, node.value, node.left, node.right, node.parent, GetColor(node.color))
}
