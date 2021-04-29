package leetcode

import (
	"reflect"
	"testing"
)

// 403.青蛙过河
// https://leetcode-cn.com/problems/frog-jump
func canCross(stones []int) bool {
	n := len(stones)
	for i := 1; i < n; i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		stones []int
		want   bool
	}{
		{
			stones: []int{0, 1, 3, 5, 6, 8, 12, 17},
			want:   true,
		},
		{
			stones: []int{0, 1, 2, 3, 4, 8, 9, 11},
			want:   false,
		},
		{
			stones: []int{0, 1, 2, 3, 4, 8, 9, 11},
			want:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := canCross(tC.stones)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
