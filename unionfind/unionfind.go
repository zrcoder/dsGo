package unionfind

// 用数组实现，也可以定义节点建立森林
type UnionFind []int

func NewUnionFind(n int) UnionFind {
	unionFind := make([]int, n)
	for i := range unionFind {
		unionFind[i] = i
	}
	return unionFind
}
func (uf UnionFind) Find(x int) int {
	for uf[x] != x {
		x, uf[x] = uf[x], uf[uf[x]]
	}
	return x
}
func (uf UnionFind) Join(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	uf[rootX] = rootY // 可以按秩合并，即高度较小的树根插入高度较大的树根下面，进一步减少整个Union、Find操作的复杂度
}
