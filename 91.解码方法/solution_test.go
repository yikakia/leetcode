package leetcode

import (
	"reflect"
	"testing"
)

// 91.解码方法
// https://leetcode-cn.com/problems/decode-ways
func numDecodings(s string) (res int) {
	n := len(s)
	a, b := 0, 1

	for i := 1; i <= n; i++ {
		res = 0
		if s[i-1] != '0' {
			res += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			res += a
		}
		a, b = b, res
	}

	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		want int
	}{
		{
			s:    "20",
			want: 1,
		},
		{
			s:    "10",
			want: 1,
		},
		{
			s:    "1212",
			want: 5,
		},
		{
			s:    "12112",
			want: 8,
		},
		{
			s:    "26",
			want: 2,
		},
		{
			s:    "12",
			want: 2,
		},
		{
			s:    "226",
			want: 3,
		},
		{
			s:    "0",
			want: 0,
		},
		{
			s:    "06",
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numDecodings(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
