/*
	A set can store unique values, without any particular order.
*/
package set

import . "github.com/zrcoder/dsGo"

type Set map[Any]Empty

func New() Set {
	return make(map[Any]Empty)
}

func NewWithCapacity(c int) Set {
	return make(map[Any]Empty, c)
}

func (s Set) Add(item Any) {
	s[item] = Empty{}
}

func (s Set) Delete(item Any) {
	delete(s, item)
}

func (s Set) Has(item Any) bool {
	_, ok := s[item]
	return ok
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) AllItems() []Any {
	r := make([]Any, s.Size())
	i := 0
	for key := range s {
		r[i] = key
		i++
	}
	return r
}

func (s Set) Range(f func(item Any) bool) {
	for key := range s {
		if f(key) {
			return
		}
	}
}
