package leetcode

import (
	"math/bits"
	"reflect"
	"testing"
)

// 剑指Offer15.二进制中1的个数
// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof
func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		num  uint32
		want int
		desc string
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
