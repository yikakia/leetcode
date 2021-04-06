package leetcode

import (
	"reflect"
	"testing"
)

// 80.删除有序数组中的重复项II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	fast, slow := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{0, 1, 2, 2, 2, 2, 2, 3, 4, 4, 4},
			want: 7,
		},
		{
			nums: []int{1, 1, 2},
			want: 3,
		},
		{
			nums: []int{1, 2},
			want: 2,
		},
		{
			nums: []int{},
			want: 0,
		},
		{
			nums: []int{1},
			want: 1,
		},

		{
			nums: []int{1, 1, 1, 1, 1},
			want: 2,
		},
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			want: 5,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: 7,
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
