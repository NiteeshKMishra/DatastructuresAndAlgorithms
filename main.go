package main

import (
	"fmt"

	"github.com/niteeshkrmishra/dsandalgo/datastructures"
)

func main() {
	fmt.Println("Data structures")
	singlyLinkedList()

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
