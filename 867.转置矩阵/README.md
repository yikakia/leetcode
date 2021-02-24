# 867.转置矩阵
[https://leetcode-cn.com/problems/transpose-matrix](https://leetcode-cn.com/problems/transpose-matrix) 
## 原题
给你一个二维整数数组 `matrix` ， 返回 `matrix` 的 **转置矩阵** 。

矩阵的 **转置** 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

<img alt="" src="https://assets.leetcode.com/uploads/2021/02/10/hint_transpose.png" style="width: 600px; height: 197px;" />

 

 **示例 1：** 

```

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]

```
 **示例 2：** 

```

输入：matrix = [[1,2,3],[4,5,6]]
输出：[[1,4],[2,5],[3,6]]

```
 

 **提示：** 
-  `m == matrix.length` 
-  `n == matrix[i].length` 
-  `1 <= m, n <= 1000` 
-  `1 <= m * n <= 10^5` 
-  `-10^9 <= matrix[i][j] <= 10^9` 
 
**标签**
`数组` 


## 简单的转置
```go
func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := range matrix {
		for j := range matrix[i] {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}
```
>执行用时: 12 ms
内存消耗: 6.1 MB

把横坐标放在纵坐标即可。