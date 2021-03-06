# 1319.连通网络的操作次数
[https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected/](https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected/) 
## 原题
用以太网线缆将 `n` 台计算机连接成一个网络，计算机的编号从 `0` 到 `n-1`。线缆用 `connections` 表示，其中 `connections[i] = [a, b]` 连接了计算机 `a` 和 `b`。

网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。

给你这个计算机网络的初始布线 `connections`，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。 

 

**示例 1：** 

**<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/11/sample_1_1677.png" style="height: 167px; width: 570px;">** 

```
输入：n = 4, connections = [[0,1],[0,2],[1,2]]
输出：1
解释：拔下计算机 1 和 2 之间的线缆，并将它插到计算机 1 和 3 上。

```
**示例 2：** 

**<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/11/sample_2_1677.png" style="height: 175px; width: 660px;">** 

```
输入：n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
输出：2

```
**示例 3：** 

```
输入：n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
输出：-1
解释：线缆数量不足。

```
**示例 4：** 

```
输入：n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]
输出：0

```
 

**提示：** 
- `1 <= n <= 10^5`
- `1 <= connections.length <= min(n*(n-1)/2, 10^5)`
- `connections[i].length == 2`
- `0 <= connections[i][0], connections[i][1] < n`
- `connections[i][0] != connections[i][1]`
- 没有重复的连接。
- 两台计算机不会通过多条线缆连接。
 
**标签**
`深度优先搜索` `广度优先搜索` `并查集` 


## 并查集
```go
type UnionFindSet struct {
	fa       []int
	rank     []int
	setCount int
}

func NewUnionFind(n int) *UnionFindSet {
	uf := UnionFindSet{fa: make([]int, n), rank: make([]int, n), setCount: n}
	for i := range uf.fa {
		uf.fa[i] = i
		uf.rank[i] = 1
	}
	return &uf
}

func (uf *UnionFindSet) reset() {
	for i := range uf.fa {
		uf.fa[i] = i
		uf.rank[i] = 1
	}
}

func (uf *UnionFindSet) Find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.Find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *UnionFindSet) Merge(x, y int) bool {
	fx, fy := uf.Find(x), uf.Find(y)
	if fx == fy {
		return false
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.fa[fy] = fx
	uf.setCount--
	return true
}
func makeConnected(n int, connections [][]int) int {
	if len(connections)+1 < n {
		return -1
	}

	uf := NewUnionFind(n)
	for _, e := range connections {
		uf.Merge(e[0], e[1])
	}

	return uf.setCount - 1
}
```
>执行用时: 76 ms
内存消耗: 11.3 MB

简单的并查集，这次试着写了一个结构体来执行。感觉还可以。
