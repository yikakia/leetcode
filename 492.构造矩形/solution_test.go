package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// 492.构造矩形
// https://leetcode-cn.com/problems/construct-the-rectangle
func constructRectangle(area int) []int {
	for w := int(math.Sqrt(float64(area))); w >= 1; w-- {
		if area%w == 0 {
			return []int{area / w, w}
		}
	}
	return []int{area, 1}
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		area int
		want []int
		desc string
	}{
		{
			area: 4,
			want: []int{2, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := constructRectangle(tC.area)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
