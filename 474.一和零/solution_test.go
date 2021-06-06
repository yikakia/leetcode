package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

// 474.一和零
// https://leetcode-cn.com/problems/ones-and-zeroes
func findMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		strs []string
		m, n int
		want int
		desc string
	}{
		{
			strs: []string{"10", "0001", "111001", "1", "0"},
			m:    5,
			n:    3,
			want: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findMaxForm(tC.strs, tC.m, tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
