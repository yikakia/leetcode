package leetcode

import (
	"reflect"
	"testing"
)

// 1004.最大连续1的个数III
// https://leetcode-cn.com/problems/max-consecutive-ones-iii/

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestOnes(A []int, K int) (res int) {
	var left, right, hasZero int
	n := len(A)
	for right < n {
		if A[right] == 0 {
			if hasZero < K {
				hasZero++
			} else {
				for hasZero >= K {
					if A[left] == 0 {
						hasZero--
					}
					left++
				}
				hasZero++
			}
		}
		right++
		res = max(res, right-left)
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		A    []int
		K    int
		want int
	}{
		{
			A:    []int{1, 0, 1, 0, 1, 1, 1, 0, 0, 1},
			K:    2,
			want: 7,
		},
		{
			A:    []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1},
			K:    2,
			want: 7,
		},
		{
			A:    []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			K:    2,
			want: 6,
		},
		{
			A:    []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			K:    3,
			want: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := longestOnes(tC.A, tC.K)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
