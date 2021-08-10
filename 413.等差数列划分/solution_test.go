package leetcode

import (
	"reflect"
	"testing"
)

// 413.等差数列划分
// https://leetcode-cn.com/problems/arithmetic-slices
func numberOfArithmeticSlices(nums []int) (ans int) {
	n := len(nums)
	if n < 3 {
		return
	}

	var delta, count = nums[1] - nums[0], 0
	for end := 2; end < n; end++ {
		if delta == nums[end]-nums[end-1] {
			count++
			ans += count
		} else {
			delta = nums[end] - nums[end-1]
			count = 0
		}
	}

	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
		desc string
	}{
		{
			nums: []int{1, 2, 3, 4, 3, 4, 5, 6, 1, 2, 3},
			want: 7,
		},
		{
			nums: []int{1, 2, 3, 8, 9, 10},
			want: 2,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: 3,
		},
		{
			nums: []int{1},
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numberOfArithmeticSlices(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
