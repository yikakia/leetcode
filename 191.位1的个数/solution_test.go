package leetcode

import (
	"reflect"
	"testing"
)

// 191.位1的个数
// https://leetcode-cn.com/problems/number-of-1-bits
func hammingWeight(num uint32) (res int) {
	for num != 0 {
		res++
		num &= num - 1
	}

	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		num  uint32
		want int
	}{
		{
			num:  00000000000000000000000000001011,
			want: 3,
		},
		{
			num:  00000000000000000000000010000000,
			want: 1,
		},
		{
			num:  0b11111111111111111111111111111101,
			want: 31,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := hammingWeight(tC.num)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
