package main

import (
	"fmt"
	"strings"
)

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
    
    // Functional programming.
    fmt.Println("Functional Traversal")
    aNode.traverse(func (node *Node) {
    	fmt.Printf("func: (%p) node=%v\n", node, node)
    })
}

