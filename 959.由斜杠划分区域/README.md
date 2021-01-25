# 959.由斜杠划分区域
[https://leetcode-cn.com/problems/regions-cut-by-slashes/](https://leetcode-cn.com/problems/regions-cut-by-slashes/) 
## 原题
在由 1 x 1 方格组成的 N x N 网格 `grid` 中，每个 1 x 1 方块由 `/`、`\` 或空格构成。这些字符会将方块划分为一些共边的区域。

（请注意，反斜杠字符是转义的，因此 `\` 用 `"\\"` 表示。）。

返回区域的数目。

 
**示例 1：** 

```
输入：
[
  " /",
  "/ "
]
输出：2
解释：2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/1.png">
```
**示例 2：** 

```
输入：
[
  " /",
  "  "
]
输出：1
解释：2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/2.png">
```
**示例 3：** 

```
输入：
[
  "\\/",
  "/\\"
]
输出：4
解释：（回想一下，因为 \ 字符是转义的，所以 "\\/" 表示 \/，而 "/\\" 表示 /\。）
2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/3.png">
```
**示例 4：** 

```
输入：
[
  "/\\",
  "\\/"
]
输出：5
解释：（回想一下，因为 \ 字符是转义的，所以 "/\\" 表示 /\，而 "\\/" 表示 \/。）
2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/4.png">
```
**示例 5：** 

```
输入：
[
  "//",
  "/ "
]
输出：3
解释：2x2 网格如下：
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/5.png">

```
 

**提示：** 
- `1 <= grid.length == grid[0].length <= 30`
- `grid[i][j]` 是 `';/';`、`';\';`、或 `'; ';`。
 
**标签**
`深度优先搜索` `并查集` `图` 


## 
```go
type UnionFind struct {
	fa, rank []int
	count    int
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &UnionFind{fa: fa, rank: rank, count: n}
}
func (uf *UnionFind) Find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.Find(uf.fa[x])
	}
	return uf.fa[x]
}
func (uf *UnionFind) Union(x, y int) bool {
	fx, fy := uf.Find(x), uf.Find(y)
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
func regionsBySlashes(grid []string) int {
	n := len(grid)
	uf := NewUnionFind(4 * n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			index := i*n + j
			/*
				______
				|\ 0/|
				|3\/ |
				| /\1|
				|/2 \|
				------
			*/
			// 把当前的▲和下面的▼相连
			if i < n-1 {
				bottom := index + n
				uf.Union(index*4+2, bottom*4)
			}
			// 把当前的最右边的三角形和右边的最左边的三角形相连
			if j < n-1 {
				right := index + 1
				uf.Union(index*4+1, right*4+3)
			}
			if grid[i][j] == '/' {
				uf.Union(index*4, index*4+3)
				uf.Union(index*4+1, index*4+2)
			} else if grid[i][j] == '\\' {
				uf.Union(index*4, index*4+1)
				uf.Union(index*4+2, index*4+3)
			} else {
				uf.Union(index*4, index*4+1)
				uf.Union(index*4+1, index*4+2)
				uf.Union(index*4+2, index*4+3)
			}

		}
	}
	return uf.count
}
```
>执行用时: 4 ms
内存消耗: 4.6 MB

这是变种的并查集。把一个正方形切成四个三角形然后进行合并即可。

还有要注意的是左边的右侧的三角形和右边的左侧的三角形是联通的，所以是可以相连的。同理，上方的下侧的三角形和下方的上侧的三角形也是联通的。

最后返回连通块的个数即可