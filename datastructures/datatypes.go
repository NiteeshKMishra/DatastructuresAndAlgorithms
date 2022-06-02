package datastructures

type Node struct {
	Value int64
	Next  *Node
}

type DoublePtrNode struct {
	Value    int64
	Next     *DoublePtrNode
	Previous *DoublePtrNode
}

type SinglyLinkList struct {
	Head   *Node
	Tail   *Node
	Length int
}

type DoublyLinkList struct {
	Head   *DoublePtrNode
	Tail   *DoublePtrNode
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
