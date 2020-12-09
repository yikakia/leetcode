# 62. 不同路径
[https://leetcode-cn.com/problems/unique-paths/](https://leetcode-cn.com/problems/unique-paths/) 
## 原题
一个机器人位于一个 `m x n`<em> </em>网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
 
**示例 1：**
<img src="https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png" />
```
输入：m = 3, n = 7
输出：28
```

**示例 2：**
```
输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
```
**示例 3：**
```
输入：m = 7, n = 3
输出：28
```
**示例 4：**
```
输入：m = 3, n = 3
输出：6
```
 
**提示：**
- `1 <= m, n <= 100`
- 题目数据保证答案小于等于 `2 * 10^9`

## 逐步累加
```go
func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		matrix[i][0] = 1
	}
	for i := 0; i < n; i++ {
		matrix[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}
	return matrix[m-1][n-1]
}
```
>执行用时: 0 ms
内存消耗: 2 MB

不过还可以优化，就是只用 `make([]int ,n)` 。因为只有当前行会被用到。

### 优化后的代码
```go
func uniquePaths(m int, n int) int {
	row := make([]int, n)
	row[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			row[j] += row[j-1]
		}
	}
	return row[n-1]
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB

简单地说就是 `row` 记录的是这个位置上在当前循环有多少种方法到这里。即当 `i == 3`，`j == 2`时，表示的是第四行的第二列有多少种方法到达。而它需要更新的话就累加上左边的元素就好了。

因为实际上如果换种方法看第一种方法的话，它可以看作复制上一行的当前列，并且加上当前行的上一列。即 
```go
matrix[i][j] = matrix[i-1][j] 
matirx[i][j] += matrix[i][j-1]
```

这个时候直接只管当前列就可以了，不断累加左边的就行。