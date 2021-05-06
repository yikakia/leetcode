package leetcode

import (
	"reflect"
	"testing"
)

// 1720.解码异或后的数组
// https://leetcode-cn.com/problems/decode-xored-array
func decode(encoded []int, first int) []int {
	n := len(encoded)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n+1)
	res[0] = first
	for i := range encoded {
		res[i+1] = res[i] ^ encoded[i]
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		encoded []int
		first   int
		want    []int
	}{
		{
			encoded: []int{1, 2, 3},
			first:   1,
			want:    []int{1, 0, 2, 1},
		},
		{
			encoded: []int{6, 2, 7, 3},
			first:   4,
			want:    []int{4, 2, 0, 7, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := decode(tC.encoded, tC.first)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
