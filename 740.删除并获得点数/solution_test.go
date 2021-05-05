package leetcode

import (
	"reflect"
	"testing"
)

// 740.删除并获得点数
// https://leetcode-cn.com/problems/delete-and-earn
func deleteAndEarn(nums []int) int {
	var maxNum int
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}
	dp := make([]int, maxNum+1)
	for _, num := range nums {
		dp[num]++
	}
	return rob(dp)
}

func rob(dp []int) int {
	pre1, pre2 := dp[0], max(dp[0], dp[1])

	for i := 2; i < len(dp); i++ {
		pre1, pre2 = pre2, max(pre1+dp[i]*i, pre2)
	}
	return pre2
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{3, 4, 2},
			want: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := deleteAndEarn(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
