package leetcode

import (
	"reflect"
	"testing"
)

// 1143.最长公共子序列
// https://leetcode-cn.com/problems/longest-common-subsequence
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
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
		text1 string
		text2 string
		want  int
	}{
		{
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			text1: "abc",
			text2: "abc",
			want:  3,
		},
		{
			text1: "abc",
			text2: "def",
			want:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := longestCommonSubsequence(tC.text1, tC.text2)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
