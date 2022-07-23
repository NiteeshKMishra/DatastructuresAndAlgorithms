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

type TreeNode struct {
	Value int64
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

type BinaryHeap struct {
	Values []int64
}

type PriorityNode struct {
	Value    int64
	Priority int
}

type PriorityQueue struct {
	Values []PriorityNode
}

type HashTable struct {
	Values [][][]string
}

type Graph struct {
	AdjacencyList map[string][]string
}

type WeightedGraphNode struct {
	Vertex string
	Weight int64
}

type WeightedGraph struct {
	AdjacencyList map[string][]WeightedGraphNode
}
