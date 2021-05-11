package leetcode

import (
	"reflect"
	"testing"
)

// 1734.解码异或后的排列
// https://leetcode-cn.com/problems/decode-xored-permutation
func decode(encoded []int) []int {
	n := len(encoded)

	totalXor := 0

	for i := 1; i <= n+1; i++ {
		totalXor ^= i
	}
	mask := 0
	for i := 1; i < n; i += 2 {
		mask ^= encoded[i]
	}
	res := make([]int, n+1)
	res[0] = totalXor ^ mask

	for i := 1; i < n+1; i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}

	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		encoded []int
		want    []int
	}{
		{
			encoded: []int{3, 1},
			want:    []int{1, 2, 3},
		},
		{
			encoded: []int{6, 5, 4, 6},
			want:    []int{2, 4, 1, 5, 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := decode(tC.encoded)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
