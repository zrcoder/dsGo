package heap

type Heap[T comparable] struct {
	cmp  Comparator
	data []T
	idx  map[T]int
	cnt  map[T]int
	size int
}

type Comparator func(a, b any) bool

const DefaultCapacity = 64

func New[T comparable](cmp ...Comparator) *Heap[T] {
	return NewWithCap[T](DefaultCapacity, cmp...)
}

func NewWithCap[T comparable](cap int, cmp ...Comparator) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0, cap),
		idx:  make(map[T]int),
		cnt:  make(map[T]int),
		cmp:  getComparator(cmp),
	}
}

func Build[T comparable, S ~[]T](data S, cmp ...Comparator) *Heap[T] {
	res := &Heap[T]{
		data: data,
		idx:  make(map[T]int),
		cnt:  make(map[T]int),
		cmp:  getComparator(cmp),
		size: len(data),
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

// build will build the heap by the given cmp.
// The complexity is O(n) where n is the size of the heap.
func (h *Heap[T]) build() {
	n := len(h.data)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	for i, v := range h.data {
		h.idx[v] = i
		h.cnt[v]++
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n is the size of the heap.
func (h *Heap[T]) Push(x T) {
	if h.cnt[x] == 0 {
		n := len(h.data)
		h.idx[x] = n
		h.data = append(h.data, x)
		h.up(n)
	}
	h.cnt[x]++
	h.size++
}

// Pop removes and returns the peek element from the heap.
// The complexity is O(log n) where n is the size of the heap.
// Pop is equivalent to Remove(h.data[0]).
func (h *Heap[T]) Pop() T {
	if h.size == 0 {
		var x T
		return x
	}
	res := h.data[0]
	if h.cnt[res] == 1 {
		last := len(h.data) - 1
		h.swap(0, last)
		h.down(0, last)
		var zero T
		h.data[last] = zero // // avoid memory leak if T is pointer
		h.data = h.data[:last]
		h.idx[res] = -1 // for safety
	}
	h.cnt[res]--
	h.size--
	return res
}

// Peek returns the peek value of the heap
// The complexity is O(1)
func (h *Heap[T]) Peek() T {
	return h.data[0]
}

// Len returns the size of the heap.
// The complexity is O(1)
func (h *Heap[T]) Len() int {
	return h.size
}

// Remove removes and returns any element from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Remove(x T) T {
	if h.cnt[x] == 0 {
		return x
	}
	if h.cnt[x] == 1 {
		i := h.idx[x]
		last := len(h.data) - 1
		if last != i {
			h.swap(i, last)
			if !h.down(i, last) {
				h.up(i)
			}
		}
		var zero T
		h.data[last] = zero // avoid memory leak if T is pointer
		h.data = h.data[:last]
		h.idx[x] = -1 // for safety
	}
	h.size--
	h.cnt[x]--
	return x
}

// Update re-establishes the heap ordering after the element has changed its value.
// Changing the value of the element x and then calling Fix is equivalent to,
// but less expensive than, calling Remove(x) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Update(x T) {
	if h.cnt[x] == 0 {
		return
	}
	i := h.idx[x]
	if !h.down(i, len(h.data)) {
		h.up(i)
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.idx[h.data[i]] = i
	h.idx[h.data[j]] = j
}

func (h *Heap[T]) up(i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || !h.cmp(h.data[i], h.data[parent]) {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

func (h *Heap[T]) down(i, n int) bool {
	cur := i
	for {
		child := 2*cur + 1 // left
		if child >= n || child < 0 {
			break
		}
		right := child + 1
		if right < n && h.cmp(h.data[right], h.data[child]) {
			child = right
		}
		if !h.cmp(h.data[child], h.data[cur]) {
			break
		}
		h.swap(cur, child)
		cur = child
	}
	return cur > i
}
