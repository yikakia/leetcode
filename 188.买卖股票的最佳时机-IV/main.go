package main

import (
	"fmt"
	"math"
)

// 188.买卖股票的最佳时机IV
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/

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
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

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
func main() {
	test()
}
