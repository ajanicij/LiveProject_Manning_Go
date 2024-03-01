package main

import (
	"fmt"
	"strings"
)

// queue

type Cell [T any] struct {
	data T
	next *Cell [T]
	prev *Cell [T]
}

type Queue [T any] struct {
	topSentinel *Cell [T]
	bottomSentinel *Cell [T]
}

func (queue *Queue [T]) isEmpty() bool {
	return queue.topSentinel.next == queue.bottomSentinel
}

func (me *Cell [T]) addAfter(after *Cell [T]) {
	after.next = me.next
	after.prev = me
	me.next.prev = after
	me.next = after
}

func (me *Cell [T]) addBefore(before *Cell [T]) {
	prev := me.prev
	prev.addAfter(before)
}

func (me *Cell [T]) delete() {
	next := me.next
	prev := me.prev
	prev.next = next
	next.prev = prev
}

func (queue *Queue [T]) push(node T) {
	cell := &Cell [T]{data: node}
	queue.topSentinel.addAfter(cell)
}

func (queue *Queue [T]) pop() T {
	if queue.isEmpty() {
		panic("trying to pop from empty stack")
	}
	cell := queue.topSentinel.next
	cell.delete()
	return cell.data
}

func MakeQueue [T any]() Queue [T] {
	topSentinel := &Cell [T]{}
	bottomSentinel := &Cell [T]{}
	topSentinel.prev = nil
	topSentinel.next = bottomSentinel
	bottomSentinel.prev = topSentinel
	bottomSentinel.next = nil
	return Queue [T] {
		topSentinel: topSentinel,
		bottomSentinel: bottomSentinel,
	}
}

// Add an item to the top of the queue.
func (queue *Queue [T]) enqueue(value T) {
	queue.push(value)
}

// Remove an item from the bottom of the queue.
func (queue *Queue [T]) dequeue() T {
	if queue.isEmpty() {
		panic("trying to dequeue from empty queue")
	}
	cell := queue.bottomSentinel.prev
	cell.delete()
	return cell.data
}

// tree

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

func buildTree() *Node {
	iTree := NewTree("I", nil, nil)
	jTree := NewTree("J", nil, nil)
	hTree := NewTree("H", iTree, jTree)
	fTree := NewTree("F", hTree, nil)
	cTree := NewTree("C", nil, fTree)
	gTree := NewTree("G", nil, nil)
	dTree := NewTree("D", nil, nil)
	eTree := NewTree("E", gTree, nil)
	bTree := NewTree("B", dTree, eTree)
	aTree := NewTree("A", bTree, cTree)
	return aTree
}

func indentString(value string, indent string, depth int) string {
	builder := strings.Builder{}
	for i := 0; i < depth; i++ {
		builder.WriteString(indent)
	}
	builder.WriteString(value)
	return builder.String()
}

func (node *Node) displayIndented(indent string, depth int) string {
	result := indentString(node.data, indent, depth) + "\n"
	if node.left != nil {
		result += node.left.displayIndented(indent, depth + 1)
	}
	if node.right != nil {
		result += node.right.displayIndented(indent, depth + 1)
	}
	return result
}

func main() {
    // Build a tree.
    aNode := buildTree()
    
    // Display with indentation.
    fmt.Println(aNode.displayIndented("  ", 0))

    // Display traversals.
    fmt.Println("Preorder:     ", aNode.preorder())
    fmt.Println("Inorder:      ", aNode.inorder())
    fmt.Println("Postorder:    ", aNode.postorder())
    fmt.Println("Breadth first:", aNode.breadthFirst())
}

