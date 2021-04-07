package leetcode

import (
	"reflect"
	"testing"
)

// 33.搜索旋转排序数组
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if r == -1 {
		return -1
	}
	if r == 0 {
		if nums[r] != target {
			return -1
		}
		return 0
	}
	mid := (l + r) / 2
	switch {
	case nums[mid] == target:
		return mid
	case nums[mid] > nums[r]:
		index := bSearch(nums[:mid], target)
		if index != -1 {
			return index
		}
		if index = search(nums[mid+1:], target); index == -1 {
			return -1
		}
		return mid + 1 + index
	default:
		index := bSearch(nums[mid:], target)
		if index != -1 {
			return mid + index
		}
		return search(nums[:mid], target)
	}
}

func bSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return -1
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{3, 1},
			target: 3,
			want:   0,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			nums:   []int{1},
			target: 0,
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
