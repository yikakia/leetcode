package leetcode

import (
	"reflect"
	"testing"
)

// 566.重塑矩阵
// https://leetcode-cn.com/problems/reshape-the-matrix/
func matrixReshape(nums [][]int, r int, c int) [][]int {
	rr := len(nums)
	cc := len(nums[0])
	if rr*cc != r*c {
		return nums
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	index := 0
	for i := 0; i < rr; i++ {
		for j := 0; j < cc; j++ {
			newI := index / c
			newJ := index % c
			index++
			res[newI][newJ] = nums[i][j]
		}
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums [][]int
		r    int
		c    int
		want [][]int
	}{
		{
			nums: [][]int{{1, 2, 3, 4, 5}},
			r:    5,
			c:    1,
			want: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			nums: [][]int{{1, 2}, {3, 4}},
			r:    1,
			c:    4,
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			nums: [][]int{{1, 2}, {3, 4}},
			r:    2,
			c:    4,
			want: [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := matrixReshape(tC.nums, tC.r, tC.c)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
