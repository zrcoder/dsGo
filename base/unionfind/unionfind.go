package unionfind

/*
一个非常有趣的讲解： https://blog.csdn.net/niushuai666/article/details/6662911
*/
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
		uf[x] = uf.find(uf[x])
	}
	return uf[x]
}
*/
func (uf UnionFind) Find(x int) int {
	root := x
	for root != uf[root] {
		root = uf[root]
	}
	for root != x {
		uf[x], x = root, uf[x]
	}
	return root
}
func (uf UnionFind) Join(x, y int) {
	uf[x] = y
}
