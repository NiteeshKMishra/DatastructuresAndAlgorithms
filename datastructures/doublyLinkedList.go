package datastructures

import (
	"fmt"
)

func (d *DoublyLinkList) AddLast(value int64) {
	newNode := DoublePtrNode{
		Value:    value,
		Next:     nil,
		Previous: nil,
	}

	if d.Length == 0 {
		d.Head = &newNode
		d.Tail = &newNode
	} else {
		newNode.Previous = d.Tail
		d.Tail.Next = &newNode
		d.Tail = &newNode
	}
	d.Length++
}

func (d *DoublyLinkList) RemoveLast() int64 {
	var value int64 = 0

	if d.Length == 0 {
		fmt.Println("Linked list is empty")
	} else if d.Length == 1 {
		value = d.Head.Value
		d.Head = nil
		d.Tail = nil
		d.Length--
	} else {
		tempNode := d.Head
		for tempNode.Next != d.Tail {
			tempNode = tempNode.Next
		}
		value = d.Tail.Value
		d.Tail.Previous = nil
		tempNode.Next = nil
		d.Tail = tempNode
		d.Length--
	}
	return value
}

func (d *DoublyLinkList) AddFirst(value int64) {
	newNode := DoublePtrNode{
		Value:    value,
		Next:     nil,
		Previous: nil,
	}

	if d.Length == 0 {
		d.Head = &newNode
		d.Tail = &newNode
	} else {
		newNode.Next = d.Head
		d.Head.Previous = &newNode
		d.Head = &newNode
	}
	d.Length++
}

func (d *DoublyLinkList) RemoveFirst() int64 {
	var value int64 = 0

	if d.Length == 0 {
		fmt.Println("Linked list is empty")
	} else if d.Length == 1 {
		value = d.Head.Value
		d.Head = nil
		d.Tail = nil
		d.Length--
	} else {
		value = d.Head.Value
		tempNode := d.Head.Next
		tempNode.Previous = nil
		d.Head.Next = nil
		d.Head = tempNode
		d.Length--
	}
	return value
}

func (d *DoublyLinkList) InsertAt(index int, value int64) {
	if index < 0 || index > d.Length {
		fmt.Printf("Index %v is out of range", index)
		return
	}
	newNode := DoublePtrNode{
		Value:    value,
		Next:     nil,
		Previous: nil,
	}
	if index == 0 {
		newNode.Next = d.Head
		if d.Length == 0 {
			d.Tail = &newNode
		} else {
			d.Head.Previous = &newNode
		}
		d.Head = &newNode
	} else if index == d.Length {
		newNode.Previous = d.Tail
		d.Tail.Next = &newNode
		d.Tail = &newNode
	} else {
		currIndex := 0
		prevNode := d.Head
		currNode := d.Head
		for currIndex != index {
			prevNode = currNode
			currNode = currNode.Next
			currIndex++
		}
		prevNode.Next = &newNode
		newNode.Previous = prevNode
		currNode.Previous = &newNode
		newNode.Next = currNode
	}
	d.Length++
}

func (d *DoublyLinkList) RemoveAt(index int) int64 {
	var value int64 = 0

	if index < 0 || index >= d.Length {
		fmt.Printf("Index %v is out of range", index)
		return value
	}
	if index == 0 {
		value = d.Head.Value
		tempNode := d.Head.Next
		d.Head.Next = nil
		if d.Length == 1 {
			d.Tail = nil
		} else {
			tempNode.Previous = nil
		}
		d.Head = tempNode
	} else if index == d.Length {
		value = d.Tail.Value
		tempNode := d.Head
		for tempNode.Next != d.Tail {
			tempNode = tempNode.Next
		}
		d.Tail.Previous = nil
		tempNode.Next = nil
		d.Tail = tempNode
	} else {
		currIndex := 0
		prevNode := d.Head
		currNode := d.Head
		for currIndex != index {
			prevNode = currNode
			currNode = currNode.Next
			currIndex++
		}
		currNode.Next.Previous = prevNode
		prevNode.Next = currNode.Next
		currNode.Next = nil
		currNode.Previous = nil
	}
	d.Length--
	return value
}

func (d *DoublyLinkList) ListContains(value int64) bool {
	hasValue := false

	tempNode := d.Head
	for tempNode != nil {
		if tempNode.Value == value {
			hasValue = true
			break
		} else {
			tempNode = tempNode.Next
		}
	}
	return hasValue
}

func (d *DoublyLinkList) Reverse() {
	if d.Length < 1 {
		fmt.Println("Linked list is empty or has single node")
		return
	} else {
		var currNode *DoublePtrNode = d.Head
		var prevNode *DoublePtrNode
		var nextNode *DoublePtrNode

		for currNode.Next != nil {
			nextNode = currNode.Next
			currNode.Next = prevNode
			currNode.Previous = nextNode
			prevNode = currNode
			currNode = nextNode
		}
		currNode.Next = prevNode
		currNode.Previous = nil
		d.Tail = d.Head
		d.Head = currNode
	}
}
