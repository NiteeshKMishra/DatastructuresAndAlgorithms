package main

import (
	"fmt"

	"github.com/niteeshkrmishra/dsandalgo/algorithms"
	"github.com/niteeshkrmishra/dsandalgo/datastructures"
)

func main() {
	fmt.Println("Data structures")
	singlyLinkedList()
	doublyLinkedList()
	stackAndQueues()
	binarySearchTree()
	treeTraversal()
	binaryHeap()
	priorityQueue()
	hashTable()
	graph()
	graphTraversal()
	fmt.Println("Algorithms")
	graphAlgorithms()
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

func binarySearchTree() {
	fmt.Println("===============BinarySearchTree===================")
	bst := datastructures.BinarySearchTree{}
	//Insert
	bst.Insert(10)
	fmt.Println(bst.Root)
	bst.Insert(12)
	fmt.Println(bst.Root)
	bst.Insert(8)
	fmt.Println(bst.Root)
	bst.Insert(15)
	bst.Insert(9)
	fmt.Println(bst.Root)
	//Find
	hasValue := bst.Search(7)
	fmt.Println(bst.Root, hasValue)
	hasValue = bst.Search(9)
	fmt.Println(bst.Root, hasValue)
}

func treeTraversal() {
	fmt.Println("===============Tree Traversal===================")
	bst := datastructures.BinarySearchTree{}
	bst.Insert(10)
	bst.Insert(12)
	bst.Insert(8)
	bst.Insert(15)
	bst.Insert(9)
	bst.Insert(1)
	//BFS
	visited := bst.BFS()
	fmt.Println(bst.Root, visited)
	//DFS PreOrder
	visited = bst.DFSPreOrder()
	fmt.Println(bst.Root, visited)
	//DFS PostOrder
	visited = bst.DFSPostOrder()
	fmt.Println(bst.Root, visited)
	//DFS InOrder
	visited = bst.DFSInOrder()
	fmt.Println(bst.Root, visited)
}

func binaryHeap() {
	fmt.Println("===============Binary Heap===================")
	//Max binary heap
	bh := datastructures.BinaryHeap{}
	bh.MaxBinaryHeapInsert(41)
	bh.MaxBinaryHeapInsert(39)
	bh.MaxBinaryHeapInsert(33)
	bh.MaxBinaryHeapInsert(18)
	bh.MaxBinaryHeapInsert(27)
	bh.MaxBinaryHeapInsert(12)
	bh.MaxBinaryHeapInsert(55)
	fmt.Println(bh.Values)

	//MaxValue Extract
	value := bh.BinaryHeapExtractMax()
	fmt.Println(value, bh.Values)
	value = bh.BinaryHeapExtractMax()
	fmt.Println(value, bh.Values)
}

func priorityQueue() {
	fmt.Println("===============Priority Queue===================")
	p := datastructures.PriorityQueue{}
	p.Enqueue("Red", 1)
	p.Enqueue("Blue", 4)
	p.Enqueue("Orange", 5)
	p.Enqueue("Yellow", 1)
	p.Enqueue("Black", 3)
	p.Enqueue("White", 2)
	p.Enqueue("Grey", 4)
	fmt.Println(p.Values)

	value := p.Dequeue()
	fmt.Println(value, p.Values)
	value = p.Dequeue()
	fmt.Println(value, p.Values)
}

func hashTable() {
	fmt.Println("===============Hash Table===================")
	h := datastructures.HashTable{}
	h.Values = make([][][]string, datastructures.HashTableLength)
	//Set Values
	h.Set("red", "#FF3F33")
	h.Set("blue", "#3352FF")
	h.Set("green", "#33FF52")
	h.Set("red", "#FF5533")
	fmt.Println(h.Values)
	//Get values
	value := h.Get("red")
	fmt.Println(value)
	value = h.Get("orange")
	fmt.Println(value)
	//Get Keys
	keys := h.Keys()
	fmt.Println(keys)
	//Get KeyValues
	keyValues := h.KeyValues()
	fmt.Println(keyValues)
}

func graph() {
	fmt.Println("===============Graph===================")
	g := datastructures.Graph{}
	g.AdjacencyList = make(map[string][]string)
	//Add vertex
	g.AddVertex("Delhi")
	g.AddVertex("Mumbai")
	fmt.Println(g.AdjacencyList)
	//Add edge
	g.AddEdge("Delhi", "Mumbai")
	g.AddEdge("Delhi", "Goa")
	fmt.Println(g.AdjacencyList)
	//Remove edge
	g.RemoveEdge("Delhi", "Goa")
	fmt.Println(g.AdjacencyList)
	//Remove vertex
	g.RemoveVertex("Mumbai")
	fmt.Println(g.AdjacencyList)
}

func graphTraversal() {
	fmt.Println("===============Graph Traversal===================")
	g := datastructures.Graph{}
	g.AdjacencyList = make(map[string][]string)
	g.AddEdge("Delhi", "Mumbai")
	g.AddEdge("Delhi", "Goa")
	g.AddEdge("Goa", "Kolkata")
	g.AddEdge("Mumbai", "Chennai")
	fmt.Println(g.AdjacencyList)
	visited := g.DFSRecursive("Delhi")
	fmt.Println(visited)
	visited = g.DFSIterative("Delhi")
	fmt.Println(visited)
	visited = g.BFS("Delhi")
	fmt.Println(visited)
}

func graphAlgorithms() {
	fmt.Println("===============Graph Algorithms===================")
	//Dijkstra
	g := datastructures.WeightedGraph{}
	g.AdjacencyList = make(map[string][]datastructures.WeightedGraphNode)
	g.AddEdge("Delhi", "Mumbai", 720)
	g.AddEdge("Delhi", "Nagpur", 560)
	g.AddEdge("Nagpur", "Mumbai", 100)
	g.AddEdge("Delhi", "Goa", 450)
	g.AddEdge("Goa", "Mumbai", 100)
	path := algorithms.ShortestPath(&g, "Delhi", "Mumbai")
	fmt.Println(path)
}
