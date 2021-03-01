# 304.二维区域和检索 - 矩阵不可变
[https://leetcode-cn.com/problems/range-sum-query-2d-immutable](https://leetcode-cn.com/problems/range-sum-query-2d-immutable) 
## 原题
给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 ( *row* 1, *col* 1) ，右下角为 ( *row* 2, *col* 2)。

<img alt="Range Sum Query 2D" src="https://assets.leetcode-cn.com/aliyun-lc-upload/images/304.png" style="width: 130px;">

<small>上图子矩阵左上角 (row1, col1) = **(2, 1)** ，右下角(row2, col2) = **(4, 3)，** 该子矩形内元素的总和为 8。</small>

 **示例:** 

```
给定 matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12

```
 **说明:** 
- 你可以假设矩阵不可变。
- 会多次调用 *sumRegion* 方法 *。* 
- 你可以假设 *row* 1 ≤ *row* 2 且 *col* 1 ≤ *col* 2。
 
**标签**
`动态规划` 


## 前缀和
```go
type NumMatrix struct {
	pres [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	if n == 0 {
		return NumMatrix{}
	}
	m := len(matrix[0])
	pres := make([][]int, n)
	for i, row := range matrix {
		pres[i] = make([]int, m+1)
		for j := range row {
			pres[i][j+1] = pres[i][j] + row[j]
		}
	}
	return NumMatrix{pres: pres}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (sum int) {
	for i := row1; i <= row2; i++ {
		sum += this.pres[i][col2+1] - this.pres[i][col1]
	}
	return sum
}
```
>执行用时: 32 ms
内存消耗: 8.2 MB

前缀和的二维表达方式。