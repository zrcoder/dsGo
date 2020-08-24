package set

import "testing"

func Test(t *testing.T) {
	set := New()
	for i := 0; i < 10; i++ {
		set.Add(i)
	}
	t.Log(set.AllItems())
	if set.Size() != 10 {
		t.Error("something wrong with func Size()")
	}
	set.Add(100)
	set.Add(100)
	set.Add(200)
	if set.Has(300) {
		t.Error("something wrong with func Has()")
	}
	if !set.Has(200) {
		t.Error("something wrong with func Has()")
	}
	set.Delete(100)
	set.Delete(555)
	if set.Size() != 11 {
		t.Error("something wrong with func Size()")
	}
	if len(set.AllItems()) != set.Size() {
		t.Error("len(set.AllItems()) != set.Size()")
	}
}

func Benchmark(b *testing.B) {
	set := NewWithCapacity(b.N)
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	for i := 0; i < b.N; i++ {
		set.Delete(i)
	}
}
