package leetcode

import (
	"reflect"
	"testing"
)

// 434.字符串中的单词数
// https://leetcode-cn.com/problems/number-of-segments-in-a-string
func countSegments(s string) int {
	var flag bool
	var res int
	for i := range s {
		if s[i] == ' ' {
			if flag {
				res++
			}
			flag = false
		} else {
			flag = true
		}
	}
	if flag {
		res++
	}
	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		s    string
		want int
		desc string
	}{
		{
			s:    "Hello, my name is John",
			want: 5,
		},
		{
			s:    "  ",
			want: 0,
		},
		{
			s:    "a  ",
			want: 1,
		},
		{
			s:    " a ",
			want: 1,
		},
		{
			s:    "  a",
			want: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := countSegments(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
