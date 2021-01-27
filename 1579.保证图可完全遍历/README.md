# 1579.保证图可完全遍历
[https://leetcode-cn.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/](https://leetcode-cn.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/) 
## 原题
Alice 和 Bob 共有一个无向图，其中包含 n 个节点和 3  种类型的边：
- 类型 1：只能由 Alice 遍历。
- 类型 2：只能由 Bob 遍历。
- 类型 3：Alice 和 Bob 都可以遍历。
给你一个数组 `edges` ，其中 `edges[i] = [type<sub>i</sub>, u<sub>i</sub>, v<sub>i</sub>]` 表示节点 `u<sub>i</sub>` 和 `v<sub>i</sub>` 之间存在类型为 `type<sub>i</sub>` 的双向边。请你在保证图仍能够被 Alice和 Bob 完全遍历的前提下，找出可以删除的最大边数。如果从任何节点开始，Alice 和 Bob 都可以到达所有其他节点，则认为图是可以完全遍历的。

返回可以删除的最大边数，如果 Alice 和 Bob 无法完全遍历图，则返回 -1 。

 

**示例 1：** 

**<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex1.png" style="height: 191px; width: 179px;">** 

```
输入：n = 4, edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]
输出：2
解释：如果删除 [1,1,2] 和 [1,1,3] 这两条边，Alice 和 Bob 仍然可以完全遍历这个图。再删除任何其他的边都无法保证图可以完全遍历。所以可以删除的最大边数是 2 。

```
**示例 2：** 

**<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex2.png" style="height: 190px; width: 178px;">** 

```
输入：n = 4, edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]
输出：0
解释：注意，删除任何一条边都会使 Alice 和 Bob 无法完全遍历这个图。

```
**示例 3：** 

**<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex3.png" style="height: 190px; width: 178px;">** 

```
输入：n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]
输出：-1
解释：在当前图中，Alice 无法从其他节点到达节点 4 。类似地，Bob 也不能达到节点 1 。因此，图无法完全遍历。
```
 

**提示：** 
- `1 <= n <= 10^5`
- `1 <= edges.length <= min(10^5, 3 * n * (n-1) / 2)`
- `edges[i].length == 3`
- `1 <= edges[i][0] <= 3`
- `1 <= edges[i][1] < edges[i][2] <= n`
- 所有元组 `(type<sub>i</sub>, u<sub>i</sub>, v<sub>i</sub>)` 互不相同
 
**标签**
`并查集` 


## 
```go
type unionFind struct {
	fa, rank []int
	count    int
}

func newUnionFind(n int) *unionFind {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &unionFind{
		fa:    fa,
		rank:  rank,
		count: n,
	}
}
func (uf *unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}
	return uf.fa[x]
}
func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.fa[fy] = fx
	uf.count--
	return true
}
func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}
func (uf *unionFind) reset() {
	uf.count = len(uf.fa)
	for i := range uf.fa {
		uf.fa[i] = i
		uf.rank[i] = 1
	}
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	res := len(edges)
	ufa, ufb := newUnionFind(n), newUnionFind(n)
	for _, e := range edges {
		if e[0] == 3 {
			if !ufa.inSameSet(e[1]-1, e[2]-1) {
				ufa.union(e[1]-1, e[2]-1)
				ufb.union(e[1]-1, e[2]-1)
				res--
			}
		}
	}
	for _, e := range edges {
		switch e[0] {
		case 1:
			if ufa.union(e[1]-1, e[2]-1) {
				res--
			}
		case 2:
			if ufb.union(e[1]-1, e[2]-1) {
				res--
			}
		}
	}
	if ufa.count > 1 || ufb.count > 1 {
		return -1
	}
	return res
}
```
>执行用时: 320 ms
内存消耗: 20.3 MB

并查集的变式，这次是先遍历公共边，再分别遍历两种单独的边。最后判断两种图里面的连通分量是不是都是 1 即可。

先假设所有的边都可以被移除，然后遇到了能够保留的边就把结果减一，最后返回这个结果即可。

遍历公共边的时候先看是不是能够合并，能够合并的话就合并(两个集合都合并)，并且修改把能移除的边数减一（因为这条边要用于合并）。

遍历单独的边的时候也是根据能不能合并来进行合并并更新可以移除的边数。

最后的时候判断一下两种图里面的联通分量有没有大于 1 的。如果有的话就说明图不能被完全遍历，返回 -1 ，没有的话就说明图可以被完全遍历，返回可以移除的边数。