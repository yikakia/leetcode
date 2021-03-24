package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// 456.132模式
// https://leetcode-cn.com/problems/132-pattern
func find132pattern(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64

	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}

	return false
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			nums: []int{3, 1, 4, 2},
			want: true,
		},
		{
			nums: []int{-1, 3, 2, 0},
			want: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := find132pattern(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
