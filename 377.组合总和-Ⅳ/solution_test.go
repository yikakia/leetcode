package leetcode

import (
	"reflect"
	"testing"
)

// 377.组合总和Ⅳ
// https://leetcode-cn.com/problems/combination-sum-iv
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 2, 3, 7, 8},
			target: 10,
			want:   291,
		},
		{
			nums:   []int{1, 2, 3, 7, 8},
			target: 14,
			want:   3499,
		},
		{
			nums:   []int{1, 2, 3},
			target: 4,
			want:   7,
		},
		{
			nums:   []int{9},
			target: 3,
			want:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := combinationSum4(tC.nums, tC.target)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
