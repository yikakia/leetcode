package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 1738.找出第K大的异或坐标值
// https://leetcode-cn.com/problems/find-kth-largest-xor-coordinate-value
func kthLargestValue(matrix [][]int, k int) (res int) {
	m, n := len(matrix), len(matrix[0])
	preXor := make([][]int, m+1)
	preXor[0] = make([]int, n+1)
	for i := range preXor[1:] {
		preXor[i+1] = make([]int, n+1)
		for j := range preXor[i+1][1:] {
			preXor[i+1][j+1] = preXor[i+1][j] ^ matrix[i][j]
		}
	}
	q := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			preXor[i+1][j+1] ^= preXor[i][j+1]
			q[i*n+j] ^= preXor[i+1][j+1]
		}
	}

	sort.Ints(q)
	return q[m*n-k]
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		k      int
		want   int
	}{
		{
			matrix: [][]int{
				{5, 2},
				{1, 6},
			},
			k:    1,
			want: 7,
		},
		{
			matrix: [][]int{
				{5, 2},
				{1, 6},
			},
			k:    2,
			want: 5,
		},
		{
			matrix: [][]int{
				{5, 2},
				{1, 6},
			},
			k:    3,
			want: 4,
		},
		{
			matrix: [][]int{
				{5, 2},
				{1, 6},
			},
			k:    4,
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := kthLargestValue(tC.matrix, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
