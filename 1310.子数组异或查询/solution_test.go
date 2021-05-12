package leetcode

import (
	"reflect"
	"testing"
)

// 1310.子数组异或查询
// https://leetcode-cn.com/problems/xor-queries-of-a-subarray
func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	res := make([]int, len(queries))

	xorStore := make([]int, n+1)
	for i := range xorStore {
		if i == 0 {
			continue
		}
		xorStore[i] = arr[i-1] ^ xorStore[i-1]
	}

	for i, query := range queries {
		l, r := query[0], query[1]
		res[i] = xorStore[r+1]
		if l > 0 {
			res[i] ^= xorStore[l]
		}
	}

	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		arr     []int
		queries [][]int
		want    []int
	}{
		{
			arr:     []int{16},
			queries: [][]int{{0, 0}, {0, 0}},
			want:    []int{16, 16},
		},
		{
			arr:     []int{1, 3, 4, 8},
			queries: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}},
			want:    []int{2, 7, 14, 8},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := xorQueries(tC.arr, tC.queries)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
