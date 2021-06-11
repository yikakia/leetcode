package leetcode

import (
	"reflect"
	"testing"
)

// 518.零钱兑换II
// https://leetcode-cn.com/problems/coin-change-2
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		amount int
		coins  []int
		want   int
		desc   string
	}{
		{
			amount: 5,
			coins:  []int{1, 2, 5},
			want:   4,
		},
		{
			amount: 3,
			coins:  []int{2},
			want:   0,
		},
		{
			amount: 10,
			coins:  []int{10},
			want:   1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := change(tC.amount, tC.coins)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
