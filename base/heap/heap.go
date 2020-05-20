package heap

import "container/heap"

type Cmp func(i, j int) bool
type Any interface{}

type Heap interface {
	InitWithCmp(cmp Cmp)
	Get(i int) Any
	Len() int
	Peek() Any
	Push(x Any)
	Pop() Any
	Fix(i int)
	Remove(i int) Any
}

func NewWithSlice(slice []Any) Heap {
	return &heapImp{inner: &helper{slice: slice}}
}

type heapImp struct {
	inner *helper
}

func (h *heapImp) InitWithCmp(cmp Cmp) {
	h.inner.cmp = cmp
	heap.Init(h.inner)
}

func (h *heapImp) Get(i int) Any {
	return h.inner.slice[i]
}

func (h *heapImp) Len() int {
	return h.inner.Len()
}

func (h *heapImp) Peek() Any {
	return h.inner.slice[0]
}

func (h *heapImp) Push(x Any) {
	heap.Push(h.inner, x)
}

func (h *heapImp) Pop() Any {
	return heap.Pop(h.inner)
}

func (h *heapImp) Fix(i int) {
	heap.Fix(h.inner, i)
}

func (h *heapImp) Remove(i int) Any {
	return heap.Remove(h.inner, i)
}
