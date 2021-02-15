# 485.最大连续1的个数
[https://leetcode-cn.com/problems/max-consecutive-ones/](https://leetcode-cn.com/problems/max-consecutive-ones/) 
## 原题
给定一个二进制数组， 计算其中最大连续1的个数。

 **示例 1:** 

```

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.

```
 **注意：** 
- 输入的数组只包含 `0` 和 `1`。
- 输入数组的长度是正整数，且不超过 10,000。
 
**标签**
`数组` 


## 
```go
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func findMaxConsecutiveOnes(nums []int) (res int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] == 0 {
			res = max(j-i, res)
			i = j + 1
		} else if j == len(nums)-1 {
			res = max(j-i+1, res)
		}
	}
	return
}
```
>执行用时: 48 ms
内存消耗: 6.5 MB

用窗口记录长度。要注意的是最后的时候可能没有更新，所以需要最后再判断一次。