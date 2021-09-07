package leetcode

import (
	"reflect"
	"testing"
)

// 1221.分割平衡字符串
// https://leetcode-cn.com/problems/split-a-string-in-balanced-strings
func balancedStringSplit(s string) (ans int) {
	lC, rC := 0, 0
	for _, ch := range s {
		if ch == 'L' {
			lC++
		}
		if ch == 'R' {
			rC++
		}
		if lC == rC {
			ans++
			lC = 0
			rC = 0
		}
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		s    string
		want int
		desc string
	}{
		{
			s:    "RLRRLLRLRL",
			want: 4,
		},
		{
			s:    "RLLLLRRRLR",
			want: 3,
		},
		{
			s:    "LLLLRRRR",
			want: 1,
		},
		{
			s:    "RLRRRLLRLL",
			want: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := balancedStringSplit(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
