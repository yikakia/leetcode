package leetcode

import (
	"reflect"
	"testing"
)

// 503.下一个更大元素II
// https://leetcode-cn.com/problems/next-greater-element-ii
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	stack := make([]int, 0, n)
	for i := 0; i < 2*n-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 1},
			want: []int{2, -1, 2},
		},
		{
			nums: []int{1, 1, 1},
			want: []int{-1, -1, -1},
		},
		{
			nums: []int{3, 2, 1},
			want: []int{-1, 3, 3},
		},
		{
			nums: []int{},
			want: []int{},
		},
		{
			nums: []int{1},
			want: []int{-1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := nextGreaterElements(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
