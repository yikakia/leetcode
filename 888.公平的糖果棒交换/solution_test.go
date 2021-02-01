package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 888.公平的糖果棒交换
// https://leetcode-cn.com/problems/fair-candy-swap/

func sum(nums ...int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}

func fairCandySwap(A []int, B []int) []int {
	sort.Ints(B)
	sa, sb := sum(A...), sum(B...)
	delta := sb - sa
	for _, num := range A {
		want := num + delta/2
		index := sort.Search(len(B), func(i int) bool { return B[i] >= want })
		if index < len(B) && B[index] == want {
			return []int{num, want}
		}
	}
	return []int{}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc       string
		A, B, want []int
	}{
		{
			A:    []int{1, 1},
			B:    []int{2, 2},
			want: []int{1, 2},
		},
		{
			A:    []int{1, 2},
			B:    []int{2, 3},
			want: []int{1, 2},
		},
		{
			A:    []int{2},
			B:    []int{1, 3},
			want: []int{2, 3},
		},
		{
			A:    []int{1, 2, 5},
			B:    []int{2, 4},
			want: []int{5, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := fairCandySwap(tC.A, tC.B)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
