package leetcode

import (
	"reflect"
	"testing"
)

// 766.托普利茨矩阵
// https://leetcode-cn.com/problems/toeplitz-matrix/
func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		want   bool
	}{
		{
			matrix: [][]int{{1}},
			want:   true,
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}},
			want:   true,
		},
		{
			matrix: [][]int{{1, 2}, {2, 2}},
			want:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := isToeplitzMatrix(tC.matrix)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
