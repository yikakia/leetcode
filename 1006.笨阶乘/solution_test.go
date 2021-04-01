package leetcode

import (
	"reflect"
	"testing"
)

// 1006.笨阶乘
// https://leetcode-cn.com/problems/clumsy-factorial
func clumsy(N int) (ans int) {
	switch {
	case N == 1:
		return 1
	case N == 2:
		return 2
	case N == 3:
		return 6
	case N == 4:
		return 7

	case N%4 == 0:
		return N + 1
	case N%4 <= 2:
		return N + 2
	default:
		return N - 1
	}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		N    int
		want int
	}{
		{
			N:    1,
			want: 1,
		},
		{
			N:    4,
			want: 7,
		},
		{
			N:    10,
			want: 12,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := clumsy(tC.N)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
