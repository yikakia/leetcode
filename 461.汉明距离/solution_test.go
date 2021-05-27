package leetcode

import (
	"math/bits"
	"reflect"
	"testing"
)

// 461.汉明距离
// https://leetcode-cn.com/problems/hamming-distance
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		x, y int
		want int
	}{
		{
			x:    1,
			y:    4,
			want: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := hammingDistance(tC.x, tC.y)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
