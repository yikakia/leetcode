package leetcode

import (
	"reflect"
	"testing"
)

// 剑指Offer38.字符串的排列
// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof
func permutation(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	r := make([]string, 0)
	hasApper := make(map[rune]bool, 0)
	for index, char := range s {
		if hasApper[char] {
			continue
		}
		tmpr := permutation(s[:index] + s[index+1:])
		for i, word := range tmpr {
			tmpr[i] = string(char) + word
		}
		r = append(r, tmpr...)
		hasApper[char] = true
	}
	return r
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		s    string
		want []string
		desc string
	}{
		{s: "abc",
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := permutation(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
