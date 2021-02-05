package leetcode

import (
	"reflect"
	"testing"
)

// 643.子数组最大平均数I
// https://leetcode-cn.com/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	max := 0
	for i := range nums[:k] {
		sum += nums[i]
	}
	max = sum
	for i := range nums[k:] {
		sum += nums[i+k] - nums[i]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		k    int
		want float64
	}{
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findMaxAverage(tC.nums, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
