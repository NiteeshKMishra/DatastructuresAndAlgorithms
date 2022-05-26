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
