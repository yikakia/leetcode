package leetcode

import (
	"reflect"
	"testing"
)

// 978.最长湍流子数组
// https://leetcode-cn.com/problems/longest-turbulent-subarray/
func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	switch {
	case n < 2:
		return 1
	case n == 2:
		if arr[0] == arr[1] {
			return 1
		}
		return 2
	}
	compare := make([]int, n-1)
	for i := range arr[1:] {
		switch {
		case arr[i] > arr[i+1]:
			compare[i] = -1
		case arr[i] < arr[i+1]:
			compare[i] = 1
		case arr[i] == arr[i+1]:
			compare[i] = 0
		}
	}
	st, end := 0, 1
	res := 1
	for end < n-1 {
		if t := compare[end-1] * compare[end]; t == -1 {
			if t := end - st + 2; t > res {
				res = t
			}
			end++
		} else if t == 0 {
			if !(compare[st] == 0 && compare[end] == 0) {
				if res < 2 {
					res = 2
				}
			}
			end++
			st = end - 1
		} else if t == 1 {
			end++
			st = end - 1
			if res < 2 {
				res = 2
			}
		} else {
			panic("unreachable")
		}
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		arr  []int
		want int
	}{
		{
			arr:  []int{8, 8, 8, 1},
			want: 2,
		},
		{
			arr:  []int{8, 8, 8, 9},
			want: 2,
		},
		{
			arr:  []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			want: 5,
		},
		{
			arr:  []int{4, 8, 12, 16},
			want: 2,
		},
		{
			arr:  []int{100},
			want: 1,
		},
		{
			arr:  []int{4, 3, 2, 1},
			want: 2,
		},

		{
			arr:  []int{8, 8, 9, 8},
			want: 3,
		},
		{
			arr:  []int{9, 8, 8, 9},
			want: 2,
		},
		{
			arr:  []int{9, 8, 9},
			want: 3,
		},
		{
			arr:  []int{9, 9, 9},
			want: 1,
		},
		{
			arr:  []int{9, 9},
			want: 1,
		},
		{
			arr:  []int{9, 8, 8},
			want: 2,
		},
		{
			arr:  []int{9, 9},
			want: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := maxTurbulenceSize(tC.arr)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
