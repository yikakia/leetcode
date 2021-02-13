package leetcode

import (
	"reflect"
	"testing"
)

// 448.找到所有数组中消失的数字
// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for _, vi := range nums {
		if vi != 0 {
			var cur int
			for j := vi - 1; nums[j] != 0; {
				cur = nums[j] - 1
				nums[j] = 0
				j = cur
			}
		}
	}
	for i := range nums {
		if nums[i] != 0 {
			res = append(res, i+1)
		}
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
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findDisappearedNumbers(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
