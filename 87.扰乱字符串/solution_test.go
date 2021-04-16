package leetcode

import (
	"reflect"
	"testing"
)

// 87.扰乱字符串
// https://leetcode-cn.com/problems/scramble-string

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(offset1, offset2, length int) int8
	dfs = func(offset1, offset2, length int) (res int8) {
		d := &dp[offset1][offset2][length]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()
		sub1 := s1[offset1 : offset1+length]
		sub2 := s2[offset2 : offset2+length]

		if sub1 == sub2 {
			return 1
		}

		if !sameFreq(sub1, sub2) {
			return 0
		}

		for i := 1; i < length; i++ {
			if dfs(offset1, offset2, i) == 1 && dfs(offset1+i, offset2+i, length-i) == 1 {
				return 1
			}
			if dfs(offset1, offset2+length-i, i) == 1 && dfs(offset1+i, offset2, length-i) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(0, 0, n) == 1
}

func sameFreq(s1, s2 string) bool {
	freq := map[byte]int{}

	for i := range s1 {
		freq[s1[i]]++
	}
	for i := range s2 {
		freq[s2[i]]--
		if freq[s2[i]] < 0 {
			return false
		}
	}
	return true
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		s1, s2 string
		want   bool
	}{
		{
			s1:   "great",
			s2:   "rgeat",
			want: true,
		},
		{
			s1:   "abcde",
			s2:   "caebd",
			want: false,
		},
		{
			s1:   "a",
			s2:   "a",
			want: true,
		},
		{
			s1:   "aabbbc",
			s2:   "bbbaac",
			want: true,
		},
		{
			s1:   "abcdbdacbdac",
			s2:   "bdacabcdbdac",
			want: true,
		},
		{
			s1:   "abcdbdac",
			s2:   "bdacabcd",
			want: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := isScramble(tC.s1, tC.s2)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
