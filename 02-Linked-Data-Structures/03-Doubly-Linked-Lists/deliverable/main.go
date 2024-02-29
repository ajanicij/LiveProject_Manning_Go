package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	data string
	next *Cell
	prev *Cell
}

type DoublyLinkedList struct {
	topSentinel *Cell
	bottomSentinel *Cell
}

func makeDoublyLinkedList() DoublyLinkedList {
	topSentinel := &Cell{}
	bottomSentinel := &Cell{}
	topSentinel.next = bottomSentinel
	bottomSentinel.prev = topSentinel
	return DoublyLinkedList{
		topSentinel: topSentinel,
		bottomSentinel: bottomSentinel,
	}
}

func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	after.prev = me
	me.next.prev = after
	me.next = after
}

func (me *Cell) addBefore(before *Cell) {
	prev := me.prev
	prev.addAfter(before)
}

func (me *Cell) delete() {
	next := me.next
	prev := me.prev
	prev.next = next
	next.prev = prev
}

func (list *DoublyLinkedList) addRange(values []string) {
	for _, value := range values {
		cell := &Cell{data: value}
		list.bottomSentinel.addBefore(cell)
	}
}

func (list *DoublyLinkedList) toString(separator string) string {
	builder := strings.Builder{}
	for cell := list.topSentinel.next; cell != list.bottomSentinel; cell = cell.next {
		builder.WriteString(cell.data)
		if cell.next != list.bottomSentinel {
			builder.WriteString(separator)
		}
	}
	return builder.String()
}

func (list *DoublyLinkedList) length() int {
	l := 0
	cell := list.topSentinel
	for {
		cell = cell.next
		if cell == list.bottomSentinel {
			break
		}
		l++
	}
	return l
}

func (list *DoublyLinkedList) isEmpty() bool {
	return list.topSentinel.next == list.bottomSentinel
}

func (list *DoublyLinkedList) contains(value string) bool {
	cell := list.topSentinel.next
	for cell != list.bottomSentinel {
		if cell.data == value {
			return true
		}
		cell = cell.next
	}
	return false
}

func (list *DoublyLinkedList) find(value string) *Cell {
	cell := list.topSentinel.next
	for cell != list.bottomSentinel {
		if cell.data == value {
			return cell
		}
		cell = cell.next
	}
	return nil
}

func (list *DoublyLinkedList) push(value string) {
	cell := &Cell{data: value}
	list.topSentinel.addAfter(cell)
}

func (list *DoublyLinkedList)  pop() string {
	if list.isEmpty() {
		panic("trying to pop from empty stack")
	}
	cell := list.topSentinel.next
	cell.delete()
	return cell.data
}

func (list *DoublyLinkedList) hasLoop() bool {
	fast := list.topSentinel.next
	slow := fast
	for {
		fast = fast.next
		if fast == list.bottomSentinel {
			break
		}
		fast = fast.next
		if fast == list.bottomSentinel {
			break
		}
		slow = slow.next
		if fast == slow {
			// Loop detected.
			return true
		}
	}
	// End of list reached.
	return false
}

func (list *DoublyLinkedList) toStringMax(separator string, max int) string {
	builder := strings.Builder{}
	count := 1
	for cell := list.topSentinel.next; cell != list.bottomSentinel; cell = cell.next {
		builder.WriteString(cell.data)
		if count >= max {
			break
		}
		if cell.next != list.bottomSentinel {
			builder.WriteString(separator)
		}
		count++
	}
	return builder.String()
}

// queue

// Add an item to the top of the queue.
func (queue *DoublyLinkedList) enqueue(value string) {
	queue.push(value)
}

// Remove an item from the bottom of the queue.
func (queue *DoublyLinkedList) dequeue() string {
	if queue.isEmpty() {
		panic("trying to dequeue from empty queue")
	}
	cell := queue.bottomSentinel.prev
	cell.delete()
	return cell.data
}

// deque

// Add an item at the bottom of the deque.
func (deque *DoublyLinkedList) pushBottom(value string) {
	cell := &Cell{data: value}
	deque.bottomSentinel.addBefore(cell)
}

// Add an item at the top of the deque.
func (deque *DoublyLinkedList) pushTop(value string) {
	deque.push(value)
}

// Remove an item from the top of the deque.
func (deque *DoublyLinkedList) popTop() string {
	return deque.pop()
}

// Add an item at the top of the deque.
func (deque *DoublyLinkedList) popBottom() string {
	if deque.isEmpty() {
		panic("trying to pop from empty stack")
	}
	cell := deque.bottomSentinel.prev
	cell.delete()
	return cell.data
}

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

