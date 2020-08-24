# dsGo
Data structures impletioned with Go

### `base` or `safe`?
For each data structure, we give a base version, which will not be thread safe.

If you want a safe one, just make a simple wrapper of the base one, for example, let's have a look at the wrapper of set. 
```
import (
	"sync"
 
	ds "github.com/zrcoder/dsGo"
	base "github.com/zrcoder/dsGo/set"
)
 
type Set struct {
	sync.RWMutex
	inner base.Set
}
 
func New() *Set {
	return &Set{inner: base.New()}
}
 
func NewWithCapacity(c int) *Set {
	return &Set{inner: base.NewWithCapacity(c)}
}
 
func (s *Set) Add(item ds.Any) {
	s.Lock()
	s.inner.Add(item)
	s.Unlock()
}
 
func (s *Set) Delete(item ds.Any) {
	s.Lock()
	s.inner.Delete(item)
	s.Unlock()
}
 
func (s *Set) Has(item ds.Any) bool {
	s.RLock()
	exist := s.inner.Has(item)
	s.RUnlock()
	return exist
}
 
func (s *Set) Size() int {
	s.RLock()
	length := s.inner.Size()
	s.RUnlock()
	return length
}
 
func (s *Set) AllItems() []ds.Any {
	s.RLock()
	items := s.inner.AllItems()
	s.RUnlock()
	return items
}

func (s *Set) Range(f func(item ds.Any) bool)  {
	s.Lock()
	s.inner.Range(f)
	s.Unlock()
}

```
the base version usually performances better than the safe one <br>
and the safe version usually used for concurrent scenes
### our data structures
[queue](queue/queue.go)
```text
A queue gives you a FIFO or first-in firs-out order.
```
[stack](stack/stack.go)
```text
A stack gives you a LIFO or last-in first-out order.
```
[heap](heap/readme.md)
```text
A heap with APIs easy to use, different from container/heap in standard lib
```
[set](set/set.go)
```text
A set can store unique values, without any particular order.
```
[bit set](bitset/bitset.go)
```text
A bit set is a fixed-size sequence of n bits.
```
[unionfind](unionfind/readme.md)
```text
union find
```
[ring buffer](ringbuffer/ringbuffer.go)
```
Also known as a circular buffer.
```
### data structures in standard library
sync.Map
```
A thread safe hash map
```
container/list.List
```
A doubly linked list
```
container/heap
```
A heap is a tree with the property that each node is the minimum-valued node in its subtree.
```
container/ring.Ring
```
A ring is an element of a circular list, or ring.
```
### references
[swift-algorithm-club in GitHub](https://github.com/raywenderlich/swift-algorithm-club)
