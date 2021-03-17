package leetcode

import (
	"reflect"
	"testing"
)

// 115.不同的子序列
// https://leetcode-cn.com/problems/distinct-subsequences
func numDistinct(s string, t string) int {
	ls, lt := len(s), len(t)
	if ls < lt {
		return 0
	}
	dp := make([][]int, ls)
	for i := range dp {
		dp[i] = make([]int, lt)
	}
	var maxInColAtRow func(col int, row int) int = func(col, row int) int {
		for i := row; i >= 0; i-- {
			if dp[i][col] > 0 {
				return dp[i][col]
			}
		}
		return 0
	}
	for i := 0; i < ls; i++ {
		for j := 0; j < lt && j <= i; j++ {
			if s[i] == t[j] {
				if j == 0 {
					dp[i][j] = 1 + maxInColAtRow(j, i-1)
				} else {
					dp[i][j] = maxInColAtRow(j-1, i-1) + maxInColAtRow(j, i-1)
				}
			}
		}
	}
	return maxInColAtRow(lt-1, ls-1)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s, t string
		want int
	}{
		{
			s:    "babgbag",
			t:    "bag",
			want: 5,
		},
		{s: "rabbbit",
			t:    "rabbit",
			want: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numDistinct(tC.s, tC.t)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
