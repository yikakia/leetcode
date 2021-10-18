package leetcode

import (
	"reflect"
	"testing"
)

// 476.数字的补数
// https://leetcode-cn.com/problems/number-complement
func findComplement(num int) int {
	highBit := 0
	for i := 1; i <= 30; i++ {
		if num < 1<<i {
			break
		}
		highBit = i
	}
	mask := 1<<(highBit+1) - 1
	return num ^ mask
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		num  int
		want int
		desc string
	}{
		{
			num:  5,
			want: 2,
		},
		{
			num:  1,
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findComplement(tC.num)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
