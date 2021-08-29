package leetcode

import (
	"reflect"
	"testing"
)

// 1588.所有奇数长度子数组的和
// https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays
func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	preSum := make([]int, len(arr)+1)
	copy(preSum[1:], arr)
	for i := 1; i < len(preSum); i++ {
		preSum[i] += preSum[i-1]
	}
	for i := 1; i < len(preSum); i++ {
		for j := i; j < len(preSum); j += 2 {
			sum += preSum[j] - preSum[i-1]
		}
	}
	return sum
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
		desc string
	}{
		{
			arr:  []int{1, 4, 2, 5, 3},
			want: 58,
		},
		{
			arr:  []int{10, 11, 12},
			want: 66,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := sumOddLengthSubarrays(tC.arr)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
