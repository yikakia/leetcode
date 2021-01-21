# 1489.找到最小生成树里的关键边和伪关键边
[https://leetcode-cn.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/](https://leetcode-cn.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/) 
## 原题
给你一个 `n` 个点的带权无向连通图，节点编号为 `0` 到 `n-1` ，同时还有一个数组 `edges` ，其中 `edges[i] = [from``<sub>i</sub>, to<sub>i</sub>, weight<sub>i</sub>]` 表示在 `from<sub>i</sub>` 和 `to<sub>i</sub>` 节点之间有一条带权无向边。最小生成树 (MST) 是给定图中边的一个子集，它连接了所有节点且没有环，而且这些边的权值和最小。

请你找到给定图中最小生成树的所有关键边和伪关键边。如果从图中删去某条边，会导致最小生成树的权值和增加，那么我们就说它是一条关键边。伪关键边则是可能会出现在某些最小生成树中但不会出现在所有最小生成树中的边。

请注意，你可以分别以任意顺序返回关键边的下标和伪关键边的下标。

 

**示例 1：** 

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/ex1.png" style="height: 262px; width: 259px;">

```
输入：n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
输出：[[0,1],[2,3,4,5]]
解释：上图描述了给定图。
下图是所有的最小生成树。
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/msts.png" style="height: 553px; width: 540px;">
注意到第 0 条边和第 1 条边出现在了所有最小生成树中，所以它们是关键边，我们将这两个下标作为输出的第一个列表。
边 2，3，4 和 5 是所有 MST 的剩余边，所以它们是伪关键边。我们将它们作为输出的第二个列表。

```
**示例 2 ：** 

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/ex2.png" style="height: 253px; width: 247px;">

```
输入：n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
输出：[[],[0,1,2,3]]
解释：可以观察到 4 条边都有相同的权值，任选它们中的 3 条可以形成一棵 MST 。所以 4 条边都是伪关键边。

```
 

**提示：** 
- `2 <= n <= 100`
- `1 <= edges.length <= min(200, n * (n - 1) / 2)`
- `edges[i].length == 3`
- `0 <= from<sub>i</sub> < to<sub>i</sub> < n`
- `1 <= weight<sub>i</sub> <= 1000`
- 所有 `(from<sub>i</sub>, to<sub>i</sub>)` 数对都是互不相同的。
 
**标签**
`深度优先搜索` `并查集` 


## 最小生成树-Kruskal 算法

### 自定义的 Kruskal 算法
```go
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	newEdges := make([][]int, len(edges))
	for i := range edges {
		newEdges[i] = make([]int, 4)
		copy(newEdges[i], edges[i])
		newEdges[i][3] = i
	}
	sort.Slice(newEdges, func(i, j int) bool {
		return newEdges[i][2] < newEdges[j][2]
	})

	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) bool {
		fx, fy := find(x), find(y)
		if fx == fy {
			return false
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx

		return true
	}

	minSum := 0
	for _, e := range newEdges {
		x, y := e[0], e[1]
		if merge(x, y) {
			minSum += e[2]
		}
	}

	critical := []int{}
	notcritical := []int{}
	for i := range newEdges {
		tm := 0
		m := 0
		reset(fa, rank)
		for j, e := range newEdges {
			if j == i {
				continue
			}
			x, y := e[0], e[1]
			if merge(x, y) {
				tm += e[2]
				m++
			}
		}
		if tm > minSum || m+1 != n {
			critical = append(critical, newEdges[i][3])
		}
	}
	sort.Sort(sort.IntSlice(critical))
	for _, e := range newEdges {
		index := sort.Search(len(critical), func(x int) bool {
			return critical[x] >= e[3]
		})
		if index < len(critical) && critical[index] == e[3] {
			continue
		}

		tm := 0
		reset(fa, rank)
		merge(e[0], e[1])
		tm += e[2]
		for _, e := range newEdges {
			x, y := e[0], e[1]
			if merge(x, y) {
				tm += e[2]
			}
		}
		if tm == minSum {
			notcritical = append(notcritical, e[3])
		}
	}

	return [][]int{critical, notcritical}
}
func reset(fa, rank []int) {
	if len(fa) != len(rank) {
		return
	}
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
}
```
>执行用时: 44 ms
内存消耗: 4.3 MB

简单来说就是最小生成树的变种，关键是怎么找关键边和伪关键边。


- 找关键边：
    当不通过这条边生成最小生成树的时候，如果得到的权值总和比最小生成树的权值总和**大**，则这条边是最小生成树里面不可替换的，这条边就是关键边。
- 找伪关键边
    **当一条边不是关键边的时候**，如果通过这条边生成的最小生成树的权值总和和最小生成树的权值总和**相等**，那么这条边就是可以在最小生成树里面被替换的，这条边就是伪关键边。

基本就是对每一条边进行一次微调过的 Kruskal 算法。先进行关键边的查找，再进行伪关键边的查找。两者可以在同一个循环中处理（像官方题解那样写一个带 ignoreID 的参数的函数就可以了），也可以分开处理。我写的时候没想到可以通过把 ignoreID 置为负数就能实现不忽略了，于是就分开写了。

对于看这个边是不是关键边，可以先排好序后用二分查找来，不过如果像官方题解一样的话就不用了，先查找好，如果是关键边就 `continue`，不是的话就判断是不是伪关键边就可以了。

不过我这里写的时候还是有些臃肿，因为其实可以直接在 `edges` 数组上进行修改的，这样就不用再开一个数组了。

### 小修改的版本

```go
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	for i := range edges {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) bool {
		fx, fy := find(x), find(y)
		if fx == fy {
			return false
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx

		return true
	}

	minSum := 0
	for _, e := range edges {
		x, y := e[0], e[1]
		if merge(x, y) {
			minSum += e[2]
		}
	}

	critical := []int{}
	notcritical := []int{}
	for i := range edges {
		tm := 0
		m := 0
		reset(fa, rank)
		for j, e := range edges {
			if j == i {
				continue
			}
			x, y := e[0], e[1]
			if merge(x, y) {
				tm += e[2]
				m++
			}
		}
		if tm > minSum || m+1 != n {
			critical = append(critical, edges[i][3])
		}
	}

	sort.Slice(critical, func(i, j int) bool {
		return critical[i] < critical[j]
	})

	for _, e := range edges {
		index := sort.Search(len(critical), func(x int) bool {
			return critical[x] >= e[3]
		})
		if index < len(critical) && critical[index] == e[3] {
			continue
		}

		tm := 0
		reset(fa, rank)
		merge(e[0], e[1])
		tm += e[2]
		for _, e := range edges {
			x, y := e[0], e[1]
			if merge(x, y) {
				tm += e[2]
			}
		}
		if tm == minSum {
			notcritical = append(notcritical, e[3])
		}
	}

	return [][]int{critical, notcritical}
}
func reset(fa, rank []int) {
	if len(fa) != len(rank) {
		return
	}
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
}
```
>执行用时: 36 ms
内存消耗: 3.9 MB

可以看到时间和空间都优化了 **10%** 这还是挺不错的。