package main

import (
	"fmt"
)

func main() {
    // Make a list from an array of values.
    values := []string {
        "0", "1", "2", "3", "4", "5",
    }
    list := makeLinkedList()
    list.addRange(values)

    fmt.Println(list.toString(" "))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 5 point to cell 2.
    list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

    fmt.Println(list.toStringMax(" ", 10))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 4 point to cell 2.
    list.sentinel.next.next.next.next.next = list.sentinel.next.next

    fmt.Println(list.toStringMax(" ", 10))
    if list.hasLoop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
}

