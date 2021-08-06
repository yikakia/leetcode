# 847.访问所有节点的最短路径
[https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes](https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes) 
## 原题
存在一个由 `n` 个节点组成的无向连通图，图中的节点按从 `0` 到 `n - 1` 编号。

给你一个数组 `graph` 表示这个图。其中， `graph[i]` 是一个列表，由所有与节点 `i` 直接相连的节点组成。

返回能够访问所有节点的最短路径的长度。你可以在任一节点开始和停止，也可以多次重访节点，并且可以重用边。

 
 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/05/12/shortest1-graph.jpg" style="width: 222px; height: 183px;" />
```

输入：graph = [[1,2,3],[0],[0],[0]]
输出：4
解释：一种可能的路径为 [1,0,2,0,3]
```
 **示例 2：** 

<img alt="" src="https://assets.leetcode.com/uploads/2021/05/12/shortest2-graph.jpg" style="width: 382px; height: 222px;" />

```

输入：graph = [[1],[0,2,4],[1,3,4],[2],[1,2]]
输出：4
解释：一种可能的路径为 [0,1,4,2,3]

```
 

 **提示：** 
-  `n == graph.length` 
-  `1 <= n <= 12` 
-  `0 <= graph[i].length < n` 
-  `graph[i]` 不包含 `i` 
- 如果 `graph[a]` 包含 `b` ，那么 `graph[b]` 也包含 `a` 
- 输入的图总是连通图
 
**标签**
`位运算` `广度优先搜索` `图` `动态规划` `状态压缩` 


## 旅行商问题青春版
```go

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, tuple{i, 1 << i, 0})
	}
	for {
		t := q[0]
		q = q[1:]
		if t.mask == (1<<n)-1 {
			return t.dist
		}
		for _, v := range graph[t.u] {
			maskV := t.mask | 1<<v
			if !seen[v][maskV] {
				q = append(q, tuple{v, maskV, t.dist + 1})
				seen[v][maskV] = true
			}
		}
	}
}
```
>执行用时: 0 ms
内存消耗: 6.3 MB

这就是旅行商问题的青春版，简单地说就是暴搜。但是暴搜的时候需要进行记忆，不对已经搜索过的再进行一次搜索。这个时候就用状态压缩记忆当前这个查找序列找过了哪些节点，如果找到这个节点的路径之前出现过，那么就不再进行搜索，避免浪费。