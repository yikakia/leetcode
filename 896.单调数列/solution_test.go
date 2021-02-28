package leetcode

import (
	"reflect"
	"testing"
)

// 896.单调数列
// https://leetcode-cn.com/problems/monotonic-array

const (
	Inc = iota + 1
	Dec
)

func isMonotonic(A []int) bool {
	n := len(A)
	if n < 3 {
		return true
	}
	flag := 0
	for i := 1; i < n; i++ {
		switch {
		case A[i] > A[i-1]:
			flag |= Inc
		case A[i] < A[i-1]:
			flag |= Dec
		default:
		}
		if flag == Inc|Dec {
			return false
		}
	}
	return true
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		A    []int
		want bool
	}{
		{
			A:    []int{1, 2, 2, 3},
			want: true,
		},
		{
			A:    []int{6, 5, 4, 4},
			want: true,
		},
		{
			A:    []int{1, 3, 2},
			want: false,
		},
		{
			A:    []int{1, 2, 4, 5},
			want: true,
		},
		{
			A:    []int{1, 1, 1},
			want: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := isMonotonic(tC.A)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
