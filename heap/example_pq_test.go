package heap

import "fmt"

type Item struct {
	Name     string
	Priority int
}

func Example_priorityQueue() {
	pq := NewWithCap(4)
	pq.InitWithCmp(func(a, b Value) bool {
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

	// get the origin index before updating
	index := pq.IndexOf(item)
	item.Priority = 5
	pq.Fix(index)

	for pq.Len() > 0 {
		item := pq.Pop().(*Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Name)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}
