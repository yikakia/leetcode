package leetcode

import (
	"reflect"
	"testing"
)

// 240.搜索二维矩阵II
// https://leetcode-cn.com/problems/search-a-2d-matrix-ii
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := n-1, 0

	for x >= 0 && y < m {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			x--
		}
		if matrix[x][y] < target {
			y++
		}
	}
	return false
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		target int
		want   bool
		desc   string
	}{
		{
			matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target: 5,
			want:   true,
		},
		{
			matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target: 20,
			want:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := searchMatrix(tC.matrix, tC.target)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
