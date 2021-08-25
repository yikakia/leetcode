# 797.所有可能的路径
[https://leetcode-cn.com/problems/all-paths-from-source-to-target](https://leetcode-cn.com/problems/all-paths-from-source-to-target) 
## 原题
给你一个有 `n` 个节点的 **有向无环图（DAG）** ，请你找出所有从节点 `0` 到节点 `n-1` 的路径并输出（ **不要求按特定顺序** ）

二维数组的第 `i` 个数组中的单元都表示有向图中 `i` 号节点所能到达的下一些节点，空就是没有下一个结点了。

译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。

 

 **示例 1：** 

<img alt="" src="https://assets.leetcode.com/uploads/2020/09/28/all_1.jpg" style="height: 242px; width: 242px;" />

```

输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3

```
 **示例 2：** 

<img alt="" src="https://assets.leetcode.com/uploads/2020/09/28/all_2.jpg" style="height: 301px; width: 423px;" />

```

输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

```
 **示例 3：** 

```

输入：graph = [[1],[]]
输出：[[0,1]]

```
 **示例 4：** 

```

输入：graph = [[1,2,3],[2],[3],[]]
输出：[[0,1,2,3],[0,2,3],[0,3]]

```
 **示例 5：** 

```

输入：graph = [[1,3],[2],[3],[]]
输出：[[0,1,2,3],[0,3]]

```
 

 **提示：** 
-  `n == graph.length` 
-  `2 <= n <= 15` 
-  `0 <= graph[i][j] < n` 
-  `graph[i][j] != i` （即，不存在自环）
-  `graph[i]` 中的所有元素 **互不相同** 
- 保证输入为 **有向无环图（DAG）** 
 
**标签**
`深度优先搜索` `广度优先搜索` `图` `回溯` 


## 回溯
```go
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph) - 1
	res := [][]int{}
	var dfs func(curNode int)
	visitList := make([]int, 0, len(graph))
	dfs = func(curNode int) {
		visitList = append(visitList, curNode)
		if curNode == n {
			tmp := make([]int, len(visitList))
			copy(tmp, visitList)
			res = append(res, tmp)
            return
		}
		for _, nxtNode := range graph[curNode] {
			dfs(nxtNode)
			visitList = visitList[:len(visitList)-1]
		}
	}
	dfs(0)
	return res
}
```
>执行用时：8 ms
内存消耗：6.4 MB

比较标准的回溯题。用深搜搜出结果然后再进行判断。因为没有环所以处理起来就比较简单。