package leetcode

import (
	"reflect"
	"testing"
)

// 1480.一维数组的动态和
// https://leetcode-cn.com/problems/running-sum-of-1d-array
func runningSum(nums []int) []int {
	for i := range nums[1:] {
		nums[i+1] += nums[i]
	}
	return nums
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
		desc string
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: []int{1, 3, 6, 10},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := runningSum(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
