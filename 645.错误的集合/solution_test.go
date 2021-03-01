package leetcode

import (
	"reflect"
	"testing"
)

// 645.错误的集合
// https://leetcode-cn.com/problems/set-mismatch
func findErrorNums(nums []int) []int {
	var ab int
	for i, num := range nums {
		ab ^= num ^ (i + 1)
	}
	mask := 1
	for mask&ab == 0 {
		mask <<= 1
	}
	var a, b int
	for i, num := range nums {
		if mask&num == 0 {
			a ^= num
		} else {
			b ^= num
		}
		if mask&(i+1) == 0 {
			a ^= i + 1
		} else {
			b ^= i + 1
		}
	}
	for _, num := range nums {
		if num == a {
			return []int{a, b}
		}
	}

	return []int{b, a}
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			nums: []int{2, 2},
			want: []int{2, 1},
		},
		{
			nums: []int{1, 2, 2, 4},
			want: []int{2, 3},
		},
		{
			nums: []int{1, 1},
			want: []int{1, 2},
		},
		{
			nums: []int{3, 2, 2},
			want: []int{2, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findErrorNums(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
