## 重新设计 heap 包
标准库 container/heap 包的设计比较特别，Api 并非面向对象的风格，
这和其他语言甚至 Go 自身的其他数据结构如 container/list 不一样
### 当前 Api 初体验
简单起见，我们先假设要管理一些整形数字；要同时用到大顶堆和小顶堆。怎么写呢？

#### 步骤1
需要先实现两个类型，MinHeap和MaxHeap，分别实现 heap.Interface 接口， 即Len、Less、Swap、Push、Pop 五个方法：
```go
type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
```
这里已经有明显的代码坏味道，重复代码太多；可以通过该自定义 heap 增加函数类型属性 cmp 来解决，这里不细说，后边重新设计后有类似实现。

#### 步骤2
现在继续使用我们的大顶堆和小顶堆：
```go
	nums := []int{2, 9, 10, 7, 4, 3}
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	for _, v := range nums {
		heap.Push(minHeap, v)
		heap.Push(maxHeap, v)
	}
	fmt.Println(heap.Pop(minHeap))
	fmt.Println(maxHeap[0])
```
可以看到，没有Peek方法，直接取第0个元素即峰顶；push 和 pop 使用了 heap包的 Push Pop方法， 而不是直接这样：
```go
minHeap.Push(5)
maxHeap.Pop()
```
综合看，步骤 1 里需要实现 heap.Interface 接口，并且步骤2使用的是heap包的 Push 和 Pop 方法，而不是类型本身的 Push 和 Pop方法

### 分析修改 Api 设计
看起来标准库当前设计并不友好，尝试修改下，让 Api 更易用。

#### 1. 使用者只需关注比较逻辑 
堆底层是一个切片， heap.Interface 里要求的五个方法 Len、Less、Swap、Push、Pop，有四个无需使用者关注，只有比较逻辑需要使用者确定

这里可以定义一个只包含 Less方法(改名为Cmp更好)的接口让使用者实现，或者直接给我们的结构体增加函数类型的 cmp 属性
#### 2. Push 和 Pop 就可以直接按照堆实例的方法调用
基于上条分析，Push 和 Pop 就可以直接按照堆实例的方法调用，而不用弄成一个包方法

综上，我们需要提供 Heap，使用起来像这样:
```go
import (
	"fmt"

	"github.com/zrcoder/dsGo/base/heap"
)
func main()  {
	nums := []int{2, 9, 10, 7, 4, 3}

	cmp := func(a, b any) bool {
		return a.(int) > b.(int)
	}
	maxHeap := heap.New(cmp)
	for _, v := range nums {
		minHeap.Push(v)
		maxHeap.Push(v)
	}

	minHeap := heap.NewWithCap(len(nums)) // use default comparator: a < b

	fmt.Println(minHeap.Pop())
	fmt.Println(maxHeap.Peek())
}
```
可以看到，使用者唯一需要确定的就是比较逻辑，创建堆实例时传入比较函数即可。

### 实现新设计
有两个实现方法
#### 1. 包装标准库已有 Api
略
#### 2. 参考标准库核心方法 up 和 down 从头写

详见[具体实现](heap.go)

### 扩展

我们还实现了扩展API，Remove 和 Update，Remove 可以删除任意元素（Pop只能删除堆顶元素），Update 可以在某个元素值发改变后调整堆。

我们借助哈希表 idx 维护了每个元素在堆里的索引，知道索引后可以调用 up 和 down 方法在对数级复杂度内完成操作。

> 同时考虑了相同元素多次入堆的情况，用了哈希表 cnt 维护了每个元素的个数，data 数组中仅维护去重后的元素。