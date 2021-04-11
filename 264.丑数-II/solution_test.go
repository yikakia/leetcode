package leetcode

import (
	"reflect"
	"testing"
)

// 264.丑数II
// https://leetcode-cn.com/problems/ugly-number-ii

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	uglyNums := make([]int, n)
	uglyNums[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		x2, x3, x5 := uglyNums[p2]*2, uglyNums[p3]*3, uglyNums[p5]*5
		uglyNums[i] = min(x2, x3, x5)
		if uglyNums[i] == x2 {
			p2++
		}
		if uglyNums[i] == x3 {
			p3++
		}
		if uglyNums[i] == x5 {
			p5++
		}
	}
	return uglyNums[n-1]
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if res > num {
			res = num
		}
	}
	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		want int
	}{
		{
			n:    10,
			want: 12,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := nthUglyNumber(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
