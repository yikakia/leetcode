package leetcode

import (
	"reflect"
	"testing"
)

// 395.至少有K个重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestSubstring(s string, k int) (res int) {
	for t := 1; t <= 26; t++ {
		cnt := [26]int{}
		l := 0
		total := 0
		lessK := 0
		for r, ch := range s {
			if cnt[ch-'a'] == 0 {
				total++
				lessK++
			}
			cnt[ch-'a']++
			if cnt[ch-'a'] == k {
				lessK--
			}
			for total > t {
				if cnt[s[l]-'a'] == k {
					lessK++
				}
				cnt[s[l]-'a']--
				if cnt[s[l]-'a'] == 0 {
					total--
					lessK--
				}
				l++
			}
			if lessK <= 0 {
				res = max(res, r-l+1)
			}
		}
	}
	return
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		k    int
		want int
	}{
		{
			s:    "aadbaaaceghadhsfebacegccdbebeabb",
			k:    3,
			want: 3,
		},
		{
			s:    "aaabb",
			k:    3,
			want: 3,
		},
		{
			s:    "ababbc",
			k:    2,
			want: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := longestSubstring(tC.s, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
