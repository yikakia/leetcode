package leetcode

import (
	"reflect"
	"testing"
)

// 28.实现strStr()
// https://leetcode-cn.com/problems/implement-strstr
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i := range haystack[:len(haystack)-len(needle)+1] {
		if haystack[i] != needle[0] {
			continue
		}
		exist := true
		for j := range needle {
			if haystack[i+j] != needle[j] {
				exist = false
				break
			}
		}
		if exist {
			return i
		}
	}

	return -1
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc             string
		haystack, needle string
		want             int
	}{
		{
			haystack: "mississippi",
			needle:   "issip",
			want:     4,
		},
		{
			haystack: "abcdefg",
			needle:   "efg",
			want:     4,
		},
		{
			haystack: "a",
			needle:   "a",
			want:     0,
		},
		{
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
		{
			haystack: "",
			needle:   "",
			want:     0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := strStr(tC.haystack, tC.needle)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
