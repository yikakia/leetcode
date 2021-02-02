package leetcode

import (
	"reflect"
	"testing"
)

// 424.替换后的最长重复字符
// https://leetcode-cn.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		k    int
		want int
	}{
		{
			s:    "BAAAB",
			k:    2,
			want: 5,
		},
		{
			s:    "ABBB",
			k:    2,
			want: 4,
		},
		{
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			s:    "AABABBABB",
			k:    1,
			want: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := characterReplacement(tC.s, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
