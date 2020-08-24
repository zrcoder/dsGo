/*
	A stack gives you a LIFO or last-in first-out order.
	You can only Push to add a new element to the top of the stack,
	Pop to remove the element from the top,
	and Peek at the top element without poping it off.
*/
package stack

import . "github.com/zrcoder/dsGo"

type stackItem struct {
	item Any
	next *stackItem
}

type Stack struct {
	peek  *stackItem
	depth uint64
}

func New() *Stack {
	return &Stack{}
}

// Add a new element to the top
func (s *Stack) Push(item Any) {
	s.peek = &stackItem{item: item, next: s.peek}
	s.depth++
}

// Remove the element from the top and returns it
func (s *Stack) Pop() Any {
	if s.depth == 0 {
		return nil
	}
	item := s.peek.item
	s.peek = s.peek.next
	s.depth--
	return item
}

// Returns the element from the top without deletion
func (s *Stack) Peek() Any {
	if s.depth == 0 {
		return nil
	}
	return s.peek.item
}
