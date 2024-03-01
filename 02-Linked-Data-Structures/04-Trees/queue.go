package main

type Cell [T any] struct {
	data T
	next *Cell [T]
	prev *Cell [T]
}

type Queue [T any] struct {
	topSentinel *Cell [T]
	bottomSentinel *Cell [T]
}

func (queue *Queue [T]) isEmpty() bool {
	return queue.topSentinel.next == queue.bottomSentinel
}

func (me *Cell [T]) addAfter(after *Cell [T]) {
	after.next = me.next
	after.prev = me
	me.next.prev = after
	me.next = after
}

func (me *Cell [T]) addBefore(before *Cell [T]) {
	prev := me.prev
	prev.addAfter(before)
}

func (me *Cell [T]) delete() {
	next := me.next
	prev := me.prev
	prev.next = next
	next.prev = prev
}

func (queue *Queue [T]) push(node T) {
	cell := &Cell [T]{data: node}
	queue.topSentinel.addAfter(cell)
}

func (queue *Queue [T]) pop() T {
	if queue.isEmpty() {
		panic("trying to pop from empty stack")
	}
	cell := queue.topSentinel.next
	cell.delete()
	return cell.data
}

func MakeQueue [T any]() Queue [T] {
	topSentinel := &Cell [T]{}
	bottomSentinel := &Cell [T]{}
	topSentinel.prev = nil
	topSentinel.next = bottomSentinel
	bottomSentinel.prev = topSentinel
	bottomSentinel.next = nil
	return Queue [T] {
		topSentinel: topSentinel,
		bottomSentinel: bottomSentinel,
	}
}

// Add an item to the top of the queue.
func (queue *Queue [T]) enqueue(value T) {
	queue.push(value)
}

// Remove an item from the bottom of the queue.
func (queue *Queue [T]) dequeue() T {
	if queue.isEmpty() {
		panic("trying to dequeue from empty queue")
	}
	cell := queue.bottomSentinel.prev
	cell.delete()
	return cell.data
}

