package leetcode

import (
	"reflect"
	"testing"
)

// 7.整数反转
// https://leetcode-cn.com/problems/reverse-integer
func reverse(x int) int {
	var newX int32 = 0
	for x != 0 {
		tNewX := newX * 10
		if tNewX/10 != newX {
			return 0
		}
		newX = tNewX
		X := x % 10
		tNewX = newX + int32(X)
		if tNewX-int32(X) != newX {
			return 0
		}
		newX = tNewX
		x /= 10
	}
	return int(newX)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		x    int
		want int
	}{
		{
			x:    1534236469,
			want: 0,
		},
		{
			x:    123,
			want: 321,
		},
		{
			x:    120,
			want: 21,
		},
		{
			x:    0,
			want: 0,
		},
		{
			x:    0x7fffffffe,
			want: 0,
		},
		{
			x:    -0x7fffffff,
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := reverse(tC.x)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
