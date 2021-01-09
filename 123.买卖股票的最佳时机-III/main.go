package main

import (
	"fmt"
	"math"
	"reflect"
)

// 123.买卖股票的最佳时机III
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
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
func test() {
	type TestType struct {
		prices []int
		want   int
	}
	ts := []TestType{
		TestType{
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   6,
		},
		TestType{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		TestType{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		TestType{
			prices: []int{1},
			want:   0,
		},
	}
	for _, t := range ts {
		get := maxProfit(t.prices)
		if !reflect.DeepEqual(get, t.want) {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")
}
func main() {
	test()
}
