package leetcode

import (
	"reflect"
	"testing"
)

// 485.最大连续1的个数
// https://leetcode-cn.com/problems/max-consecutive-ones/

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func findMaxConsecutiveOnes(nums []int) (res int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] == 0 {
			res = max(j-i, res)
			i = j + 1
		} else if j == len(nums)-1 {
			res = max(j-i+1, res)
		}
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{0, 1, 1, 0, 1, 1, 1, 0},
			want: 3,
		},
		{
			nums: []int{0, 1, 1, 0, 1, 0, 1, 1},
			want: 2,
		},
		{
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findMaxConsecutiveOnes(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
