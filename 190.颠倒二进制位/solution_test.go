package leetcode

import (
	"reflect"
	"testing"
)

// 190.颠倒二进制位
// https://leetcode-cn.com/problems/reverse-bits

const (
	mask1  uint32 = 0x55555555
	mask2  uint32 = 0x33333333
	mask4  uint32 = 0x0f0f0f0f
	mask8  uint32 = 0x00ff00ff
	mask16 uint32 = 0x0000ffff
)

func reverseBits(num uint32) uint32 {
	num = ((num >> 16) & mask16) | ((num & mask16) << 16)
	num = ((num >> 8) & mask8) | ((num & mask8) << 8)
	num = ((num >> 4) & mask4) | ((num & mask4) << 4)
	num = ((num >> 2) & mask2) | ((num & mask2) << 2)
	num = ((num >> 1) & mask1) | ((num & mask1) << 1)
	return num
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		num  uint32
		want uint32
	}{
		{
			num:  0b00000010100101000001111010011100,
			want: 0b00111001011110000010100101000000,
		},
		{
			num:  0b00000010100101000001111010011100,
			want: 0b00111001011110000010100101000000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := reverseBits(tC.num)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
