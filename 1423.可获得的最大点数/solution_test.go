package leetcode

import (
	"reflect"
	"testing"
)

// 1423.可获得的最大点数
// https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards/

func sum(nums ...int) (res int) {
	for _, t := range nums {
		res += t
	}
	return
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxScore(cardPoints []int, k int) (res int) {
	if k == len(cardPoints) {
		return sum(cardPoints...)
	}
	n := len(cardPoints)

	window := sum(cardPoints[:n-k]...)
	minSum := window
	for i := range cardPoints[n-k:] {
		window += cardPoints[i+n-k] - cardPoints[i]
		minSum = min(minSum, window)
	}

	return sum(cardPoints[:k]...) + window - minSum
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc       string
		cardPoints []int
		k          int
		want       int
	}{
		{
			cardPoints: []int{1, 2, 3, 4, 5, 6, 1},
			k:          3,
			want:       12,
		},
		{
			cardPoints: []int{2, 2, 2},
			k:          2,
			want:       4,
		},
		{
			cardPoints: []int{9, 7, 7, 9, 7, 7, 9},
			k:          7,
			want:       55,
		},
		{
			cardPoints: []int{1, 1000, 1},
			k:          1,
			want:       1,
		},
		{
			cardPoints: []int{1, 79, 80, 1, 1, 1, 200, 1},
			k:          3,
			want:       202,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := maxScore(tC.cardPoints, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
