# 547.省份数量
[https://leetcode-cn.com/problems/number-of-provinces/](https://leetcode-cn.com/problems/number-of-provinces/) 
## 原题

有 `n` 个城市，其中一些彼此相连，另一些没有相连。如果城市 `a` 与城市 `b` 直接相连，且城市 `b` 与城市 `c` 直接相连，那么城市 `a` 与城市 `c` 间接相连。
**省份**  是一组直接或间接相连的城市，组内不含其他没有相连的城市。
给你一个 `n x n` 的矩阵 `isConnected` ，其中 `isConnected[i][j] = 1` 表示第 `i` 个城市和第 `j` 个城市直接相连，而 `isConnected[i][j] = 0` 表示二者不直接相连。
返回矩阵中 **省份**  的数量。
 
**示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/12/24/graph1.jpg" style="width: 222px; height: 142px;" />
```
输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
输出：2
```
**示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/12/24/graph2.jpg" style="width: 222px; height: 142px;" />
```
输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
输出：3
```
 
**提示：** 
- `1 <= n <= 200`
- `n == isConnected.length`
- `n == isConnected[i].length`
- `isConnected[i][j]` 为 `1` 或 `0`
- `isConnected[i][i] == 1`
- `isConnected[i][j] == isConnected[j][i]`
 
**标签**
`深度优先搜索` `并查集` 


## 并查集
```go
func findCircleNum(isConnected [][]int) int {
	fa := make([]int, len(isConnected))
	for i := range fa {
		fa[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for i := range isConnected {
		for j := range isConnected[i] {
			if isConnected[i][j] == 1 {
				fFrom, fTo := find(i), find(j)
				fa[fFrom] = fTo
			}
		}
	}

	ans := make(map[int]bool)
	for _, x := range fa {
		ans[find(x)] = true
	}
	return len(ans)
}
```
>执行用时: 28 ms
内存消耗: 6.6 MB

并查集和昨天类似，就是先指向自己，然后指向别人。因此最后全部再进行一次 `find` 就可以完全更新。

还有一种就是深度搜索，暂时先留着等之后再来。