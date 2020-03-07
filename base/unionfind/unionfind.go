package unionfind

/*
一个非常有趣的讲解： https://blog.csdn.net/niushuai666/article/details/6662911
*/

// 用数组实现，也可以定义节点建立森林
type UnionFind []int

func NewUnionFind(n int) UnionFind {
	unionFind := make([]int, n)
	for i := range unionFind {
		unionFind[i] = i
	}
	return unionFind
}

/* 递归实现find：
func (uf UnionFind) Find(x int) int {
	if uf[x] != x {
		uf[x] = uf.find(uf[x]) // 路径压缩
	}
	return uf[x]
}
*/
func (uf UnionFind) Find(x int) int {
	for uf[x] != x { // 路径压缩
		uf[x] = uf[uf[x]]
		x = uf[x]
	}
	return uf[x]
}
func (uf UnionFind) Join(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	uf[rootX] = rootY   // 可以按秩合并，即高度较小的树根插入高度较大的树根下面，进一步减少整个Union、Find操作的复杂度
}
