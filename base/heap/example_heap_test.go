package heap

import (
	"fmt"
	"math/rand"
)

func Example_1() {
	myHeap := NewWithSlice([]Any{1, 9, 4, 3})
	myHeap.InitWithCmp(func(i, j int) bool {
		return myHeap.Get(i).(int) > myHeap.Get(j).(int)
	})
	myHeap.Push(5)
	myHeap.Push(8)
	fmt.Println(myHeap.Pop())
	fmt.Println(myHeap.Peek())
	for myHeap.Len() > 0 {
		fmt.Println(myHeap.Pop())
	}
	// output:
	// 9
	// 8
	// 8
	// 5
	// 4
	// 3
	// 1
}
func Example() {
	minHeap, maxHeap := NewWithSlice(nil), NewWithSlice(nil)
	minHeap.InitWithCmp(func(i, j int) bool {
		return minHeap.Get(i).(int) < minHeap.Get(j).(int)
	})
	maxHeap.InitWithCmp(func(i, j int) bool {
		return maxHeap.Get(i).(int) > maxHeap.Get(j).(int)
	})
	fmt.Println(minHeap.Len() == 0 && maxHeap.Len() == 0)

	min, max := 0, 100
	maxHeap.Push(max)
	minHeap.Push(min)
	for i := 0; i < max; i++ {
		minHeap.Push(rand.Intn(max))
		maxHeap.Push(rand.Intn(max))
	}
	fmt.Println(minHeap.Peek())
	fmt.Println(maxHeap.Peek())
	// output:
	// true
	// 0
	// 100
}
