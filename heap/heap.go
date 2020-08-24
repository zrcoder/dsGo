package heap

type Value interface{}
type Cmp func(a, b Value) bool

type Heap interface {
	InitWithCmp(cmp Cmp)

	Push(x Value)
	Pop() Value
	Peek() Value

	Len() int
	IndexOf(x Value) int
	Fix(i int)
	Remove(i int) Value
	Update(i int, value Value)
}

func New() Heap {
	return NewWithCap(0)
}

func NewWithCap(cap int) Heap {
	return &heapImp{slice: make([]Value, 0, cap)}
}

func NewWithSlice(slice []Value) Heap {
	return &heapImp{slice: slice}
}

type heapImp struct {
	cmp   Cmp
	slice []Value
}

// InitWithCmp will build the heap by the given cmp.
// The complexity is O(n) where n = h.Len().
func (h *heapImp) InitWithCmp(cmp Cmp) {
	h.cmp = cmp
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *heapImp) Push(x Value) {
	h.slice = append(h.slice, x)
	h.up(h.Len() - 1)
}

// Pop removes and returns the peek element from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *heapImp) Pop() Value {
	last := h.Len() - 1
	h.swap(0, last)
	h.down(0, last)
	res := h.slice[last]
	h.slice = h.slice[:last]
	return res
}

// Peek returns the peek value of the heap
// The complexity is O(1)
func (h *heapImp) Peek() Value {
	return h.slice[0]
}

// Len returns the size of the heap.
// The complexity is O(1)
func (h *heapImp) Len() int {
	return len(h.slice)
}

// IndexOf returns the index of x in the inner slice of the heap
// If x is not in the heap, returns -1
// The complexity is O(n)
func (h *heapImp) IndexOf(x Value) int {
	return h.indexOf(x, 0)
}

func (h *heapImp) indexOf(x Value, i int) int {
	if i >= h.Len() || h.cmp(x, h.slice[i]) {
		return -1
	}
	if h.slice[i] == x {
		return i
	}
	// search in left child
	if r := h.indexOf(x, 2*i+1); r != -1 {
		return r
	}
	// search in right child
	return h.indexOf(x, 2*i+2)
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h *heapImp) Fix(i int) {
	if !h.down(i, h.Len()) {
		h.up(i)
	}
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *heapImp) Remove(i int) Value {
	n := h.Len() - 1
	if n != i {
		h.swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	res := h.slice[n]
	h.slice = h.slice[:n]
	return res
}

// Update update the value of the element at index i, and then fix the heap
// The complexity is O(log n) where n = h.Len().
func (h *heapImp) Update(i int, value Value) {
	h.slice[i] = value
	h.Fix(i)
}

func (h *heapImp) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func (h *heapImp) up(i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || !h.cmp(h.slice[i], h.slice[parent]) {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

func (h *heapImp) down(i, n int) bool {
	cur := i
	for {
		child := 2*cur + 1 // left
		if child >= n || child < 0 {
			break
		}
		right := child + 1
		if right < n && h.cmp(h.slice[right], h.slice[child]) {
			child = right
		}
		if !h.cmp(h.slice[child], h.slice[cur]) {
			break
		}
		h.swap(cur, child)
		cur = child
	}
	return cur > i
}
