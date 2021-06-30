package leetcode

import (
	"reflect"
	"testing"
)

// 168.Excel表列名称
// https://leetcode-cn.com/problems/excel-sheet-column-title
func convertToTitle(columnNumber int) (res string) {
	for columnNumber > 0 {
		a0 := (columnNumber - 1) % 26
		res = string('A'+byte(a0)) + res
		columnNumber = (columnNumber - a0) / 26
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		columnNumber int
		want         string
		desc         string
	}{
		{
			columnNumber: 1,
			want:         "A",
		},
		{
			columnNumber: 28,
			want:         "AB",
		},
		{
			columnNumber: 701,
			want:         "ZY",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := convertToTitle(tC.columnNumber)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
