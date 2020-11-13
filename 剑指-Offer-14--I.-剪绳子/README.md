# 剑指 Offer 14- I. 剪绳子
[https://leetcode-cn.com/problems/jian-sheng-zi-lcof/](https://leetcode-cn.com/problems/jian-sheng-zi-lcof/)

## 动规
```go
func cuttingRope(n int) int {
	dp := make([]int, n+2)
	dp[1] = 1
	dp[2] = 1
	for i := 2; i < n+1; i++ {
		r := dp[1]
		for j := 1; j < i; j++ {
			r = max(r, max(j, dp[j])*max(i-j, dp[i-j]))
		}
		dp[i] = r
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了36.72%的用户

动规的思想是前面的解可以被后面使用。其中最重要的是找到合适的状态转移方程。对于本题而言，长度为`n`的绳子的切法可以利用长度为`n-1`的绳子的切法作为参考。

对于动规一直不是很懂，之后需要加强这方面的练习。

## 数论
```go
func cuttingRope(n int) int {
	r := 1
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	rest := n % 3
	for n > 2 {
		r *= 3
		n -= 3
	}
	switch rest {
	case 1:
		r *= 4
		r /= 3
	case 2:
		r *= 2
	default:
	}
	return r
}
```
从数学上看这个问题的话要先知道几个前提

- 任意大于2的整数都可以表示为2和3的组合
- `3`>`2*1`(长度为3时切分成3比2和1好)
- `2*2`>`3*1` (长度为4时切分成两个2比1和3好)
- `3*2`>`3*1*1`>`2*2*1`(长度为5时切分成3*2最优)
- `3*3`>`3*2*1`>`2*2*2`(长度为6时切分成两个3最优)

总的来说就是对2，3特殊处理，对大于3的n，先竟可能多的分3，再根据n除以3的余数r来求得最后的结果。

- r = 1 应该把`3*1`切分成`2*2`
- r = 2 应该直接乘以2