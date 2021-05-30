package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// 231.2的幂
// https://leetcode-cn.com/problems/power-of-two
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		want bool
	}{
		{
			n:    math.MinInt64 + 1,
			want: false,
		},
		{
			n:    math.MinInt64,
			want: false,
		},
		{
			n:    -1,
			want: false,
		},
		{
			n:    0,
			want: false,
		},
		{
			n:    1,
			want: true,
		},
		{
			n:    16,
			want: true,
		},
		{
			n:    218,
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := isPowerOfTwo(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
