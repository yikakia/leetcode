package leetcode

import (
	"reflect"
	"testing"
)

// 剑指Offer10-I.斐波那契数列
// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof

const (
	mod = 1e9 + 7
)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b, c := 0, 1, 0
	for i := 0; i < n-1; i++ {
		c = (a + b) % mod
		a = b
		b = c
	}
	return c
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		n    int
		want int
		desc string
	}{
		{
			n:    2,
			want: 1,
		},
		{
			n:    3,
			want: 2,
		},
		{
			n:    4,
			want: 3,
		},
		{
			n:    5,
			want: 5,
		},
		{
			n:    6,
			want: 8,
		},
		{
			n:    7,
			want: 13,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := fib(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
