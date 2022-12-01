package heap

import (
	"fmt"
)

func Example_ints() {
	h := NewWithCap[int](4)
	h.Push(2)
	h.Push(1)
	h.Push(5)
	h.Push(3)
	fmt.Printf("minimum: %d\n", h.Peek())

	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

type Item struct {
	Name     string
	Priority int
}

func Example_custom() {
	h := New[*Item](func(a, b any) bool {
		return a.(*Item).Priority > b.(*Item).Priority
	})

	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	for name, priority := range items {
		h.Push(&Item{
			Name:     name,
			Priority: priority,
		})
	}

	item := &Item{
		Name:     "orange",
		Priority: 1,
	}

	h.Push(item)

	item.Priority = 3
	h.Update(item)

	h.Remove(item)

	item.Priority = 5
	h.Push(item)

	for h.Len() > 0 {
		item := h.Pop()
		fmt.Printf("%.2d:%s ", item.Priority, item.Name)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}

func Example_build() {
	nums := []int{6, 8, 5, 9, 3}
	h := Build(nums)
	h.Push(1)
	h.Remove(3)
	h.Remove(100)
	for h.Len() > 0 {
		fmt.Print(h.Pop())
		fmt.Print(",")
	}
	// Output:
	// 1,5,6,8,9,
}
