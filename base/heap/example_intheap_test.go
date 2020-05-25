package heap

import (
	"fmt"
)

func Example_intHeap() {
	h := NewWithSlice([]Any{2, 1, 5})
	h.InitWithCmp(func(i, j int) bool {
		return h.Get(i).(int) < h.Get(j).(int)
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
