package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 561.数组拆分I
// https://leetcode-cn.com/problems/array-partition-i/
func arrayPairSum(nums []int) (res int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{1, 4, 3, 2},
			want: 4,
		},
		{
			nums: []int{6, 2, 6, 5, 1, 2},
			want: 9,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := arrayPairSum(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
