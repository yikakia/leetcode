package leetcode

import (
	"reflect"
	"testing"
)

// 704.二分查找
// https://leetcode-cn.com/problems/binary-search
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
		desc   string
	}{
		{
			nums:   []int{2, 5},
			target: 5,
			want:   1,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := search(tC.nums, tC.target)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
