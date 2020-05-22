package heap

type Cmp func(i, j int) bool
type Any interface{}

type Heap interface {
	InitWithCmp(cmp Cmp)
	Get(i int) Any
	Len() int
	Peek() Any
	Push(x Any)
	Pop() Any
	Fix(i int)
	Remove(i int) Any
}
