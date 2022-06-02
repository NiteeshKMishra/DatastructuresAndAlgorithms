package datastructures

import "fmt"

func (b *BinarySearchTree) Insert(value int64) {
	newNode := TreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}

	if b.Root == nil {
		b.Root = &newNode
		return
	}
	currentNode := b.Root
	for currentNode != nil {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = &newNode
				break
			}
			currentNode = currentNode.Left
		} else {
			if currentNode.Right == nil {
				currentNode.Right = &newNode
				break
			}
			currentNode = currentNode.Right
		}
	}
}

func (b *BinarySearchTree) Search(value int64) bool {
	hasValue := false

	if b.Root == nil {
		fmt.Println("Empty Binary search tree")
		return hasValue
	}
	tempNode := b.Root
	for tempNode != nil {
		if value == tempNode.Value {
			hasValue = true
			break
		} else if value < tempNode.Value {
			tempNode = tempNode.Left
		} else {
			tempNode = tempNode.Right
		}
	}
	return hasValue
}
