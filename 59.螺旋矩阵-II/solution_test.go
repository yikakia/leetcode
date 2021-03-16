package leetcode

import (
	"reflect"
	"testing"
)

// 59.螺旋矩阵II
// https://leetcode-cn.com/problems/spiral-matrix-ii
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	x, y := -1, 0
	dist := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	xcnt, ycnt := 0, 0

	dir := 0
	for cnt := 1; cnt <= n*n; cnt++ {
		x += dist[dir][0]
		y += dist[dir][1]
		dirChange := false
		if x == n-xcnt-1 && dir == 0 {
			xcnt++
			dirChange = true
		} else if y == n-ycnt-1 && dir == 1 {
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
		res[y][x] = cnt
	}

	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		want [][]int
	}{
		{
			n:    1,
			want: [][]int{{1}},
		},
		{
			n:    3,
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := generateMatrix(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
