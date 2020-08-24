package heap

import (
	"fmt"
)

func Example_intHeap() {
	h := NewWithSlice([]Value{2, 1, 5})
	h.InitWithCmp(func(a, b Value) bool {
		return a.(int) < b.(int)
	})
	h.Push(3)
	fmt.Printf("minimum: %d\n", h.Peek())

	h.Update(h.IndexOf(1), 8)
	h.Remove(h.IndexOf(3))

	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	// Output:
	// minimum: 1
	// 2 5 8
}
