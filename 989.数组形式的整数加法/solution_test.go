package leetcode

import (
	"reflect"
	"testing"
)

// 989.数组形式的整数加法
// https://leetcode-cn.com/problems/add-to-array-form-of-integer/

func addToArrayForm(A []int, K int) []int {
	if K == 0 {
		return A
	}
	ks := []int{}
	for K != 0 {
		ks = append(ks, K%10)
		K /= 10
	}
	reverse(A)
	if len(ks) > len(A) {
		A, ks = ks, A
	}
	A = append(A, 0)
	overflow := 0
	for i := range A {
		if i < len(ks) {
			A[i] += ks[i]
		} else if overflow == 0 {
			break
		}
		A[i] += overflow
		if A[i] > 9 {
			overflow = 1
			A[i] %= 10
		} else {
			overflow = 0
		}
	}
	if A[len(A)-1] == 0 {
		A = A[:len(A)-1]
	}
	reverse(A)
	return A
}

func reverse(a []int) {
	n := len(a)
	end := n / 2
	for i := 0; i < end; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want []int
		A    []int
		K    int
	}{
		{
			desc: "",
			A:    []int{1, 2, 0, 0},
			K:    34,
			want: []int{1, 2, 3, 4},
		},
		{
			A:    []int{2, 7, 4},
			K:    181,
			want: []int{4, 5, 5},
		},
		{
			A:    []int{2, 1, 5},
			K:    806,
			want: []int{1, 0, 2, 1},
		},
		{
			A:    []int{9, 9, 9},
			K:    1,
			want: []int{1, 0, 0, 0},
		},
		{
			A:    []int{9, 0, 9},
			K:    1,
			want: []int{9, 1, 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := addToArrayForm(tC.A, tC.K)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
