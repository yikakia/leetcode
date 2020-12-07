# 861. 翻转矩阵后的得分
[https://leetcode-cn.com/problems/score-after-flipping-matrix/](https://leetcode-cn.com/problems/score-after-flipping-matrix/) 
## 贪心
```go
func matrixScore(A [][]int) (res int) {
	for row := range A {
		if A[row][0] == 0 {
			flippingRow(A[row])
		}
	}
	for col := range A[0] {
		res = res << 1
		res += countOneInCol(A, col)
	}
	return
}
func countOneInCol(A [][]int, col int) (res int) {
	for row := range A {
		if A[row][col] == 1 {
			res++
		}
	}
	if compare := len(A) - res; compare > res {
		return compare
	}
	return

}
func flippingRow(nums []int) {
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = 1
		} else {
			nums[i] = 0
		}
	}
}
```
>执行用时: 4 ms
内存消耗: 2.2 MB

无限次数的情况下，要找最大的和，那么首先第一列的值要全部为`0`，那么就先遍历第一列，如果第一列的某一个元素为 `0` 就把这个元素对应的列翻转。

然后就是找第`2`列至第`len(A[0])`列了。统计这一列的是`0`多还是`1`多，`0`多就翻转这一列，如此循环最后再计算即可。

但是实际上不需要翻转，只需要统计翻转之后的`1`的个数就可以了。因为翻转的只是这一列，就相当于是局部的贪心，并不影响其他的部分。然后把`1`的个数累加到`res`中，再让`res *= 2`就可以了。这就相当于是统计某个二进制位上的数量有多少。而乘二的操作可以直接换成左移一位。数据范围内不会发生溢出，因此是可行的。