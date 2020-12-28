# 188. 买卖股票的最佳时机 IV
[https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/) 
## 原题
给定一个整数数组 `prices` ，它的第`i` 个元素 `prices[i]` 是一支给定的股票在第 `i` 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 **k**  笔交易。
**注意：** 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 
**示例 1：** 
```
输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
```
**示例 2：** 
```
输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
```
 
**提示：** 
- `0 <= k <= 10^9`
- `0 <= prices.length <= 1000`
- `0 <= prices[i] <= 1000`


## 动规
```go
func max(nums ...int) (res int) {
	res = nums[0]
	for _, num := range nums[1:] {
		if res < num {
			res = num
		}
	}
	return
}
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := make([][]int, n)
	sell := make([][]int, n)
	if k > n/2 {
		k = n / 2
	}
	for i := range prices {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	
	for i := range buy[0]  {
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
func test() {
	type TestType struct {
		k      int
		prices []int
		want   int
	}
	ts := []TestType{
		TestType{k: 2, prices: []int{2, 4, 1}, want: 2},
		TestType{k: 2, prices: []int{3, 2, 6, 5, 0, 3}, want: 7},
	}
	for _, t := range ts {
		get := maxProfit(t.k, t.prices)
		if get != t.want {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")

}
```
>执行用时: 4 ms
内存消耗: 4.7 MB

这个买股票的问题说难不难，说简单不简单。因为重要的就是写出状态转移方程。剩下的实现就很简单了。

对于这类题而言，通常可以开两个数组分别用来记录当天结束后持有或不持有股票时的最大利润。

那么第一天的`buy[0]`自然就是`-prices[0]` 。怎么更新的呢？自然要从前一天开始判断，目的是得到最大利润，那么就可以对比前一天的`buy[i-1]`和 `sell[i-1]-prices[i]`中大的那一个。这里就是模拟更新买点了。同样的可以类比到`sell`的更新上，这里就是更新卖点。