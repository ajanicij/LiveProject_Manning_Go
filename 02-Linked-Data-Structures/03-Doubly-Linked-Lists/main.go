package main

import (
	"fmt"
)

func main() {
    // Make a list from a slice of values.
    list := makeDoublyLinkedList()
    animals := []string {
        "Ant",
        "Bat",
        "Cat",
        "Dog",
        "Elk",
        "Fox",
    }
    list.addRange(animals)
    fmt.Println(list.toString(" "))
}

