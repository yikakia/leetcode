package leetcode

import (
	"reflect"
	"testing"
)

// 1486.数组异或操作
// https://leetcode-cn.com/problems/xor-operation-in-an-array
func xorOperation(n int, start int) int {
	nums := start
	for i := 1; i < n; i++ {
		start += 2
		nums ^= start
	}

	return nums
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		n     int
		start int
		want  int
	}{
		{
			n:     5,
			start: 0,
			want:  8,
		},
		{
			n:     4,
			start: 3,
			want:  8,
		},
		{
			n:     1,
			start: 7,
			want:  7,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := xorOperation(tC.n, tC.start)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
