package leetcode

import (
	"reflect"
	"testing"
)

// 724.寻找数组的中心索引
// https://leetcode-cn.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, 0
	for _, num := range nums[1:] {
		right += num
	}
	if left == right {
		return 0
	}
	for i := range nums {
		if i != 0 {
			left += nums[i-1]
			right -= nums[i]
		}
		if left == right {
			return i
		}
	}
	return -1
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		nums []int
	}{
		{
			desc: "",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			nums: []int{9},
			want: 0,
		},
		{
			nums: []int{1, 2},
			want: -1,
		},
		{
			nums: []int{1, 0},
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := pivotIndex(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
