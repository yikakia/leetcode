package leetcode

import (
	"reflect"
	"testing"
)

// 832.翻转图像
// https://leetcode-cn.com/problems/flipping-an-image
func flipAndInvertImage(A [][]int) [][]int {
	m := len(A[0])
	for i := range A {
		for j := 0; j < (m+1)/2; j++ {
			A[i][j], A[i][m-j-1] = 1^A[i][m-j-1], 1^A[i][j]
		}
	}
	return A
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		A    [][]int
		want [][]int
	}{
		{
			A:    [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			want: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			A:    [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			want: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := flipAndInvertImage(tC.A)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
