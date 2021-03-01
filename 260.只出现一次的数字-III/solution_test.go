package leetcode

import (
	"reflect"
	"testing"
)

// 260.只出现一次的数字III
// https://leetcode-cn.com/problems/single-number-iii
func singleNumber(nums []int) []int {
	var ab, a, b int
	for _, num := range nums {
		ab ^= num
	}
	var mask int = 1
	for mask&ab == 0 {
		mask <<= 1
	}
	for _, num := range nums {
		if num&mask == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 1, 3, 2, 5},
			want: []int{3, 5},
		},
		{
			nums: []int{-1, 0},
			want: []int{-1, 0},
		},
		{
			nums: []int{0, 1},
			want: []int{1, 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := singleNumber(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
