package heap

import "fmt"

func Example_intHeap() {
	h := NewWithCap(4)
	h.Push(2)
	h.Push(1)
	h.Push(5)
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

type Item struct {
	Name     string
	Priority int
}

func Example_priorityQueue() {
	pq := New(func(a, b any) bool {
		return a.(*Item).Priority > b.(*Item).Priority
	})

	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	for name, priority := range items {
		pq.Push(&Item{
			Name:     name,
			Priority: priority,
		})
	}
	item := &Item{
		Name:     "orange",
		Priority: 1,
	}
	pq.Push(item)

	item.Priority = 5
	pq.Fix(pq.IndexOf(item))

	for pq.Len() > 0 {
		item := pq.Pop().(*Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Name)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}
