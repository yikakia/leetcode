package leetcode

import (
	"reflect"
	"testing"
)

// 446.等差数列划分II-子序列
// https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence
func numberOfArithmeticSlices(nums []int) (ans int) {
	n := len(nums)
	everyMap := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		everyMap[i] = map[int]int{}
	}
	for i, num := range nums {
		for j := i - 1; j >= 0; j-- {
			d := num - nums[j]
			ans += everyMap[j][d]
			everyMap[i][d] += everyMap[j][d] + 1
		}
	}

	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
		desc string
	}{
		{
			nums: []int{2, 4, 6, 8, 10},
			want: 7,
		},
		{
			nums: []int{7, 7, 7, 7, 7},
			want: 16,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numberOfArithmeticSlices(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
