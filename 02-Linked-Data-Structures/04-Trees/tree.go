package main

type Node struct {
	data  string
	left  *Node
	right *Node
}

func NewTree(data string, left, right *Node) *Node {
	root := Node { data: data, left: left, right: right }
	return &root
}

func (node *Node) preorder() string {
	result := node.data
	if node.left != nil {
		result += " " + node.left.preorder()
	}
	if node.right != nil {
		result += " " + node.right.preorder()
	}
	return result
}

func (node *Node) preorderFunc(fn func(node *Node)) string {
	fn(node)
	result := node.data
	if node.left != nil {
		result += " " + node.left.preorderFunc(fn)
	}
	if node.right != nil {
		result += " " + node.right.preorderFunc(fn)
	}
	return result
}

func (node *Node) traverse(fn func(node *Node)) {
	fn(node)
	if node.left != nil {
		node.left.traverse(fn)
	}
	if node.right != nil {
		node.right.traverse(fn)
	}
}

func (node *Node) inorder() string {
	result := ""
	if node.left != nil {
		result += node.left.inorder() + " "
	}
	result += node.data
	if node.right != nil {
		result += " " + node.right.inorder()
	}
	return result
}
func (node *Node) postorder() string {
	result := ""
	if node.left != nil {
		result += node.left.postorder() + " "
	}
	if node.right != nil {
		result += node.right.postorder() + " "
	}
	result += node.data
	return result
}

func (node *Node) breadthFirst() string {
	result := ""
	queue := MakeQueue [*Node]()
	queue.enqueue(node)
	for !queue.isEmpty() {
		curr := queue.dequeue()
		result += curr.data
		if curr.left != nil {
			queue.push(curr.left)
		}
		if curr.right != nil {
			queue.push(curr.right)
		}
		if !queue.isEmpty() {
			result += " "
		}
	}
	return result
}

