package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// 414.第三大的数
// https://leetcode-cn.com/problems/third-maximum-number
func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
			continue
		}
		if num > b && num < a {
			b, c = num, b
			continue
		}
		if num > c && num < b {
			c = num
			continue
		}
	}

	if c == math.MinInt64 {
		return a
	}
	return c
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
		desc string
	}{
		{
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			nums: []int{1, 2},
			want: 2,
		},
		{
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := thirdMax(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
