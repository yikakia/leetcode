package leetcode

import (
	"reflect"
	"testing"
)

// 1052.爱生气的书店老板
// https://leetcode-cn.com/problems/grumpy-bookstore-owner

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxSatisfied(customers []int, grumpy []int, X int) (res int) {
	n := len(customers)
	for i := range customers {
		if grumpy[i] == 0 {
			res += customers[i]
		}
	}
	var l, r, update, maxupdate int
	for ; r < X; r++ {
		if grumpy[r] == 1 {
			update += customers[r]
		}
	}
	maxupdate = update
	for r < n {
		if grumpy[l] == 1 {
			update -= customers[l]
		}
		l++
		if grumpy[r] == 1 {
			update += customers[r]
		}
		r++
		maxupdate = max(maxupdate, update)
	}
	return res + maxupdate
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc      string
		customers []int
		grumpy    []int
		X         int
		want      int
	}{
		{
			customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
			grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
			X:         3,
			want:      16,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := maxSatisfied(tC.customers, tC.grumpy, tC.X)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
