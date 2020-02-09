/*
	BitSet is a fixed-size sequence of n bits.
	Also known as bit array or bit vector.

	To store whether something is true or false you use a Bool.
	But what if you need to remember whether 10,000 things are true or not?

	You could make an array of 10,000 booleans
	but you can also use 10,000 bits instead.
	That's a lot more compact because 10,000 bits fit in less than 1250 bytes(10,000 / 8).
*/
package bitset

import (
	"sync"

	base "github.com/zrcoder/dsGo/base/bitset"
)

type BitSet struct {
	lock  sync.RWMutex
	inner base.BitSet
}

func New() *BitSet {
	return &BitSet{inner: base.New()}
}

func NewWithSize(size int) *BitSet {
	return &BitSet{inner: base.NewWithSize(size)}
}

// Set true at the index
func (bs *BitSet) Set(index int) {
	bs.lock.Lock()
	bs.inner.Set(index)
	bs.lock.Unlock()
}

// Set false at the index
func (bs *BitSet) Unset(index int) {
	bs.lock.Lock()
	bs.inner.Unset(index)
	bs.lock.Unlock()
}

// Returns the bool value at the index
func (bs *BitSet) Get(index int) bool {
	bs.lock.RLock()
	r := bs.inner.Get(index)
	bs.lock.RUnlock()
	return r
}
