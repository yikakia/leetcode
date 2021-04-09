package leetcode

import (
	"reflect"
	"testing"
)

// 154.寻找旋转排序数组中的最小值II
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if right < 1 {
		return nums[0]
	}
	if right == 1 {
		return min(nums[0], nums[1])
	}

	prev := func(i int) int {
		if i > 0 {
			return nums[i-1]
		}
		return nums[len(nums)-1]
	}

	mid := right / 2

	for left <= right {
		switch {
		case nums[mid] < prev(mid):
			return nums[mid]
		case nums[mid] > nums[right]:
			left = mid + 1
		case nums[mid] < nums[right]:
			right = mid
		case nums[mid] == nums[right]:
			right--
		default:
		}
		mid = (left + right) / 2
	}

	return nums[left]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 2, 2, 3, 3, 4, 1, 1},
			want: 1,
		},
		{
			nums: []int{1, 1, 1},
			want: 1,
		},
		{
			nums: []int{1, 3, 5},
			want: 1,
		},
		{
			nums: []int{2, 2, 2, 0, 1},
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findMin(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
