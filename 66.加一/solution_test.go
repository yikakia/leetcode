package leetcode

import (
	"reflect"
	"testing"
)

// 66.加一
// https://leetcode-cn.com/problems/plus-one/
func plusOne(digits []int) []int {
	n := len(digits)
	overflow := 0
	digits[n-1]++
	for i := n - 1; i >= 0; i-- {
		digits[i] += overflow
		if digits[i] > 9 {
			overflow = 1
			digits[i] %= 10
		} else {
			overflow = 0
			break
		}
	}
	if overflow > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		want   []int
		digits []int
	}{
		{
			digits: []int{9},
			want:   []int{1, 0},
		},
		{
			desc:   "",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := plusOne(tC.digits)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
