package datastructures

import "fmt"

func (q *Queue) Enqueue(value int64) int {
	newNode := Node{
		Value: value,
		Next:  nil,
	}
	if q.Length == 0 {
		q.First = &newNode
		q.Last = &newNode
	} else {
		q.Last.Next = &newNode
		q.Last = &newNode
	}
	q.Length++
	return q.Length
}

func (q *Queue) Dequeue() int64 {
	if q.Length == 0 {
		fmt.Println("Queue is Empty")
		return 0
	} else {
		temp := q.First
		q.First = temp.Next
		temp.Next = nil
		q.Length--
		return temp.Value
	}
}
