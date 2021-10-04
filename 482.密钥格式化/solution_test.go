package leetcode

import (
	"reflect"
	"testing"
	"unicode"
)

// 482.密钥格式化
// https://leetcode-cn.com/problems/license-key-formatting
func licenseKeyFormatting(s string, k int) string {
	ans := []byte{}
	for i, cnt := len(s)-1, 0; i >= 0; i-- {
		if s[i] != '-' {
			ans = append(ans, byte(unicode.ToUpper(rune(s[i]))))
			cnt++
			if cnt%k == 0 {
				ans = append(ans, '-')
			}
		}
	}
	if len(ans) > 0 && ans[len(ans)-1] == '-' {
		ans = ans[:len(ans)-1]
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		s    string
		k    int
		want string
		desc string
	}{
		{
			s:    "---",
			k:    1,
			want: "",
		},
		{
			s:    "rr",
			k:    1,
			want: "R-R",
		},
		{
			s:    "rrdd",
			k:    2,
			want: "RR-DD",
		},
		{
			s:    "r",
			k:    1,
			want: "R",
		},
		{
			s:    "5F3Z-2e-9-w",
			k:    4,
			want: "5F3Z-2E9W",
		},
		{
			s:    "2-5g-3-J",
			k:    2,
			want: "2-5G-3J",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := licenseKeyFormatting(tC.s, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
