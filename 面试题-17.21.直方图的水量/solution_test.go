package leetcode

import (
	"reflect"
	"testing"
)

// 面试题17.21.直方图的水量
// https://leetcode-cn.com/problems/volume-of-histogram-lcci

func trap(height []int) (ans int) {
	stk := []int{}
	for i, h := range height {
		for len(stk) > 0 && h > height[stk[len(stk)-1]] {
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				break
			}
			left := stk[len(stk)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stk = append(stk, i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		height []int
		want   int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := trap(tC.height)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
