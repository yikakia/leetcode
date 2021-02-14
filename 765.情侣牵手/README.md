# 765.情侣牵手
[https://leetcode-cn.com/problems/couples-holding-hands/](https://leetcode-cn.com/problems/couples-holding-hands/) 
## 原题
N 对情侣坐在连续排列的 2N 个座位上，想要牵到对方的手。 计算最少交换座位的次数，以便每对情侣可以并肩坐在一起。 *一*次交换可选择任意两人，让他们站起来交换座位。

人和座位用 `0` 到 `2N-1` 的整数表示，情侣们按顺序编号，第一对是 `(0, 1)`，第二对是 `(2, 3)`，以此类推，最后一对是 `(2N-2, 2N-1)`。

这些情侣的初始座位 `row[i]` 是由最初始坐在第 i 个座位上的人决定的。

 **示例 1:** 

```

输入: row = [0, 2, 1, 3]
输出: 1
解释: 我们只需要交换row[1]和row[2]的位置即可。

```
 **示例 2:** 

```

输入: row = [3, 2, 0, 1]
输出: 0
解释: 无需交换座位，所有的情侣都已经可以手牵手了。

```
 **说明:** 
-  `len(row)` 是偶数且数值在 `[4, 60]`范围内。
- 可以保证 `row` 是序列 `0...len(row)-1` 的一个全排列。
 
**标签**
`贪心算法` `并查集` `图` 


## 并查集
```go
type unionFind struct {
	fa   []int
	rank []int
	size int
}

func newUnionFind(n int) *unionFind {
	tFa := make([]int, n)
	tRank := make([]int, n)
	for i := 0; i < n; i++ {
		tFa[i] = i
		tRank[i] = 1
	}
	return &unionFind{fa: tFa, rank: tRank, size: n}
}
func (uf *unionFind) Find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.Find(uf.fa[x])
	}
	return uf.fa[x]
}
func (uf *unionFind) Union(x, y int) {
	fx, fy := uf.Find(x), uf.Find(y)
	if fx == fy {
		return
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.fa[fy] = fx
	uf.rank[fx] += uf.rank[fy]
	uf.size--
}

func minSwapsCouples(row []int) int {
	n := len(row) / 2
	uf := newUnionFind(len(row))
	for i := 0; i < n; i++ {
		uf.Union(row[i*2], row[i*2+1])
		uf.Union(i*2, i*2+1)
	}

	return n - uf.size
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

并查集，统计共有多少对合在一起的，最后减去对数即可。
