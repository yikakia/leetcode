package leetcode

import (
	"reflect"
	"testing"
)

// 453.最小操作次数使数组元素相等
// https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements
func minMoves(nums []int) int {
	minNum := min(nums)
	ans := 0

	for _, num := range nums {
		ans += num - minNum
	}

	return ans
}

func min(nums []int) int {
	ans := nums[0]

	for _, num := range nums[1:] {
		if ans > num {
			ans = num
		}
	}

	return ans
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
		desc string
	}{
		{
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			nums: []int{1, 1, 1},
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := minMoves(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
