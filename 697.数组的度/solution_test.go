package leetcode

import (
	"reflect"
	"testing"
)

// 697.数组的度
// https://leetcode-cn.com/problems/degree-of-an-array/

func findShortestSubArray(nums []int) (res int) {
	type entry struct {
		frequency, l, dis int
	}
	N := 0
	res = len(nums)
	cntMap := map[int]entry{}
	for i, num := range nums {
		e := cntMap[num]
		if e.frequency == 0 {
			e = entry{frequency: 1, l: i, dis: 1}
		} else {
			e.frequency++
			e.dis = i - e.l + 1
		}
		cntMap[num] = e
		if N < e.frequency {
			N = e.frequency
			res = e.dis
		} else if N == e.frequency && res > e.dis {
			res = e.dis
		}
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 2, 3, 1},
			want: 2,
		},
		{
			nums: []int{1, 2, 2, 3, 1, 4, 2},
			want: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findShortestSubArray(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
