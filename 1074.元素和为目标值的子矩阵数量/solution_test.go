package leetcode

import (
	"reflect"
	"testing"
)

// 1074.元素和为目标值的子矩阵数量
// https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target
func numSubmatrixSumTarget(matrix [][]int, target int) (ret int) {
	n, m := len(matrix), len(matrix[0])
	preSum := make([][]int, n+1)
	for i := range preSum {
		preSum[i] = make([]int, m+1)
		if i != 0 {
			for j := range preSum[i] {
				if j == 0 {
					continue
				}
				preSum[i][j] = matrix[i-1][j-1] +
					preSum[i][j-1] +
					preSum[i-1][j] - preSum[i-1][j-1]
			}
		}
	}

	for i := range preSum {
		for j := range preSum[i] {
			for up := i - 1; up >= 0; up-- {
				for left := j - 1; left >= 0; left-- {
					area := preSum[i][j]
					area += preSum[up][left]
					area -= preSum[i][left]
					area -= preSum[up][j]
					if area == target {
						ret++
					}
				}
			}
		}
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		target int
		want   int
	}{
		{
			matrix: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
			target: 0,
			want:   4,
		},
		{
			matrix: [][]int{{1, -1}, {-1, 1}},
			target: 0,
			want:   5,
		},
		{
			matrix: [][]int{{904}},
			target: 0,
			want:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numSubmatrixSumTarget(tC.matrix, tC.target)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
