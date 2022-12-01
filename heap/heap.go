package heap

type Heap interface {
	Push(x any)
	Pop() any
	Peek() any

	Len() int
	IndexOf(x any) int
	Fix(index int)
	Remove(index int) any
	Update(index int, value any)
}

type Comparator func(a, b any) bool

const DefaultCapacity = 64

func New(cmp ...Comparator) Heap {
	return NewWithCap(DefaultCapacity, cmp...)
}

func NewWithCap(cap int, cmp ...Comparator) Heap {
	return &heapImp{
		slice: make([]any, 0, cap),
		cmp:   getComparator(cmp),
	}
}

func NewWithSlice(slice []any, cmp ...Comparator) Heap {
	res := &heapImp{
		slice: slice,
		cmp:   getComparator(cmp),
	}
	res.build()
	return res
}

func getComparator(cmp []Comparator) Comparator {
	if len(cmp) == 0 {
		return func(a, b any) bool {
			switch av := a.(type) {
			case int:
				return av < b.(int)
			case int8:
				return av < b.(int8)
			case int16:
				return av < b.(int16)
			case int32:
				return av < b.(int32)
			case int64:
				return av < b.(int64)
			case uint:
				return av < b.(uint)
			case uint8:
				return av < b.(uint8)
			case uint16:
				return av < b.(uint16)
			case uint32:
				return av < b.(uint32)
			case uint64:
				return av < b.(uint64)
			case uintptr:
				return av < b.(uintptr)
			case float32:
				return av < b.(float32)
			case float64:
				return av < b.(float64)
			default:
				panic("you should pass a comprator for the items")
			}
		}
	}
	return cmp[0]
}

type heapImp struct {
	cmp   Comparator
	slice []any
}

// build will build the heap by the given cmp.
// The complexity is O(n) where n = h.Len().
func (h *heapImp) build() {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *heapImp) Push(x any) {
	h.slice = append(h.slice, x)
	h.up(h.Len() - 1)
}

// Pop removes and returns the peek element from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *heapImp) Pop() any {
	last := h.Len() - 1
	h.swap(0, last)
	h.down(0, last)
	res := h.slice[last]
	h.slice = h.slice[:last]
	return res
}

// Peek returns the peek value of the heap
// The complexity is O(1)
func (h *heapImp) Peek() any {
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
func (h *heapImp) IndexOf(x any) int {
	res := -1
	for i, v := range h.slice {
		if v == x {
			return i
		}
	}
	return res
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
func (h *heapImp) Remove(i int) any {
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
func (h *heapImp) Update(i int, value any) {
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
