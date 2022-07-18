package datastructures

func (b *BinarySearchTree) BFS() []int64 {
	queue := []*TreeNode{}
	visited := []int64{}

	if b.Root != nil {
		queue = append(queue, b.Root)
		for len(queue) != 0 {
			currentNode := queue[0]
			queue = queue[1:]
			visited = append(visited, currentNode.Value)
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}

	return visited
}

func (b *BinarySearchTree) DFSPreOrder() []int64 {
	var visted = []int64{}
	var visitNode func(node *TreeNode) 
	visitNode = func (node *TreeNode) {
		visted = append(visted, node.Value)
		if(node.Left != nil){
			visitNode(node.Left)
		}
		if(node.Right != nil){
			visitNode(node.Right)
		}
	}
	if b.Root != nil {
		visitNode(b.Root)
	}
	return visted
}

func (b *BinarySearchTree) DFSInOrder() []int64 {
	var visited = []int64{}
	var visitNode func(node *TreeNode)

	visitNode = func(node *TreeNode) {
		if node.Left != nil {
			visitNode(node.Left)
		}
		visited = append(visited, node.Value)
		if node.Right != nil {
			visitNode(node.Right)
		}
	}

	if b.Root != nil {
		visitNode(b.Root)
	}
	return visited
}

func (b *BinarySearchTree) DFSPostOrder() []int64 {
	var visited = []int64{}
	var visitNode func(node *TreeNode)

	visitNode = func(node *TreeNode) {
		if node.Left != nil {
			visitNode(node.Left)
		}
		if node.Right != nil {
			visitNode(node.Right)
		}
		visited = append(visited, node.Value)
	}

	if b.Root != nil {
		visitNode(b.Root)
	}
	return visited
}
