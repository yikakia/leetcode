package leetcode

import (
	"reflect"
	"testing"
)

// 1035.不相交的线
// https://leetcode-cn.com/problems/uncrossed-lines
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i, v := range nums1 {
		for j, w := range nums2 {
			if v == w {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{1, 4, 2},
			nums2: []int{1, 2, 4},
			want:  2,
		},
		{
			nums1: []int{2, 5, 1, 2, 5},
			nums2: []int{10, 5, 2, 1, 5, 2},
			want:  3,
		},
		{
			nums1: []int{1, 3, 7, 1, 7, 5},
			nums2: []int{1, 9, 2, 5, 1},
			want:  2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := maxUncrossedLines(tC.nums1, tC.nums2)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
