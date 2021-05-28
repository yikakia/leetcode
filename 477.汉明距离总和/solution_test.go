package leetcode

import (
	"reflect"
	"testing"
)

// 477.汉明距离总和
// https://leetcode-cn.com/problems/total-hamming-distance
func totalHammingDistance(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i <= 30; i++ {
		oneCounts := 0
		for _, num := range nums {
			oneCounts += (num >> i) & 1
		}
		res += oneCounts * (n - oneCounts)
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{4, 14, 2},
			want: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := totalHammingDistance(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
