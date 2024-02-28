package main

import (
	"fmt"
)

func main() {
    // smallListTest()

    // Make a list from an array of values.
    greekLetters := []string {
        "α", "β", "γ", "δ", "ε",
    }
    list := makeLinkedList()
    list.addRange(greekLetters)
    fmt.Println(list.toString(" "))
    fmt.Println()

    // Demonstrate a stack.
    stack := makeLinkedList()
    stack.push("Apple")
    stack.push("Banana")
    stack.push("Coconut")
    stack.push("Date")
    for !stack.isEmpty() {
        fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
            stack.pop(),
            stack.length(),
            stack.toString(" "))
    }
}

