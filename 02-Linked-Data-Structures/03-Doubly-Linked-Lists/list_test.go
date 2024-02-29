package main

import (
	"testing"
)

func Test_makeDoublyLinkedList(t *testing.T) {
	values := []string{"one", "two", "three"}
	list := makeDoublyLinkedList()
	list.addRange(values)
	if list.length() != 3 {
		t.Errorf("wrong list length: %d", list.length())
	}
	first := list.topSentinel.next
	if first.data != "one" {
		t.Errorf("wrong first element: %s", first.data)
	}
}

func Test_contains(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"one", "two"})
	if list.contains("thousand") {
		t.Errorf("contains returns true but should be false")
	}
	if !list.contains("two") {
		t.Errorf("contains returns false but should be true")
	}
}

func Test_find(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"one", "two", "three"})
	cell := list.find("two")
	if cell == nil {
		t.Errorf("find returned nil cell")
	}
	if cell.data != "two" {
		t.Errorf("find returned wrong cell")
	}
	
	// find value that is not in the list.
	cell = list.find("four")
	if cell != nil {
		t.Errorf("find returned non-nil cell")
	}
}

func Test_isEmpty(t *testing.T) {
	list := makeDoublyLinkedList()
	if !list.isEmpty() {
		t.Errorf("isEmpty should return true")
	}
	list.push("ZERO")
	if list.isEmpty() {
		t.Errorf("isEmpty should return false")
	}
}

func Test_pop(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"one", "two", "three"})
	list.push("ZERO")
	value := list.pop()
	if value != "ZERO" {
		t.Errorf("pop returned wrong value: %s", value)
	}
}

func Test_hasLoop(t *testing.T) {
	cell1 := &Cell{data: "one"}
	cell2 := &Cell{data: "two"}
	cell3 := &Cell{data: "three"}
	cell4 := &Cell{data: "four"}
	list := makeDoublyLinkedList()
	list.topSentinel.addAfter(cell1)
	cell1.addAfter(cell2)
	cell2.addAfter(cell3)
	cell3.addAfter(cell4)
	if list.hasLoop() {
		t.Errorf("hasLoop must return false")
	}
	cell4.next = cell1
	if !list.hasLoop() {
		t.Errorf("hasLoop must return true")
	}
}

// queue

func Test_enqueue(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"one", "two", "three"})
	list.enqueue("ZERO")
	value := list.pop()
	if value != "ZERO" {
		t.Errorf("pop returned wrong value: %s", value)
	}
}

func Test_dequeue(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"one", "two", "three"})
	list.enqueue("ZERO")
	value := list.dequeue()
	if value != "three" {
		t.Errorf("dequeue returned wrong value: %s", value)
	}
	value = list.dequeue()
	if value != "two" {
		t.Errorf("dequeue returned wrong value: %s", value)
	}
}

// deque

func Test_pushBottom(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"one"})
	list.pushBottom("ZERO")
	value := list.pop()
	if value != "one" {
		t.Errorf("pop returned wrong value: %s", value)
	}
	value = list.pop()
	if value != "ZERO" {
		t.Errorf("pop returned wrong value: %s", value)
	}
}

