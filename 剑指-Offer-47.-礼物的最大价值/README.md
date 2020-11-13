# 剑指 Offer 47. 礼物的最大价值
[https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/](https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/)

## 贪心动规最大值
```go
func maxValue(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	r := make([][]int, row)
	for i := 0; i < row; i++ {
		r[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r[i][j] = grid[i][j]
			up := 0
			left := 0
			if i != 0 {
				up = r[i-1][j]
			}
			if j != 0 {
				left = r[i][j-1]
			}
			r[i][j] += max(up, left)
		}
	}
	return r[row-1][col-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```

>执行用时：8 ms, 在所有 Go 提交中击败了87.13%的用户
内存消耗：4.3 MB, 在所有 Go 提交中击败了32.00%的用户

就是简单的计算左边和上边的最大值作为到这里的最优解就行。唯一要注意的就是边界问题。