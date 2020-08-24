package queue

import "testing"

func Test(t *testing.T) {
	queue := New()
	total := 10
	for i := 0; i < total; i++ {
		queue.Enqueue(i)
	}
	for i := 0; i < total; i++ {
		item := queue.Dequeue()
		if i != item {
			t.Errorf("expected %d, but got %d", i, item)
		}
	}
}

func Benchmark(b *testing.B) {
	queue := New()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}
