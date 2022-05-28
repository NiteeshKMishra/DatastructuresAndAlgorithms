package datastructures

import "fmt"

func (s *Stack) Push(value int64) int {
	newNode := Node{
		Value: value,
		Next:  nil,
	}

	newNode.Next = s.Top
	s.Top = &newNode
	s.Length++
	return s.Length
}

func (s *Stack) Pop() int64 {
	if s.Length == 0 {
		fmt.Println("Stack is empty")
		return 0
	} else {
		temp := s.Top
		s.Top = s.Top.Next
		temp.Next = nil
		s.Length--
		return temp.Value
	}
}
