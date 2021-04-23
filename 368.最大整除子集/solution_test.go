package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 368.最大整除子集
// https://leetcode-cn.com/problems/largest-divisible-subset
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]struct{ index, lenth int }, n)
	for i := range dp {
		dp[i].index = -1
	}

	max := 0
	maxIndex := 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if t := dp[j].lenth + 1; dp[i].lenth < t {
					dp[i].index = j
					dp[i].lenth = t
				}
				if dp[i].lenth > max {
					max = dp[i].lenth
					maxIndex = i
				}
			}
		}
	}

	res := make([]int, 0, max+2)

	for i := maxIndex; i != -1; i = dp[i].index {
		res = append(res, nums[i])
	}

	sort.Ints(res)
	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			nums: []int{1},
			want: []int{1},
		},
		{
			nums: []int{1, 2, 3},
			want: []int{1, 2},
		},
		{
			nums: []int{1, 2, 4, 8},
			want: []int{1, 2, 4, 8},
		},
		{
			nums: []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720},
			want: []int{9, 18, 90, 180, 360, 720},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := largestDivisibleSubset(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
