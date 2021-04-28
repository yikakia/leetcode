package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// 633.平方数之和
// https://leetcode-cn.com/problems/sum-of-square-numbers
func judgeSquareSum(c int) bool {

	for a := 0; a*a <= c; a++ {
		t := math.Sqrt(float64(c - a*a))
		if t == math.Floor(t) {
			return true
		}
	}

	return false
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		c    int
		want bool
	}{
		{
			c:    5,
			want: true,
		},
		{
			c:    3,
			want: false,
		},
		{
			c:    4,
			want: true,
		},
		{
			c:    2,
			want: true,
		},
		{
			c:    1,
			want: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := judgeSquareSum(tC.c)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
