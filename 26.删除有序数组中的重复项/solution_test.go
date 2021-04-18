package leetcode

import (
	"reflect"
	"testing"
)

// 26.删除有序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	i := 1
	for j := 1; j < n; j++ {
		if nums[j] != nums[i-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{1, 2},
			want: 2,
		},
		{
			nums: []int{0, 0, 0, 0},
			want: 1,
		},
		{
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3},
			want: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := removeDuplicates(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
