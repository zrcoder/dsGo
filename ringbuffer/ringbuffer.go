package ringbuffer

import (
	. "github.com/zrcoder/dsGo"
)

type RingBuffer struct {
	slice      []Any
	readIndex  int
	writeIndex int
}

func NewWithSize(size int) *RingBuffer {
	return &RingBuffer{slice: make([]Any, size)}
}

func (rb *RingBuffer) Write(x Any) bool {
	if rb.spaceForWriting() == 0 {
		return false
	}
	rb.slice[rb.writeIndex%len(rb.slice)] = x
	rb.writeIndex++
	return true
}

func (rb *RingBuffer) Read() Any {
	if rb.spaceForReading() == 0 {
		return nil
	}
	res := rb.slice[rb.readIndex%len(rb.slice)]
	rb.readIndex++
	return res
}

func (rb *RingBuffer) spaceForReading() int {
	return rb.writeIndex - rb.readIndex

}

func (rb *RingBuffer) spaceForWriting() int {
	return len(rb.slice) - rb.spaceForReading()
}
