package main

import (
	"testing"
)

func Test_makeLinkedList(t *testing.T) {
	values := []string{"one", "two", "three"}
	list := makeLinkedList()
	list.addRange(values)
	if list.length() != 3 {
		t.Errorf("wrong list length: %d", list.length())
	}
}

func Test_contains(t *testing.T) {
	list := makeLinkedList()
	list.addRange([]string{"one", "two"})
	if list.contains("thousand") {
		t.Errorf("contains returns true but should be false")
	}
	if !list.contains("two") {
		t.Errorf("contains returns false but should be true")
	}
}

func Test_find(t *testing.T) {
	list := makeLinkedList()
	list.addRange([]string{"one", "two", "three"})
	cell := list.find("two")
	if cell == nil {
		t.Errorf("find returned nil cell")
	}
	if cell.data != "one" {
		t.Errorf("find returned wrong cell")
	}
	
	// find value that is not in the list.
	cell = list.find("four")
	if cell != nil {
		t.Errorf("find returned non-nil cell")
	}
}

func Test_push(t *testing.T) {
	list := makeLinkedList()
	list.addRange([]string{"one", "two", "three"})
	list.push("ZERO")
}

func Test_isEmpty(t *testing.T) {
	list := makeLinkedList()
	if !list.isEmpty() {
		t.Errorf("isEmpty should return true")
	}
	list.push("ZERO")
	if list.isEmpty() {
		t.Errorf("isEmpty should return false")
	}
}

func Test_pop(t *testing.T) {
	list := makeLinkedList()
	list.addRange([]string{"one", "two", "three"})
	list.push("ZERO")
	value := list.pop()
	if value != "ZERO" {
		t.Errorf("pop returned wrong value: %s", value)
	}
}

