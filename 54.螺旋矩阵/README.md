# 54.螺旋矩阵
[https://leetcode-cn.com/problems/spiral-matrix](https://leetcode-cn.com/problems/spiral-matrix) 
## 原题
给你一个 `m` 行 `n` 列的矩阵  `matrix` ，请按照 **顺时针螺旋顺序** ，返回矩阵中的所有元素。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg" style="width: 242px; height: 242px;" />
```

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg" style="width: 322px; height: 242px;" />
```

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

```
 

 **提示：** 
-  `m == matrix.length` 
-  `n == matrix[i].length` 
-  `1 <= m, n <= 10` 
-  `-100 <= matrix[i][j] <= 100` 
 
**标签**
`数组` 


## 方向数组
```go

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	cnt := m * n
	x, y := -1, 0
	dist := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	xcnt, ycnt := 0, 0
	dir := 0
	res := make([]int, 0, cnt)
	for cnt != 0 {
		x += dist[dir][0]
		y += dist[dir][1]
		dirChange := false
		if x == n-xcnt-1 && dir == 0 {
			xcnt++
			dirChange = true
		} else if y == m-ycnt-1 && dir == 1 {
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
		res = append(res, matrix[y][x])
		cnt--
	}

	return res
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB

经典的循环输出。这里用的是方向数组。同时对边界条件进行判断，需要注意的是每一次方向变化后，下一次这个方向进行变化的时候判断条件的值是不同的。因此需要另外处理。x 方向上第一次变化后下一次判断的时候条件需要减一。剩下的一次类推即可。
