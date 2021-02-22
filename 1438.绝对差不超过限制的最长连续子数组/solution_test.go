package leetcode

import (
	"reflect"
	"testing"
)

// 1438.绝对差不超过限制的最长连续子数组
// https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

func longestSubarray(nums []int, limit int) (ans int) {
	var minQ, maxQ []int
	left := 0
	for right, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		nums  []int
		limit int
		want  int
	}{
		// {
		// 	nums:  []int{2, 2, 4, 8, 3, 4, 5, 2, 1, 6, 8, 7, 4},
		// 	limit: 4,
		// 	want:  5,
		// },
		{
			nums:  []int{8, 2, 4, 7},
			limit: 4,
			want:  2,
		},
		{
			nums:  []int{10, 1, 2, 4, 7, 2},
			limit: 5,
			want:  4,
		},
		{
			nums:  []int{4, 2, 2, 2, 4, 4, 2, 2},
			limit: 0,
			want:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := longestSubarray(tC.nums, tC.limit)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
