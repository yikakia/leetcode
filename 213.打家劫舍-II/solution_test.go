package leetcode

import (
	"reflect"
	"testing"
)

// 213.打家劫舍II
// https://leetcode-cn.com/problems/house-robber-ii
func rob(nums []int) int {
	n := len(nums)
	if n < 4 {
		return max(nums...)
	}
	dp := make([]int, n)  // 抢 i
	dp2 := make([]int, n) // 不抢 i
	res1, res2 := 0, 0    // res1 表示选第一个的最大结果 res2 表示不选第一个的最大的结果

	dp[0] = nums[0]
	for i := 1; i < n-1; i++ {
		dp[i] = nums[i] + max(dp[(i+n-2)%n], dp2[(i+n-1)%n])
		dp2[i] = max(dp[(i+n-1)%n], dp2[(i+n-1)%n])
	}
	dp2[n-1] = max(dp[n-2], dp2[n-2])
	res1 = max(max(dp[n-3:]...), max(dp2[n-3:]...))

	dp = make([]int, n)
	dp2 = make([]int, n)

	for i := 1; i < n; i++ {
		dp[i] = nums[i] + max(dp[(i+n-2)%n], dp2[(i+n-1)%n])
		dp2[i] = max(dp[(i+n-1)%n], dp2[(i+n-1)%n])
	}
	res2 = max(max(dp[n-3:]...), max(dp2[n-3:]...))

	return max(res1, res2)
}

func max(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if res < num {
			res = num
		}
	}
	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			nums: []int{0},
			want: 0,
		},
		{
			nums: []int{200, 3, 140, 20, 10},
			want: 340,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := rob(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
