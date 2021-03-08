package leetcode

import (
	"reflect"
	"testing"
)

// 132.分割回文串II
// https://leetcode-cn.com/problems/palindrome-partitioning-ii

func minCut(s string) int {
	n := len(s)
	isPa := make([][]bool, n)
	for i := range isPa {
		isPa[i] = make([]bool, n)
		for j := range isPa {
			isPa[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			isPa[i][j] = s[i] == s[j] && isPa[i+1][j-1]
		}
	}

	f := make([]int, n)

	for i := range f {
		if isPa[0][i] {
			continue
		}
		f[i] = n
		for j := 0; j < i; j++ {
			if isPa[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}

	return f[n-1]
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		want int
	}{
		{
			s:    "aaabaa",
			want: 1,
		},
		{
			s:    "cabac",
			want: 0,
		},
		{
			s:    "ababababababababababababcbabababababababababababa",
			want: 0,
		},
		{
			s:    "abbaba",
			want: 2,
		},
		{
			s:    "aab",
			want: 1,
		},
		{
			s:    "a",
			want: 0,
		},
		{
			s:    "ab",
			want: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := minCut(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
