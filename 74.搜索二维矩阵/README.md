# 74.搜索二维矩阵
[https://leetcode-cn.com/problems/search-a-2d-matrix](https://leetcode-cn.com/problems/search-a-2d-matrix) 
## 原题
编写一个高效的算法来判断  `m x n`  矩阵中，是否存在一个目标值。该矩阵具有如下特性：
- 每行中的整数从左到右按升序排列。
- 每行的第一个整数大于前一行的最后一个整数。
 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/05/mat.jpg" style="width: 322px; height: 242px;" />
```

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/25/mat2.jpg" style="width: 322px; height: 242px;" />
```

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false

```
 

 **提示：** 
-  `m == matrix.length` 
-  `n == matrix[i].length` 
-  `1 <= m, n <= 100` 
-  `-10^4 <= matrix[i][j], target <= 10^4` 
 
**标签**
`数组` `二分查找` 


## 二分查找
```go
func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	start, end := 0, n*m-1
	mid := (start + end) / 2

	for start <= end {
		x, y := mid/m, mid%m
		// if x >= n || y >= m {
		// 	return false
		// }
		if matrix[x][y] < target {
			start = mid + 1
		} else if matrix[x][y] > target {
			end = mid - 1
		} else {
			return true
		}
		mid = (start + end) / 2
	}

	return false
}
```
>执行用时: 4 ms
内存消耗: 2.7 MB

就是把矩阵看成一维的，然后用一维的下标映射到二维的下标上。