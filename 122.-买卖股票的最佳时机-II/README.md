# 122. 买卖股票的最佳时机 II
[https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)
## 找最低点和最高点
```go
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	buy, sell := -1, -1
	cont := 0
	if prices[0] <= prices[1] {
		buy = prices[0]
	}
	for i := 1; i < n; {
		if buy < 0 {
			for i+1 < n && !(prices[i] <= prices[i+1] && prices[i] <= prices[i-1]) {
				i++
			}
			if i+1 < n && (prices[i] <= prices[i+1] && prices[i] <= prices[i-1]) {
				buy = prices[i]
			}
		}
		if sell < 0 {
			for i+1 < n && !(prices[i] >= prices[i+1] && prices[i] >= prices[i-1]) {
				i++
			}
			if i+1 < n && (prices[i] >= prices[i+1] && prices[i] >= prices[i-1]) {
				sell = prices[i]
			}
		}
		if buy >= 0 && sell >= 0 {
			cont += sell - buy
			sell, buy = -1, -1
		}
		i++
	}
	if prices[n-1] > prices[n-2] && prices[n-1] > buy && buy != -1 {
		cont += prices[n-1] - buy
	}
	return cont
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：3 MB, 在所有 Go 提交中击败了76.89%的用户

写完看别人写的才发现自己想复杂了。就是很简单的算每个上升区间就可以了。
```go
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	maxPro := 0
	for i := 1; i < n; i++ {
		if prices[i]-prices[i-1] > 0 {
			maxPro += prices[i] - prices[i-1]
		}
	}

	return maxPro
}
```

就是这样，很简单，根本不用模拟实际的股票操作找买点和卖点，只用算上升区间上升了多少就够了。