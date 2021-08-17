package leetcode

import (
	"reflect"
	"testing"
)

// 551.学生出勤记录I
// https://leetcode-cn.com/problems/student-attendance-record-i
func checkRecord(s string) bool {
	preL, lCount := -1, 0
	aCount := 0
	for i, ch := range s {
		switch ch {
		case 'A':
			aCount++
			if aCount > 2 {
				return false
			}
		case 'L':
			if i-preL != 1 {
				lCount = 0
			}
			lCount++
			preL = i
			if lCount >= 3 {
				return false
			}
		}
	}

	return true
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
		desc string
	}{
		{
			s:    "PPALLP",
			want: true,
		},
		{
			s:    "PPALLL",
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := checkRecord(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
