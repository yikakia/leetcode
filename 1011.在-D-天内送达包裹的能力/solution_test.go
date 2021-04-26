package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 1011.在D天内送达包裹的能力
// https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days

func shipWithinDays(weights []int, D int) int {
	// 确定二分查找左右边界
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1 // 需要运送的天数
		sum := 0 // 当前这一天已经运送的包裹重量之和
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		weights []int
		D       int
		want    int
	}{
		{
			weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			D:       5,
			want:    15,
		},
		{
			weights: []int{3, 2, 2, 4, 1, 4},
			D:       3,
			want:    6,
		},
		{
			weights: []int{1, 2, 3, 1, 1},
			D:       4,
			want:    3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := shipWithinDays(tC.weights, tC.D)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
