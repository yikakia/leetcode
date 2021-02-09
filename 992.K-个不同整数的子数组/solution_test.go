package leetcode

import (
	"reflect"
	"testing"
)

// 992.K个不同整数的子数组
// https://leetcode-cn.com/problems/subarrays-with-k-different-integers/
func subarraysWithKDistinct(A []int, K int) int {
	n := len(A)
	num1 := make([]int, n+1)
	num2 := make([]int, n+1)
	res := 0
	cnt1, cnt2 := 0, 0
	left1, left2 := 0, 0
	for _, v := range A {
		if num1[v] == 0 {
			cnt1++
		}
		num1[v]++
		if num2[v] == 0 {
			cnt2++
		}
		num2[v]++
		for cnt1 > K {
			num1[A[left1]]--
			if num1[A[left1]] == 0 {
				cnt1--
			}
			left1++
		}
		for cnt2 > K-1 {
			num2[A[left2]]--
			if num2[A[left2]] == 0 {
				cnt2--
			}
			left2++
		}
		res += left2 - left1
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		A    []int
		K    int
		want int
	}{
		{
			A:    []int{1, 2, 1, 2, 1, 2, 3, 3},
			K:    2,
			want: 9,
		},
		{
			A:    []int{1, 2, 1, 2, 3},
			K:    2,
			want: 7,
		},
		{
			A:    []int{1, 2, 1, 3, 4},
			K:    3,
			want: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := subarraysWithKDistinct(tC.A, tC.K)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
