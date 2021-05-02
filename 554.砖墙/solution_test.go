package leetcode

import (
	"reflect"
	"testing"
)

// 554.砖墙
// https://leetcode-cn.com/problems/brick-wall
func leastBricks(wall [][]int) (res int) {
	brickStart := map[int]int{}

	for _, row := range wall {
		pre := 0
		for _, brick := range row {
			brickStart[pre]++
			pre += brick
		}
	}

	n := len(wall)
	res = n
	for k, v := range brickStart {
		if k > 0 {
			res = min(res, n-v)
		}
	}

	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		wall [][]int
		want int
	}{
		{
			wall: [][]int{
				{1, 2, 2, 1},
				{3, 1, 2},
				{1, 3, 2},
				{2, 4},
				{3, 1, 2},
				{1, 3, 1, 1},
			},
			want: 2,
		},
		{
			wall: [][]int{
				{1},
				{1},
				{1},
			},
			want: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := leastBricks(tC.wall)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
