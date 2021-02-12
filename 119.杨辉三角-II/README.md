# 119.杨辉三角 II
[https://leetcode-cn.com/problems/pascals-triangle-ii/](https://leetcode-cn.com/problems/pascals-triangle-ii/) 
## 原题
给定一个非负索引 *k*，其中 *k* &le; 33，返回杨辉三角的第 *k*行。

<img alt="" src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif">

<small>在杨辉三角中，每个数是它左上方和右上方的数的和。</small>

 **示例:** 

```
输入: 3
输出: [1,3,3,1]

```
 **进阶：** 

你可以优化你的算法到 *O*(*k*) 空间复杂度吗？

 
**标签**
`数组` 


## 滚动数组-1
```go
func getRow(rowIndex int) []int {
	cur := []int{1}
	pre := cur
	for i := 0; i < rowIndex; i++ {
		cur = make([]int, len(pre)+1)
		cur[0], cur[len(cur)-1] = 1, 1
		for j := 0; j < len(pre)-1; j++ {
			cur[j+1] = pre[j] + pre[j+1]
		}
		pre = cur
	}
	return cur
}
```
>执行用时: 0 ms
内存消耗: 2 MB

一个储存之前的序列，一个储存当前的序列。

## 滚动数组-2
```go
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB

直接生成最后的长度的数组，然后从后往前生成一层的数据