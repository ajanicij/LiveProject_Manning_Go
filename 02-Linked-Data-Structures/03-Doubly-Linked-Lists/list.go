package main

import (
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

