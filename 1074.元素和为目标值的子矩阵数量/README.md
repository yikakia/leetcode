# 1074.元素和为目标值的子矩阵数量
[https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target](https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target) 
## 原题
给出矩阵  `matrix`  和目标值  `target` ，返回元素总和等于目标值的非空子矩阵的数量。

子矩阵  `x1, y1, x2, y2`  是满足 `x1 <= x <= x2`  且  `y1 <= y <= y2`  的所有单元  `matrix[x][y]`  的集合。

如果  `(x1, y1, x2, y2)` 和  `(x1', y1', x2', y2')`  两个子矩阵中部分坐标不同（如： `x1 != x1'` ），那么这两个子矩阵也不同。

 

 **示例 1：** 

<img alt="" src="https://assets.leetcode.com/uploads/2020/09/02/mate1.jpg" style="width: 242px; height: 242px;" />

```

输入：matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
输出：4
解释：四个只含 0 的 1x1 子矩阵。

```
 **示例 2：** 

```

输入：matrix = [[1,-1],[-1,1]], target = 0
输出：5
解释：两个 1x2 子矩阵，加上两个 2x1 子矩阵，再加上一个 2x2 子矩阵。

```
 **示例 3：** 

```

输入：matrix = [[904]], target = 0
输出：0

```
 
**提示：**
-  `1 <= matrix.length <= 100` 
-  `1 <= matrix[0].length <= 100` 
-  `-1000 <= matrix[i] <= 1000` 
-  `-10^8 <= target <= 10^8` 
 
**标签**
`数组` `动态规划` `Sliding Window` 


## 二维前缀和
```go
func numSubmatrixSumTarget(matrix [][]int, target int) (ret int) {
	n, m := len(matrix), len(matrix[0])
	preSum := make([][]int, n+1)
	for i := range preSum {
		preSum[i] = make([]int, m+1)
		if i != 0 {
			for j := range preSum[i] {
				if j == 0 {
					continue
				}
				preSum[i][j] = matrix[i-1][j-1] +
					preSum[i][j-1] +
					preSum[i-1][j] - preSum[i-1][j-1]
			}
		}
	}

	for i := range preSum {
		for j := range preSum[i] {
			for up := i - 1; up >= 0; up-- {
				for left := j - 1; left >= 0; left-- {
					area := preSum[i][j]
					area += preSum[up][left]
					area -= preSum[i][left]
					area -= preSum[up][j]
					if area == target {
						ret++
					}
				}
			}
		}
	}
	return
}
```
>执行用时: 144 ms
内存消耗: 4.6 MB

就是个标准的前缀和。只不过是二维的，所以写起来要注意怎么生成二维的前缀和。