package leetcode

import (
	"reflect"
	"testing"
)

// 995.K连续位的最小翻转次数
// https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/

func minKBitFlips(A []int, K int) (ans int) {
	n := len(A)
	diff := make([]int, n+1)
	revCnt := 0
	for i, v := range A {
		revCnt ^= diff[i]
		if v == revCnt { // v^revCnt == 0
			if i+K > n {
				return -1
			}
			ans++
			revCnt ^= 1
			diff[i+K] ^= 1
		}

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
			A:    []int{0, 1, 0},
			K:    1,
			want: 2,
		},
		{
			A:    []int{1, 0, 0, 0, 0},
			K:    2,
			want: 2,
		},
		{
			A:    []int{0, 0, 0, 1, 0, 1, 1, 0},
			K:    3,
			want: 3,
		},

		{
			A:    []int{1, 1, 0},
			K:    2,
			want: -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := minKBitFlips(tC.A, tC.K)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
