/*
	A queue gives you a FIFO or first-in firs-out order.
	The element you inserted first is also the first one to come out again.
*/
package queue

import . "github.com/zrcoder/dsGo"

type queueItem struct {
	item Any
	next *queueItem
}

type Queen struct {
	head  *queueItem
	tail  *queueItem
	depth uint64
}

func New() *Queen {
	return &Queen{}
}

// Put a given item into the queue
func (q *Queen) Enqueue(item Any) {
	newItem := &queueItem{item: item}
	if q.depth == 0 {
		q.head = newItem
	} else {
		q.tail.next = newItem
	}
	q.tail = newItem
	q.depth++
}

// Remove the first item from the queue and returns it
func (q *Queen) Dequeue() Any {
	if q.depth == 0 {
		return nil
	}
	item := q.head.item
	q.head = q.head.next
	q.depth--
	return item
}
