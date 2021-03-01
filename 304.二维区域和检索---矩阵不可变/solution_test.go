package leetcode

import (
	"reflect"
	"testing"
)

// 304.二维区域和检索-矩阵不可变
// https://leetcode-cn.com/problems/range-sum-query-2d-immutable
type NumMatrix struct {
	pres [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	if n == 0 {
		return NumMatrix{}
	}
	m := len(matrix[0])
	pres := make([][]int, n)
	for i, row := range matrix {
		pres[i] = make([]int, m+1)
		for j := range row {
			pres[i][j+1] = pres[i][j] + row[j]
		}
	}
	return NumMatrix{pres: pres}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (sum int) {
	for i := row1; i <= row2; i++ {
		sum += this.pres[i][col2+1] - this.pres[i][col1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		ops    [][]int
		want   int
	}{
		{
			desc:   "",
			matrix: [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}},
			ops: [][]int{
				{2, 1, 4, 3, 8},
				{1, 1, 2, 2, 11},
				{1, 2, 2, 4, 12},
				{0, 0, 0, 0, 3},
				{4, 4, 4, 4, 4},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			n := Constructor(tC.matrix)
			for i, op := range tC.ops {
				get := n.SumRegion(op[0], op[1], op[2], op[3])
				if !reflect.DeepEqual(op[4], get) {
					t.Errorf("%d input: %+v get: %v\n", i, tC, get)
				}
			}

		})
	}
}
