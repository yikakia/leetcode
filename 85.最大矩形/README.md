# 85. 最大矩形
[https://leetcode-cn.com/problems/maximal-rectangle/](https://leetcode-cn.com/problems/maximal-rectangle/) 
## 原题
给定一个仅包含 `0` 和 `1` 、大小为 `rows x cols` 的二维二进制矩阵，找出只包含 `1` 的最大矩形，并返回其面积。
 
**示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/14/maximal.jpg" style="width: 402px; height: 322px;" />
```
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
```
**示例 2：** 
```
输入：matrix = []
输出：0
```
**示例 3：** 
```
输入：matrix = [["0"]]
输出：0
```
**示例 4：** 
```
输入：matrix = [["1"]]
输出：1
```
**示例 5：** 
```
输入：matrix = [["0","0"]]
输出：0
```
 
**提示：** 
- `rows == matrix.length`
- `cols == matrix[0].length`
- `0 <= row, cols <= 200`
- `matrix[i][j]` 为 `'0'` 或 `'1'`


## 
```go
func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i, row := range matrix {
		dp[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := dp[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, dp[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
>执行用时: 16 ms
内存消耗: 4 MB

先统计每一行以该点为最右点的形状为`1*N`的矩形的最大长度（宽为1）。

然后统计以每个点作为最右下角的点的矩形的面积。具体为向上查找，把当前的最小长度和当前点的最大长度进行比较，当最大长度更小时更新当前的最小长度。面积就是长乘宽。

本来是想到了这种做法的，已经想到了分别统计，结果我就只想在一次循环里面搞定，于是就把自己绕死了。代码的简洁和实现的快捷上，应该还是偏向实现的快捷。先实现才好修改。