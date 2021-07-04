package leetcode

import (
	"reflect"
	"testing"
)

// 1833.雪糕的最大数量
// https://leetcode-cn.com/problems/maximum-ice-cream-bars
func maxIceCream(costs []int, coins int) int {
	rCos := make([]int, coins+1)
	for _, cost := range costs {
		for i := coins; i >= cost; i-- {
			rCos[i] = max(rCos[i-cost]+1, rCos[i])
		}
	}
	return rCos[coins]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		costs []int
		coins int
		want  int
		desc  string
	}{
		{
			costs: []int{1, 3, 2, 4, 1},
			coins: 7,
			want:  4,
		},
		{
			costs: []int{10, 6, 8, 7, 7, 8},
			coins: 5,
			want:  0,
		},
		{
			costs: []int{1, 6, 3, 1, 2, 5},
			coins: 20,
			want:  6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := maxIceCream(tC.costs, tC.coins)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
