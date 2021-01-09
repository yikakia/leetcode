# 123.买卖股票的最佳时机 III
[https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/) 
## 原题
给定一个数组，它的第`i` 个元素是一支给定的股票在第 `i`天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 **两笔** 交易。
**注意：** 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 
**示例 1:** 
```
输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
```
**示例 2：** 
```
输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。   
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。   
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
```
**示例 3：** 
```
输入：prices = [7,6,4,3,1] 
输出：0 
解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
```
**示例 4：** 
```
输入：prices = [1]
输出：0
```
 
**提示：** 
- `1 <= prices.length <= 10^5`
- `0 <= prices[i] <= 10^5`
 
**标签**
`数组` `动态规划` 


## 
```go
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := make([][]int, n)
	sell := make([][]int, n)

	k := 2
	for i := 0; i < n; i++ {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}

	for i := range buy[0] {
		buy[0][i] = math.MinInt64 >> 1
		sell[0][i] = math.MinInt64 >> 1
	}

	buy[0][0] = -prices[0]
	sell[0][0] = 0

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	return max(sell[n-1]...)
}

func max(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if num > res {
			res = num
		}
	}
	return res
}
```
>执行用时: 248 ms
内存消耗: 31.3 MB

这道题和 [188. 买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/) 基本相似，不过本题是它的弱化版本，因为本题的`k`是指定的。因此可以看那道题的[解析](../188.买卖股票的最佳时机-IV/README.md)