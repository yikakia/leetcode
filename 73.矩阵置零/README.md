# 73.矩阵置零
[https://leetcode-cn.com/problems/set-matrix-zeroes](https://leetcode-cn.com/problems/set-matrix-zeroes) 
## 原题
给定一个 *m* x *n* 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用 **<a href="http://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95" target="_blank">原地</a>** 算法 **。** 

 **示例 1:** 

```
输入: 
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出: 
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]

```
 **示例 2:** 

```
输入: 
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出: 
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
```
 **进阶:** 
- 一个直接的解决方案是使用  O( *m* *n* ) 的额外空间，但这并不是一个好的解决方案。
- 一个简单的改进方案是使用 O( *m* + *n* ) 的额外空间，但这仍然不是最好的解决方案。
- 你能想出一个常数空间的解决方案吗？
 
**标签**
`数组` 


## 
```go
func setZeroes(matrix [][]int) {
	firstContain := false
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			firstContain = true
		}
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if firstContain {
			matrix[i][0] = 0
		}
	}
}
```
>执行用时: 12 ms
内存消耗: 6.2 MB

用某一行/列的开头为 0 表示这一行/列是否需要置零。但是为了避免第一列被更新，所以需要在更新这一行的开头的时候判断这个是否本来就是 0 。