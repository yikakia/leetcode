# 剑指 Offer 13. 机器人的运动范围
[https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/](https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)

## `DFS`
```go
func countDig(n int) (r int) {
	for n != 0 {
		r += n % 10
		n /= 10
	}
	return
}
func movingCount(m int, n int, k int) int {
	visited := make([][]int, 0)
	for i := 0; i < m; i++ {
		visited = append(visited, make([]int, n))
	}
	return dfs(m, n, 0, 0, k, visited)
}
func dfs(m, n, x, y int, k int, visited [][]int) int {
	if x < 0 || y < 0 || x >= m || y >= n {
		return 0
	}
	if visited[x][y] != 0 {
		return 0
	}
	if countDig(x)+countDig(y) > k {
		return 0
	}
	visited[x][y] = 1
	steps := 1
	steps += dfs(m, n, x+1, y, k, visited)
    steps += dfs(m, n, x, y+1, k, visited)
    steps += dfs(m, n, x-1, y, k, visited)
	steps += dfs(m, n, x, y-1, k, visited) 
	return steps
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：3.3 MB, 在所有 Go 提交中击败了12.64%的用户


和之前的题一样，就是`DFS`的应用。不同的是这次需要记录哪些已经访问过了。重要的是`DFS`要返回的是什么要自己搞清楚。

因为在经过前面的检查后，这个点是一定能访问的，所以基础值为1，之后向四周遍历，累积周围的点的合法的砖块数。到最后返回的就是(0,0)到(m,n)的砖块数。因为不用回溯，所以向下和向左搜索可以删掉。删掉之后的内存消耗如下，少了一大截。
```go
func dfs(m, n, x, y int, k int, visited [][]int) int {
	if x < 0 || y < 0 || x >= m || y >= n {
		return 0
	}
	if visited[x][y] != 0 {
		return 0
	}
	if countDig(x)+countDig(y) > k {
		return 0
	}
	visited[x][y] = 1
	steps := 1
	steps += dfs(m, n, x+1, y, k, visited)
    steps += dfs(m, n, x, y+1, k, visited)
	return steps
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了65.59%的用户