package leetcode

import (
	"reflect"
	"testing"
)

// 494.目标和
// https://leetcode-cn.com/problems/target-sum
func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 1 {
		res := 0
		if nums[0] == target {
			res++
		}
		if nums[0] == -target {
			res++
		}
		return res
	}
	return findTargetSumWays(nums[1:], target-nums[0]) +
		findTargetSumWays(nums[1:], target+nums[0])
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
		desc   string
	}{

		{
			nums:   []int{1, 0},
			target: 1,
			want:   2,
		},
		{
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			want:   5,
		},
		{
			nums:   []int{1},
			target: 1,
			want:   1,
		},
		{
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
		{
			nums:   []int{1},
			target: -1,
			want:   1,
		},
		{
			nums:   []int{1, 1},
			target: -2,
			want:   1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findTargetSumWays(tC.nums, tC.target)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
