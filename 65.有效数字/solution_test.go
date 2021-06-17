package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

// 65.有效数字
// https://leetcode-cn.com/problems/valid-number

func isNum(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	if len(s) == 0 {
		return false
	}
	return isNum(s)
}

func isDecimal(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	index := strings.Index(s, ".")
	if index == -1 {
		return false
	}
	if index == 0 {
		return isNum(s[1:])
	}
	if index == len(s)-1 {
		return isInt(s[:index])
	}

	return isInt(s[:index]) && isNum(s[index+1:])

}

func isNumber(s string) bool {
	s = strings.ReplaceAll(s, "E", "e")
	index := strings.Index(s, "e")
	if index != -1 {
		return (isDecimal(s[:index]) || isInt(s[:index])) &&
			isInt(s[index+1:])
	}

	return isInt(s) || isDecimal(s)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
		desc string
	}{
		{
			s:    "0..",
			want: false,
		},
		{
			s:    ".-4",
			want: false,
		},
		{
			s:    "+-.",
			want: false,
		},

		{
			s:    "..1",
			want: false,
		},
		{
			s:    ".1.1",
			want: false,
		},

		{
			s:    "1..1",
			want: false,
		},
		{
			s:    "10.2",
			want: true,
		},
		{
			s:    "0",
			want: true,
		},
		{
			s:    "e",
			want: false,
		},
		{
			s:    ".",
			want: false,
		},
		{
			s:    ".1",
			want: true,
		},
		{
			s:    "0089",
			want: true,
		},
		{
			s:    "-90E3",
			want: true,
		},
		{
			s:    "+6e-1",
			want: true,
		},
		{
			s:    "53.5e93",
			want: true,
		},
		{
			s:    "99e2.5",
			want: false,
		},
		{
			s:    "--6",
			want: false,
		},
		{
			s:    "-+3",
			want: false,
		},
		{
			s:    "95a54e53",
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := isNumber(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
