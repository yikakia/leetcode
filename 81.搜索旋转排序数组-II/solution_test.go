package leetcode

import (
	"reflect"
	"testing"
)

// 81.搜索旋转排序数组II
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii
func search(nums []int, target int) bool {
	mid, n := len(nums)/2, len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[mid] == target
	}
	if mid == 0 {
		return false
	}
	switch {
	case nums[mid] > nums[0]:
		return bSearch(nums[:mid+1], target) || search(nums[mid+1:], target)
	case nums[mid] < nums[n-1]:
		return bSearch(nums[mid:], target) || search(nums[:mid], target)
	default:
		return search(nums[:mid], target) || search(nums[mid:], target)
	}
}

func bSearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		want   bool
	}{
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			want:   true,
		},
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			want:   false,
		},
		{
			nums:   []int{2, 2},
			target: 3,
			want:   false,
		},
		{
			nums:   []int{2},
			target: 3,
			want:   false,
		},
		{
			nums:   []int{22, 22, 22},
			target: 3,
			want:   false,
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
