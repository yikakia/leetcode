# 59.螺旋矩阵 II
[https://leetcode-cn.com/problems/spiral-matrix-ii](https://leetcode-cn.com/problems/spiral-matrix-ii) 
## 原题
给你一个正整数  `n` ，生成一个包含 `1` 到  `n^2`  所有元素，且元素按顺时针顺序螺旋排列的  `n x n` 正方形矩阵 `matrix` 。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/spiraln.jpg" style="width: 242px; height: 242px;" />
```

输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

```
 **示例 2：** 

```

输入：n = 1
输出：[[1]]

```
 

 **提示：** 
-  `1 <= n <= 20` 
 
**标签**
`数组` 


## 方向数组
```go
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	x, y := -1, 0
	dist := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	xcnt, ycnt := 0, 0

	dir := 0
	for cnt := 1; cnt <= n*n; cnt++ {
		x += dist[dir][0]
		y += dist[dir][1]
		dirChange := false
		if x == n-xcnt-1 && dir == 0 {
			xcnt++
			dirChange = true
		} else if y == n-ycnt-1 && dir == 1 {
			ycnt++
			dirChange = true
		} else if x == xcnt-1 && dir == 2 {
			dirChange = true
		} else if y == ycnt && dir == 3 {
			dirChange = true
		}
		if dirChange {
			dir++
			dir %= 4
		}
		res[y][x] = cnt
	}

	return res
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

和昨天的[54.螺旋矩阵](../54.螺旋矩阵/README.md)类似，就是改成了在二维数组中写数据。