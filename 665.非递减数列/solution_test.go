package leetcode

import (
	"reflect"
	"testing"
)

// 665.非递减数列
// https://leetcode-cn.com/problems/non-decreasing-array/
func checkPossibility(nums []int) bool {
	sum := 0
	for i := range nums[:len(nums)-1] {
		if nums[i] > nums[i+1] {
			if sum > 0 {
				return false
			}
			sum++
			if i > 0 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			} else {
				nums[i] = nums[i+1]
			}
		}
	}

	return true
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want bool
	}{
		{
			nums: []int{3, 4, 2, 3},
			want: false,
		},
		{
			nums: []int{4, 2, 3},
			want: true,
		},
		{
			nums: []int{4, 2, 1},
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := checkPossibility(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
