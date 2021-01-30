# 778.水位上升的泳池中游泳
[https://leetcode-cn.com/problems/swim-in-rising-water/](https://leetcode-cn.com/problems/swim-in-rising-water/) 
## 原题
在一个 N x N 的坐标方格 `grid` 中，每一个方格的值 `grid[i][j]` 表示在位置 `(i,j)` 的平台高度。

现在开始下雨了。当时间为 `t` 时，此时雨水导致水池中任意位置的水位为 `t` 。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。

你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台 `(N-1, N-1)`？

 

**示例 1:** 

```

输入: [[0,2],[1,3]]
输出: 3
解释:
时间为0时，你位于坐标方格的位置为 (0, 0)。
此时你不能游向任意方向，因为四个相邻方向平台的高度都大于当前时间为 0 时的水位。

等时间到达 3 时，你才可以游向平台 (1, 1). 因为此时的水位是 3，坐标方格中的平台没有比水位 3 更高的，所以你可以游向坐标方格中的任意位置

```
**示例2:** 

```

输入: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
输出: 16
解释:
 0  1  2  3  4
24 23 22 21  5
12 13 14 15 16
11 17 18 19 20
10  9  8  7  6

最终的路线用加粗进行了标记。
我们必须等到时间为 16，此时才能保证平台 (0, 0) 和 (4, 4) 是连通的

```
 

**提示:** 
- `2 <= N <= 50`.
- `grid[i][j]` 是 `[0, ..., N*N - 1]` 的排列。
 
**标签**
`堆` `深度优先搜索` `并查集` `二分查找` 


## 
```go
func swimInWater(grid [][]int) int {
	n := len(grid)
	edges := make([]edge, 0, n*n)
	for i, row := range grid {
		for j := range row {
			baseIndex := i*n + j
			if i > 0 {
				edges = append(edges,
					edge{baseIndex - n, baseIndex, max(grid[i-1][j], grid[i][j])})
			}
			if j > 0 {
				edges = append(edges,
					edge{baseIndex - 1, baseIndex, max(grid[i][j-1], grid[i][j])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].value < edges[j].value })
	uf := newUnionFind(n * n)
	for _, e := range edges {
		uf.union(e.start, e.end)
		if uf.inSameSet(0, n*n-1) {
			return e.value
		}
	}
	return 0
}

type edge struct {
	start, end, value int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type unionFind struct {
	fa, rank []int
}

func newUnionFind(n int) *unionFind {
	fa, rank := make([]int, n), make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &unionFind{fa: fa, rank: rank}
}

func (uf *unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *unionFind) union(x, y int) {
	if uf.inSameSet(x, y) {
		return
	}
	fx, fy := uf.fa[x], uf.fa[y]
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.fa[fy] = fx
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}
```
>执行用时: 16 ms
内存消耗: 6.5 MB

用并查集的话和昨天没有什么大的区别。都是创好边之后判断连通性。
