package leetcode

import (
	"reflect"
	"testing"
)

// 54.螺旋矩阵
// https://leetcode-cn.com/problems/spiral-matrix
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	cnt := m * n
	x, y := -1, 0
	dist := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	xcnt, ycnt := 0, 0
	dir := 0
	res := make([]int, 0, cnt)
	for cnt != 0 {
		x += dist[dir][0]
		y += dist[dir][1]
		dirChange := false
		if x == n-xcnt-1 && dir == 0 {
			xcnt++
			dirChange = true
		} else if y == m-ycnt-1 && dir == 1 {
			ycnt++
			dirChange = true
		} else if x == xcnt-1 && dir == 2 {
			dirChange = true
		} else if y == ycnt && dir == 3 {
			dirChange = true
		}
		if dirChange {
			dir++
			dir %= 4
		}
		res = append(res, matrix[y][x])
		cnt--
	}

	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		want   []int
	}{
		{
			matrix: [][]int{{2}},
			want:   []int{2},
		},
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := spiralOrder(tC.matrix)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
