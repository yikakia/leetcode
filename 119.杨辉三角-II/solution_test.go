package leetcode

import (
	"reflect"
	"testing"
)

// 119.杨辉三角II
// https://leetcode-cn.com/problems/pascals-triangle-ii/
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		rowIndex int
		want     []int
	}{
		{
			rowIndex: 3,
			want:     []int{1, 3, 3, 1},
		},
		{
			rowIndex: 4,
			want:     []int{1, 4, 6, 4, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := getRow(tC.rowIndex)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
