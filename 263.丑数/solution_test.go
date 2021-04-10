package leetcode

import (
	"reflect"
	"testing"
)

// 263.丑数
// https://leetcode-cn.com/problems/ugly-number
func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		want bool
	}{
		{
			n:    6,
			want: true,
		},
		{
			n:    8,
			want: true,
		},
		{
			n:    14,
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := isUgly(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
