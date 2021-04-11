# 264.丑数 II
[https://leetcode-cn.com/problems/ugly-number-ii](https://leetcode-cn.com/problems/ugly-number-ii) 
## 原题
给你一个整数 `n` ，请你找出并返回第 `n` 个 **丑数** 。

 **丑数** 就是只包含质因数  `2` 、 `3` 和/或  `5`  的正整数。

 

 **示例 1：** 

```

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。

```
 **示例 2：** 

```

输入：n = 1
输出：1
解释：1 通常被视为丑数。

```
 

 **提示：** 
-  `1 <= n <= 1690` 
 
**标签**
`堆` `数学` `动态规划` 


## 
```go
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	uglyNums := make([]int, n)
	uglyNums[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		x2, x3, x5 := uglyNums[p2]*2, uglyNums[p3]*3, uglyNums[p5]*5
		uglyNums[i] = min(x2, x3, x5)
		if uglyNums[i] == x2 {
			p2++
		}
		if uglyNums[i] == x3 {
			p3++
		}
		if uglyNums[i] == x5 {
			p5++
		}
	}
	return uglyNums[n-1]
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if res > num {
			res = num
		}
	}
	return res
}
```
>执行用时: 0 ms
内存消耗: 4.2 MB

简单地说就是记录上一次用 2|3|5 的下标，然后当此时更新的值与用上次的结果乘出来的相同，则进行移动到下一个下标。