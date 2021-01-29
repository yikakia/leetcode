# 1631.最小体力消耗路径
[https://leetcode-cn.com/problems/path-with-minimum-effort/](https://leetcode-cn.com/problems/path-with-minimum-effort/) 
## 原题
你准备参加一场远足活动。给你一个二维 `rows x columns` 的地图 `heights` ，其中 `heights[row][col]` 表示格子 `(row, col)` 的高度。一开始你在最左上角的格子 `(0, 0)` ，且你希望去最右下角的格子 `(rows-1, columns-1)` （注意下标从 **0**  开始编号）。你每次可以往 **上** ，**下** ，**左** ，**右**  四个方向之一移动，你想要找到耗费 **体力**  最小的一条路径。

一条路径耗费的 **体力值**  是路径上相邻格子之间 **高度差绝对值**  的 **最大值**  决定的。

请你返回从左上角走到右下角的最小 **体力消耗值**  。

 

**示例 1：** 

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex1.png" style="width: 300px; height: 300px;" />

```

输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
输出：2
解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。

```
**示例 2：** 

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex2.png" style="width: 300px; height: 300px;" />

```

输入：heights = [[1,2,3],[3,8,4],[5,3,5]]
输出：1
解释：路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。

```
**示例 3：** 
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex3.png" style="width: 300px; height: 300px;" />
```

输入：heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
输出：0
解释：上图所示路径不需要消耗任何体力。

```
 

**提示：** 
- `rows == heights.length`
- `columns == heights[i].length`
- `1 <= rows, columns <= 100`
- `1 <= heights[i][j] <= 10^6`
 
**标签**
`深度优先搜索` `并查集` `图` `二分查找` 


## 
```go
func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	edges := make([]edge, 0, n*m)
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			// 和上方的差值
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h - heights[i-1][j])})
			}
			// 和左边的差值
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].value < edges[j].value
	})

	uf := newUnionFind(n * m)
	for _, e := range edges {
		uf.union(e.start, e.end)
		if uf.inSameSet(0, n*m-1) {
			return e.value
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type edge struct {
	start, end, value int
}

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
		count: n}
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
	return true
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

```
>执行用时: 144 ms
内存消耗: 7.7 MB

并查集的思路来做，简单来说就是看题目中的格子看成是点，那么边就是格子与上下左右的格子之间的。而边长就是它们的高度之差的绝对值。

于是先求出所有的边，然后把它们按升序排序，从小到大进行合并的操作，直到左上角的点和右下角的点在同一个集合时，当前的值就是要求的那个值，即最小的最大高度之差。