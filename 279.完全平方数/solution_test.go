package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// 279.完全平方数
// https://leetcode-cn.com/problems/perfect-squares

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j])
		}
		dp[i]++
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		n    int
		want int
		desc string
	}{
		{
			n:    12,
			want: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numSquares(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
