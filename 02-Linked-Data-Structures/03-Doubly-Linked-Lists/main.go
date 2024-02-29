package main

import (
	"fmt"
)

func main() {
    // Test queue functions.
    fmt.Printf("*** Queue Functions ***\n")
    queue := makeDoublyLinkedList()
    queue.enqueue("Agate")
    queue.enqueue("Beryl")
    fmt.Printf("%s ", queue.dequeue())
    queue.enqueue("Citrine")
    fmt.Printf("%s ", queue.dequeue())
    fmt.Printf("%s ", queue.dequeue())
    queue.enqueue("Diamond")
    queue.enqueue("Emerald")
    for !queue.isEmpty() {
        fmt.Printf("%s ", queue.dequeue())
    }
    fmt.Printf("\n\n")

    // Test deque functions. Names starting
    // with F have a fast pass.
    fmt.Printf("*** Deque Functions ***\n")
    deque := makeDoublyLinkedList()
    deque.pushTop("Ann")
    deque.pushTop("Ben")
    fmt.Printf("%s ", deque.popBottom())
    deque.pushBottom("F-Cat")
    fmt.Printf("%s ", deque.popBottom())
    fmt.Printf("%s ", deque.popBottom())
    deque.pushBottom("F-Dan")
    deque.pushTop("Eva")
    for !deque.isEmpty() {
        fmt.Printf("%s ", deque.popBottom())
    }
    fmt.Printf("\n")
}

