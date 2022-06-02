package main

import (
	"fmt"

	"github.com/niteeshkrmishra/dsandalgo/datastructures"
)

func main() {
	fmt.Println("Data structures")
	singlyLinkedList()
	doublyLinkedList()
	stackAndQueues()
}

func singlyLinkedList() {
	fmt.Println("===============SinglyLinkedList===================")
	singlyLinkedList := datastructures.SinglyLinkList{}
	//Push
	singlyLinkedList.AddFirst(23)
	//Insert
	singlyLinkedList.InsertAt(1, 25)
	fmt.Println(*singlyLinkedList.Head, *singlyLinkedList.Tail, singlyLinkedList.Length)
	//Unshift
	singlyLinkedList.AddLast(28)
	fmt.Println(*singlyLinkedList.Head, *singlyLinkedList.Tail, singlyLinkedList.Length)
	//Contains
	hasValue := singlyLinkedList.ListContains(25)
	fmt.Printf("Does Linked list has value %v: %v\n", 25, hasValue)
	//Reverse
	singlyLinkedList.Reverse()
	fmt.Println(*singlyLinkedList.Head, *singlyLinkedList.Tail, singlyLinkedList.Length)
	//Pop
	singlyLinkedList.RemoveLast()
	//Shift
	singlyLinkedList.RemoveFirst()
	fmt.Println(*singlyLinkedList.Head, *singlyLinkedList.Tail, singlyLinkedList.Length)
	//RemoveAt
	singlyLinkedList.RemoveAt(0)
	fmt.Println(singlyLinkedList, singlyLinkedList.Length)
}

func doublyLinkedList() {
	fmt.Println("===============DoublyLinkedList===================")
	doublyLinkedList := datastructures.DoublyLinkList{}
	//Push
	doublyLinkedList.AddFirst(23)
	//Insert
	doublyLinkedList.InsertAt(1, 25)
	fmt.Println(*doublyLinkedList.Head, *doublyLinkedList.Tail, doublyLinkedList.Length)
	//Unshift
	doublyLinkedList.AddLast(28)
	fmt.Println(*doublyLinkedList.Head, *doublyLinkedList.Tail, doublyLinkedList.Length)
	//Contains
	hasValue := doublyLinkedList.ListContains(25)
	fmt.Printf("Does Linked list has value %v: %v\n", 25, hasValue)
	//Reverse
	doublyLinkedList.Reverse()
	fmt.Println(*doublyLinkedList.Head, *doublyLinkedList.Tail, doublyLinkedList.Length)
	//Pop
	doublyLinkedList.RemoveLast()
	//Shift
	doublyLinkedList.RemoveFirst()
	fmt.Println(*doublyLinkedList.Head, *doublyLinkedList.Tail, doublyLinkedList.Length)
	//RemoveAt
	doublyLinkedList.RemoveAt(0)
	fmt.Println(doublyLinkedList, doublyLinkedList.Length)
}

func stackAndQueues() {
	fmt.Println("===============Stack===================")
	stack := datastructures.Stack{}
	queue := datastructures.Queue{}
	//Stack Push
	stack.Push(10)
	stack.Push(987980089789790898)
	fmt.Println(*stack.Top, stack.Length)
	//Stack Pop
	value := stack.Pop()
	fmt.Println(value, *stack.Top, stack.Length)
	stack.Pop()
	stack.Pop()
	fmt.Println("===============Queue===================")
	//Queue Enqueue
	queue.Enqueue(100)
	queue.Enqueue(200)
	fmt.Println(*queue.First, queue.Length)
	//Queue Dequeue
	value = queue.Dequeue()
	fmt.Println(value, *queue.First, queue.Length)
	value = queue.Dequeue()
	fmt.Println(value, queue, queue.Length)
	queue.Dequeue()
}
