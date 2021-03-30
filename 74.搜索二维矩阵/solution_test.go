package leetcode

import (
	"reflect"
	"testing"
)

// 74.搜索二维矩阵
// https://leetcode-cn.com/problems/search-a-2d-matrix
func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	start, end := 0, n*m-1
	mid := (start + end) / 2

	for start <= end {
		x, y := mid/m, mid%m
		// if x >= n || y >= m {
		// 	return false
		// }
		if matrix[x][y] < target {
			start = mid + 1
		} else if matrix[x][y] > target {
			end = mid - 1
		} else {
			return true
		}
		mid = (start + end) / 2
	}

	return false
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			matrix: [][]int{{1}},
			target: 33,
			want:   false,
		},
		{
			matrix: [][]int{{1, 2, 3, 4, 5, 6, 7, 8}},
			target: 33,
			want:   false,
		},
		{
			matrix: [][]int{{1, 2, 3, 4, 5, 6, 7, 8}},
			target: 4,
			want:   true,
		},
		{
			matrix: [][]int{{1}, {2}, {3}, {4}},
			target: 50,
			want:   false,
		},
		{
			matrix: [][]int{{1}, {2}, {3}, {4}},
			target: 4,
			want:   true,
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
