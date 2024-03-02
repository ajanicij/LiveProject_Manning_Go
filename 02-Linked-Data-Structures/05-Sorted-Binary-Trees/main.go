package main

import (
	"fmt"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func (tree *Node) insertValue(value string) {
	if value < tree.data {
		if tree.left == nil {
			newNode := &Node{value, nil, nil}
			tree.left = newNode
			return
		} else {
			tree.left.insertValue(value)
		}
	} else if value > tree.data {
		if tree.right == nil {
			newNode := &Node{value, nil, nil}
			tree.right = newNode
			return
		} else {
			tree.right.insertValue(value)
		}
	}
}

func (tree *Node) findValue(value string) *Node {
	if value < tree.data {
		if tree.left == nil {
			return nil
		} else {
			return tree.left.findValue(value)
		}
	} else if value > tree.data {
		if tree.right == nil {
			return nil
		} else {
			return tree.right.findValue(value)
		}
	} else {
		return tree
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

func main() {
    // Make a root node to act as sentinel.
    root := Node { "", nil, nil }

    // Add some values.
    root.insertValue("I")
    root.insertValue("G")
    root.insertValue("C")
    root.insertValue("E")
    root.insertValue("B")
    root.insertValue("K")
    root.insertValue("S")
    root.insertValue("Q")
    root.insertValue("M")

    // Add F.
    root.insertValue("F")

    // Display the values in sorted order.
    fmt.Printf("Sorted values: %s\n", root.right.inorder())

    // Let the user search for values.
    for {
        // Get the target value.
        target := ""
        fmt.Printf("String: ")
        fmt.Scanln(&target)
        if len(target) == 0 { break }

        // Find the value's node.
        node := root.findValue(target)
        if node == nil {
            fmt.Printf("%s not found\n", target)
        } else {
            fmt.Printf("Found value %s\n", target)
        }
    }
}

