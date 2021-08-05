package main

import (
	"reflect"
	"testing"
)

// 171.Excel表列序号
// https://leetcode-cn.com/problems/excel-sheet-column-number
func titleToNumber(columnTitle string) int {
	res := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res *= 26
		res += int(columnTitle[i] - 'A' + 1)
	}

	return 0
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		columnTitle string
		want        int
		desc        string
	}{
		{
			columnTitle: "A",
			want:        1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := titleToNumber(tC.columnTitle)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
