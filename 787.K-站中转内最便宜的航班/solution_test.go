package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// 787.K站中转内最便宜的航班
// https://leetcode-cn.com/problems/cheapest-flights-within-k-stops
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([][]int, n)

	for i := range cost {
		cost[i] = make([]int, k+2)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt64 / 2
		}
	}
	cost[src][0] = 0
	for i := 1; i <= k+1; i++ {
		for _, flight := range flights {
			to := flight[1]
			from := flight[0]
			cost[to][i] = min(cost[to][i], cost[from][i-1]+flight[2])
		}
	}
	res := math.MaxInt64 >> 1
	for _, d := range cost[dst][1:] {
		res = min(res, d)
	}
	if res == math.MaxInt64/2 {
		return -1
	}

	return res
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if num < res {
			res = num
		}
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		want    int
		desc    string
	}{
		{
			n: 3,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			src:  0,
			dst:  2,
			k:    1,
			want: 200,
		}, {
			n: 3,
			flights: [][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			src:  0,
			dst:  2,
			k:    0,
			want: 500,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findCheapestPrice(tC.n, tC.flights, tC.src, tC.dst, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
