package unionfind

import (
	base "github.com/zrcoder/dsGo/base/unionfind"
	"sync"
)

type UnionFind struct {
	inner base.UnionFind
	lock  sync.RWMutex
}

func NewUnionFind(n int) *UnionFind {
	return &UnionFind{inner: base.NewUnionFind(n)}
}

func (uf *UnionFind) Find(x int) int {
	uf.lock.RLock()
	root := uf.inner.Find(x)
	uf.lock.RUnlock()
	return root
}

func (uf *UnionFind) Join(x, y int) {
	uf.lock.Lock()
	uf.inner.Join(x, y)
	uf.lock.Unlock()
}
