package leetcode

import (
	"reflect"
	"testing"
)

// 496.下一个更大元素I
// https://leetcode-cn.com/problems/next-greater-element-i
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextMap := map[int]int{}
	stack := []int{nums2[0]}

	for _, num := range nums2[1:] {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextMap[pre] = num
		}
		stack = append(stack, num)
	}

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		if v, ok := nextMap[num]; ok {
			ans[i] = v
		} else {
			ans[i] = -1
		}
	}
	return ans
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
		desc  string
	}{
		{
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		{
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := nextGreaterElement(tC.nums1, tC.nums2)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
