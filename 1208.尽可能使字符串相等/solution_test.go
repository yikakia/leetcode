package leetcode

import (
	"reflect"
	"testing"
)

// 1208.尽可能使字符串相等
// https://leetcode-cn.com/problems/get-equal-substrings-within-budget/
func equalSubstring(s string, t string, maxCost int) int {
	if len(s) > len(t) {
		s, t = t, s
	}
	costs := make([]int, len(t))
	for i := range s {
		costs[i] = abs(int(s[i]) - int(t[i]))
	}
	left, right := 0, 0
	res := 0
	curSum := 0
	for right < len(costs) {
		curSum += costs[right]
		for curSum > maxCost {
			curSum -= costs[left]
			left++
		}
		if t := right - left + 1; t > res {
			res = t
		}
		right++
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		s       string
		t       string
		maxCost int
		want    int
	}{
		{
			s:       "abcd",
			t:       "bcdf",
			maxCost: 3,
			want:    3,
		},
		{
			desc:    "",
			s:       "abcd",
			t:       "cdef",
			maxCost: 3,
			want:    1,
		},
		{
			desc:    "",
			s:       "abcd",
			t:       "acde",
			maxCost: 0,
			want:    1,
		},
		{
			desc:    "",
			s:       "abcd",
			t:       "cdef",
			maxCost: 1,
			want:    0,
		},
		{
			s:       "abcd",
			t:       "abcd",
			maxCost: 0,
			want:    4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := equalSubstring(tC.s, tC.t, tC.maxCost)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
