package leetcode

import (
	"reflect"
	"testing"
)

// 73.矩阵置零
// https://leetcode-cn.com/problems/set-matrix-zeroes
func setZeroes(matrix [][]int) {
	firstContain := false
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			firstContain = true
		}
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if firstContain {
			matrix[i][0] = 0
		}
	}
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		want   [][]int
	}{
		{
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			setZeroes(tC.matrix)
			if !reflect.DeepEqual(tC.want, tC.matrix) {
				t.Errorf("input: %+v get: %v\n", tC, tC.matrix)
			}
		})
	}
}
