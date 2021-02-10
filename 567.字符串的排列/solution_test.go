package leetcode

import (
	"reflect"
	"testing"
)

// 567.字符串的排列
// https://leetcode-cn.com/problems/permutation-in-string/

func isSame(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)

	if n > m {
		return false
	}
	cnt1 := [26]int{}
	cnt2 := [26]int{}

	for i := 0; i < n; i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}

	for i := n; i < m; i++ {
		if isSame(cnt1, cnt2) {
			return true
		}
		cnt2[s2[i-n]-'a']--
		cnt2[s2[i]-'a']++
	}

	return isSame(cnt1, cnt2)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s1   string
		s2   string
		want bool
	}{
		{
			s1:   "adc",
			s2:   "dcda",
			want: true,
		},
		{
			s1:   "a",
			s2:   "b",
			want: false,
		},
		{
			s1:   "a",
			s2:   "a",
			want: true,
		},
		{
			s1:   "bbba",
			s2:   "aab",
			want: false,
		},
		{
			s1:   "ab",
			s2:   "bbaa",
			want: true,
		},
		{
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := checkInclusion(tC.s1, tC.s2)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
