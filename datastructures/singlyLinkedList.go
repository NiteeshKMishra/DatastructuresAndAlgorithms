package datastructures

import "fmt"

func (s *SinglyLinkList) AddLast(value int64) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}
	if s.Length == 0 {
		s.Head = &newNode
		s.Tail = s.Head
	} else {
		s.Tail.Next = &newNode
		s.Tail = &newNode
	}
	s.Length++
}

func (s *SinglyLinkList) RemoveLast() {
	if s.Length == 0 {
		fmt.Println("No node to pop from linked list")
		return
	} else if s.Length == 1 {
		s.Head = nil
		s.Tail = nil
	} else {
		tempNode := s.Head
		for tempNode.Next != s.Tail {
			tempNode = tempNode.Next
		}
		tempNode.Next = nil
		s.Tail = tempNode
	}
	s.Length--
}

func (s *SinglyLinkList) InsertAt(index int, value int64) {
	if index < 0 || index > s.Length {
		fmt.Printf("Index %v is out of range", index)
		return
	}
	newNode := Node{
		Value: value,
		Next:  nil,
	}
	if s.Length == 0 { //index value could only be 0 here, insert at start
		s.Head = &newNode
		s.Tail = &newNode
	} else {
		tempNode := s.Head
		prevNode := s.Head
		currentIndex := 0
		for currentIndex < index {
			prevNode = tempNode
			tempNode = tempNode.Next
			currentIndex++
		}
		if tempNode == nil { //We have reached till tail, insert at end
			s.Tail.Next = &newNode
			s.Tail = &newNode
		} else { //insert in middle
			prevNode.Next = &newNode
			newNode.Next = tempNode
		}
	}
	s.Length++
}

func (s *SinglyLinkList) AddFirst(value int64) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}
	if s.Length == 0 {
		s.Head = &newNode
		s.Tail = &newNode
	} else {
		newNode.Next = s.Head
		s.Head = &newNode
	}
	s.Length++
}

func (s *SinglyLinkList) RemoveFirst() {
	if s.Length == 0 {
		fmt.Println("No node to shift from linked list")
		return
	} else if s.Length == 1 {
		s.Head = nil
		s.Tail = nil
	} else {
		tempNode := s.Head
		s.Head = tempNode.Next
		tempNode.Next = nil
	}
	s.Length--
}

func (s *SinglyLinkList) RemoveAt(index int) {
	if index < 0 || index > s.Length-1 {
		fmt.Printf("Index %v is out of range", index)
		return
	}
	if s.Length == 1 { //index can only be 0 here
		s.Head = nil
		s.Tail = nil
	} else {
		tempNode := s.Head
		prevNode := s.Head
		currentIndex := 0
		for currentIndex < index {
			prevNode = tempNode
			tempNode = tempNode.Next
			currentIndex++
		}
		if tempNode == s.Tail {
			prevNode.Next = nil
			s.Tail = prevNode
		} else {
			prevNode.Next = tempNode.Next
			tempNode.Next = nil
		}
	}
	s.Length--
}

func (s *SinglyLinkList) ListContains(value int64) bool {
	if s.Length == 0 {
		fmt.Println("Empty linked list")
		return false
	}
	tempNode := s.Head
	hasValue := false
	for tempNode != nil {
		if tempNode.Value == value {
			hasValue = true
			break
		}
		tempNode = tempNode.Next
	}
	return hasValue
}

func (s *SinglyLinkList) Reverse() {
	if s.Length == 0 {
		fmt.Println("Empty linked list")
		return
	}
	var curr = s.Head
	var prev *Node
	var next *Node
	for curr.Next != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr.Next = prev
	s.Tail = s.Head
	s.Head = curr
}
