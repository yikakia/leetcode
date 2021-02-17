# 566.重塑矩阵
[https://leetcode-cn.com/problems/reshape-the-matrix/](https://leetcode-cn.com/problems/reshape-the-matrix/) 
## 原题
在MATLAB中，有一个非常有用的函数 `reshape` ，它可以将一个矩阵重塑为另一个大小不同的新矩阵，但保留其原始数据。

给出一个由二维数组表示的矩阵，以及两个正整数 `r` 和 `c` ，分别表示想要的重构的矩阵的行数和列数。

重构后的矩阵需要将原始矩阵的所有元素以相同的 **行遍历顺序** 填充。

如果具有给定参数的 `reshape` 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。

 **示例 1:** 

```

输入: 
nums = 
[[1,2],
 [3,4]]
r = 1, c = 4
输出: 
[[1,2,3,4]]
解释:
行遍历nums的结果是 [1,2,3,4]。新的矩阵是 1 * 4 矩阵, 用之前的元素值一行一行填充新矩阵。

```
 **示例 2:** 

```

输入: 
nums = 
[[1,2],
 [3,4]]
r = 2, c = 4
输出: 
[[1,2],
 [3,4]]
解释:
没有办法将 2 * 2 矩阵转化为 2 * 4 矩阵。 所以输出原矩阵。

```
 **注意：** 
- 给定矩阵的宽和高范围在 [1, 100]。
- 给定的 r 和 c 都是正数。
 
**标签**
`数组` 


## 
```go
func matrixReshape(nums [][]int, r int, c int) [][]int {
	rr := len(nums)
	cc := len(nums[0])
	if rr*cc != r*c {
		return nums
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	index := 0
	for i := 0; i < rr; i++ {
		for j := 0; j < cc; j++ {
			newI := index / c
			newJ := index % c
			index++
			res[newI][newJ] = nums[i][j]
		}
	}
	return res
}
```
>执行用时: 12 ms
内存消耗: 6.9 MB

把长度求出来，如果不能转换就返回原数组，不然地话就生成新的结果。