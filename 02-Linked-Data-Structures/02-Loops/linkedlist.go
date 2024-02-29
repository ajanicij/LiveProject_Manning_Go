package main

import (
	"strings"
)

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

func makeLinkedList() LinkedList {
	sentinel := &Cell{}
	return LinkedList{
		sentinel: sentinel,
	}
}

func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	me.next = after
}

func (me *Cell) deleteAfter() *Cell {
	next := me.next
	if next == nil {
		panic("Can't delete after end of list")
	}
	me.next = next.next
	return next
}

func (list *LinkedList) addRange(values []string) {
	var lastCell *Cell
	for lastCell = list.sentinel; lastCell.next != nil; lastCell = lastCell.next {}
	for _, value := range values {
		cell := &Cell{data: value}
		lastCell.addAfter(cell)
		lastCell = cell
	}
}

func (list *LinkedList) toString(separator string) string {
	builder := strings.Builder{}
	for cell := list.sentinel.next; cell != nil; cell = cell.next {
		builder.WriteString(cell.data)
		if cell.next != nil {
			builder.WriteString(separator)
		}
	}
	return builder.String()
}

func (list *LinkedList) length() int {
	l := 0
	cell := list.sentinel
	for {
		cell = cell.next
		if cell == nil {
			break
		}
		l++
	}
	return l
}

func (list *LinkedList) isEmpty() bool {
	return list.sentinel.next == nil
}

func (list *LinkedList) contains(value string) bool {
	cell := list.find(value)
	return cell != nil
}

func (list *LinkedList) find(value string) *Cell {
	cell := list.sentinel.next
	for cell != nil {
		if cell.data == value {
			return cell
		}
		cell = cell.next
	}
	return nil
}

func (list *LinkedList) push(value string) {
	cell := &Cell{data: value}
	list.sentinel.addAfter(cell)
}

func (list *LinkedList)  pop() string {
	if list.isEmpty() {
		panic("trying to pop from empty stack")
	}
	cell := list.sentinel.next
	list.sentinel.deleteAfter()
	return cell.data
}

func (me *Cell) hasLoop() bool {
	fast := me
	slow := me
	for {
		fast = fast.next
		if fast == nil {
			break
		}
		fast = fast.next
		if fast == nil {
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

func (list *LinkedList) hasLoop() bool {
	return list.sentinel.hasLoop()
}

func (list *LinkedList) toStringMax(separator string, max int) string {
	builder := strings.Builder{}
	count := 1
	for cell := list.sentinel.next; cell != nil; cell = cell.next {
		builder.WriteString(cell.data)
		if count >= max {
			break
		}
		if cell.next != nil {
			builder.WriteString(separator)
		}
		count++
	}
	return builder.String()
}

