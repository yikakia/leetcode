package leetcode

import (
	"reflect"
	"testing"
)

// 867.转置矩阵
// https://leetcode-cn.com/problems/transpose-matrix
func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := range matrix {
		for j := range matrix[i] {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		want   [][]int
	}{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := transpose(tC.matrix)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
