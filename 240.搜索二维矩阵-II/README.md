# 240.搜索二维矩阵 II
[https://leetcode-cn.com/problems/search-a-2d-matrix-ii](https://leetcode-cn.com/problems/search-a-2d-matrix-ii) 
## 原题
编写一个高效的算法来搜索 ` m x n ` 矩阵 `matrix` 中的一个目标值 `target` 。该矩阵具有以下特性：
- 每行的元素从左到右升序排列。
- 每列的元素从上到下升序排列。
 

<b>示例 1：</b>
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/25/searchgrid2.jpg" />
```

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

```
<b>示例 2：</b>
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/25/searchgrid.jpg" />
```

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false

```


**提示：** 
-  `m == matrix.length` 
-  `n == matrix[i].length` 
-  `1 <= n, m <= 300` 
-  `-10^9 <= matrix[i][j] <= 10^9` 
- 每行的所有元素从左到右升序排列
- 每列的所有元素从上到下升序排列
-  `-10^9 <= target <= 10^9` 
 
**标签**
`数组` `二分查找` `分治` `矩阵` 


## 每行进行二分
```go
func searchMatrix(matrix [][]int, target int) bool {
	for _, col := range matrix {
		if col[0] > target || col[len(col)-1] < target {
			continue
		}

		index := sort.SearchInts(col, target)

		if index < len(col) && col[index] == target {
			return true
		}
	}
	return false
}
```
>执行用时: 24 ms	
内存消耗: 6.7 MB

## 每次去掉一行或者一列
```go
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := n-1, 0

	for x >= 0 && y < m {
		if matrix[y][x] == target {
			return true
		}
		if matrix[y][x] > target {
			x--
		}
		if matrix[y][x] < target {
			y++
		}
	}
	return false
}
```
>执行用时: 24 ms	
内存消耗: 6.7 MB

思路就是右上角的元素是这一列的最小的，也是这一行的最大的元素，这样就能每次排除掉一列或者一行