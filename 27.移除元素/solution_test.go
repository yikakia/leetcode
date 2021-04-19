package leetcode

import (
	"reflect"
	"testing"
)

// 27.移除元素
// https://leetcode-cn.com/problems/remove-element
func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	for l < r {
		if nums[l] == val {
			if nums[r] != val {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			} else {
				for r > l && nums[r] == val {
					r--
				}
			}
		} else {
			for l < r && nums[l] != val {
				l++
			}
		}
	}
	if nums[l] == val {
		return l
	}
	return l + 1
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		val  int
		want int
	}{
		{
			nums: []int{4, 5},
			val:  4,
			want: 1,
		},

		{
			nums: []int{1, 4, 5},
			val:  4,
			want: 2,
		},
		{
			nums: []int{},
			val:  3,
			want: 0,
		},
		{
			nums: []int{3, 3},
			val:  3,
			want: 0,
		},
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := removeElement(tC.nums, tC.val)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
