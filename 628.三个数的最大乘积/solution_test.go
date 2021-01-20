package leetcode

import (
	"reflect"
	"testing"
)

// 628.三个数的最大乘积
// https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
func maximumProduct(nums []int) int {
	const m = 1001
	max1, max2, max3 := -m, -m, -m
	min1, min2 := m, m
	for _, x := range nums {
		if x >= max1 {
			max1, max2, max3 = x, max1, max2
		} else if x >= max2 {
			max2, max3 = x, max2
		} else if x >= max3 {
			max3 = x
		}
		if x <= min1 {
			min1, min2 = x, min1
		} else if x <= min2 {
			min2 = x
		}
	}
	if min2 == m {
		min2 = 0
	}
	if max1 < 0 {
		return max1 * max2 * max3
	}

	return max1 * max(max2*max3, min1*min2)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		nums []int
	}{
		{
			nums: []int{-1, -2, -3, -4},
			want: -6,
		},
		{
			nums: []int{-100, -98, -1, 2, 3, 4},
			want: 39200,
		},
		{
			desc: "",
			nums: []int{1, 2, 3},
			want: 6,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: 24,
		},
		{
			nums: []int{1, 3, 4, -5, -8},
			want: 160,
		},
		{
			nums: []int{4, 3, 2, 1},
			want: 24,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := maximumProduct(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
