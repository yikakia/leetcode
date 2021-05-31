package leetcode

import (
	"reflect"
	"testing"
)

// 342.4的幂
// https://leetcode-cn.com/problems/power-of-four
func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		want bool
	}{
		{
			n:    4,
			want: true,
		},
		{
			n:    16,
			want: true,
		},
		{
			n:    5,
			want: false,
		},
		{
			n:    1,
			want: true,
		},
		{
			n:    2,
			want: false,
		},
		{
			n:    3,
			want: false,
		},
		{
			n:    8,
			want: false,
		},
		{
			n:    7,
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := isPowerOfFour(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
