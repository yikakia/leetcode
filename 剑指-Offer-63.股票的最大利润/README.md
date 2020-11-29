# 剑指 Offer 63. 股票的最大利润
[https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/](https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/)
## 维护买点与最大盈利
```go
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	res := 0
	min := prices[0]
	for i := 1; i < n; i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		if prices[i] > min && prices[i]-min > res {
			res = prices[i] - min
		}
	}
	return res
}
```
>执行用时：8 ms, 在所有 Go 提交中击败了24.88%的用户
内存消耗：3 MB, 在所有 Go 提交中击败了83.09%的用户

简单地说就是维护买点与最大盈利。买点是到目前的最小值，盈利是动态维护的，当前减去买点的值和当前的最大盈利的最大值。