package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

// 345.反转字符串中的元音字母
// https://leetcode-cn.com/problems/reverse-vowels-of-a-string
func reverseVowels(s string) string {
	t := []byte(s)
	n := len(t)
	i, j := 0, n-1
	for i < j {
		for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
			i++
		}
		for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}
	}
	return string(t)
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		s    string
		want string
		desc string
	}{
		{
			s:    "hello",
			want: "holle",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := reverseVowels(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
