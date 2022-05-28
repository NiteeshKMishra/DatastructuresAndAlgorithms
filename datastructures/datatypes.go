package datastructures

type Node struct {
	Value int64
	Next  *Node
}

type SinglyLinkList struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Stack struct {
	Top    *Node
	Length int
}

type Queue struct {
	First  *Node
	Last   *Node
	Length int
}
