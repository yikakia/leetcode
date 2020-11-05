# 剑指 Offer 49. 丑数
[https://leetcode-cn.com/problems/chou-shu-lcof/](https://leetcode-cn.com/problems/chou-shu-lcof/)

## 动规 
```go

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	dpa, dpb, dpc := 0, 0, 0
	for i := 1; i < n; i++ {
		dpa, dpb, dpc = dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(dpa, min(dpb, dpc))
		if dp[i] == dpa {
			a++
		}
		if dp[i] == dpb {
			b++
		}
		if dp[i] == dpc {
			c++
		}
	}
	return dp[n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：4.2 MB, 在所有 Go 提交中击败了62.92%的用户

记 `dp[n]` 为第 `n` 个丑数。那么要求`dp[n+1]`的话，我们可以知道，它一定是等于`dp[a]*2`或`dp[b]*3`或`dp[c]*5`的最小值（这三者可以同时为真，比如`30`就可以由 `15*2 10*3 6*5`这三种组成）。而`a`,`b`,`c`则可以记为第`a,b,c`个丑数。对`a`而言它满足`dp[a]*2`大于等于`dp[n+1]`。对于`b`,`c`可以类推。同时每次找到对应的位置时，都往后面移动一次。