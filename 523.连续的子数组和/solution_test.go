package leetcode

import (
	"reflect"
	"testing"
)

// 523.连续的子数组和
// https://leetcode-cn.com/problems/continuous-subarray-sum

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	preExist := map[int]int{0: -1}
	preSum := 0
	for i, v := range nums {
		preSum += v
		preSum %= k
		if preIndex, ok := preExist[preSum]; ok {
			if i-preIndex >= 2 {
				return true
			}
		} else {
			preExist[preSum] = i
		}
	}
	return false
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{0, 0},
			k:    1,
			want: true,
		},
		{
			nums: []int{23, 2, 4, 6, 7},
			k:    6,
			want: true,
		},
		{
			nums: []int{23, 2, 6, 4, 7},
			k:    6,
			want: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := checkSubarraySum(tC.nums, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
