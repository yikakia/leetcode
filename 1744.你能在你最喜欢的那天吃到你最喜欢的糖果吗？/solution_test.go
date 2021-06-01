package leetcode

import (
	"reflect"
	"testing"
)

// 1744.你能在你最喜欢的那天吃到你最喜欢的糖果吗？
// https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day
func canEat(candiesCount []int, queries [][]int) []bool {
	res := make([]bool, len(queries))
	preCandies := make([]int, len(candiesCount)+1)
	for i, v := range candiesCount {
		preCandies[i+1] = preCandies[i] + v
	}
	for i, query := range queries {
		favoriteType, favoriteDay, dailyCap := query[0], query[1], query[2]
		if dailyCap*(favoriteDay+1) <= preCandies[favoriteType] ||
			favoriteDay+1 > preCandies[favoriteType+1] {
			res[i] = false
			continue
		}
		res[i] = true
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc         string
		candiesCount []int
		queries      [][]int
		want         []bool
	}{
		{
			candiesCount: []int{7, 4, 5, 3, 8},
			queries:      [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 1000000000}},
			want:         []bool{true, false, true},
		},

		{
			candiesCount: []int{5, 2, 6, 4, 1},
			queries:      [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}},
			want:         []bool{false, true, true, false, false},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := canEat(tC.candiesCount, tC.queries)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
