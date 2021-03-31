package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 90.子集II
// https://leetcode-cn.com/problems/subsets-ii
func subsetsWithDup(nums []int) (res [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res = append(res, []int{})
	n := 0
	for i := range nums {
		preN := n
		var start, end int
		n = len(res)
		if i > 0 && nums[i] == nums[i-1] {
			start, end = preN, n
		} else {
			start, end = 0, n
		}
		for j := start; j < end; j++ {
			t := append([]int(nil), res[j]...)
			t = append(t, nums[i])
			res = append(res, t)
		}
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			nums: []int{2, 2, 2},
			want: [][]int{{}, {2}, {2, 2}, {2, 2, 2}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := subsetsWithDup(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
