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

type BitSet []byte

const (
	byteLen     = 8
	defaultSize = 1024
)

func New() BitSet {
	return NewWithSize(defaultSize)
}

func NewWithSize(size int) BitSet {
	realSize := 1 + (size-1)/byteLen
	return make([]byte, realSize)
}

// Set true at the index
func (bs BitSet) Set(index int) {
	i, mask := bs.caculateInnerIndexAndMask(index)
	bs[i] |= mask
}

// Set false at the index
func (bs BitSet) Unset(index int) {
	i, mask := bs.caculateInnerIndexAndMask(index)
	bs[i] &= ^mask
}

// Returns the bool value at the index
func (bs BitSet) Get(index int) bool {
	i, mask := bs.caculateInnerIndexAndMask(index)
	return bs[i]&mask != 0
}

func (bs BitSet) caculateInnerIndexAndMask(index int) (int, byte) {
	return index / byteLen, 1 << uint(index%byteLen)
}
