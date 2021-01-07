package main

import (
	"fmt"
	"sort"
)

// 435.无重叠区间
// https://leetcode-cn.com/problems/non-overlapping-intervals/

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	res := 1
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	maxRight := intervals[0][1]
	for _, interval := range intervals {
		if interval[0] >= maxRight {
			maxRight = interval[1]
			res++
		}
	}

	return len(intervals) - res
}
func test() {
	type TestType struct {
		intervals [][]int
		want      int
	}
	ts := []TestType{
		TestType{
			intervals: [][]int{
				[]int{1, 2},
				[]int{2, 3},
				[]int{3, 4},
				[]int{1, 3},
			},
			want: 1,
		},
		TestType{
			intervals: [][]int{
				[]int{1, 2},
				[]int{1, 2},
				[]int{1, 2},
			},
			want: 2,
		},
		TestType{
			intervals: [][]int{
				[]int{1, 2},
				[]int{2, 3},
			},
			want: 0,
		},
	}
	for _, t := range ts {
		get := eraseOverlapIntervals(t.intervals)
		if get != t.want {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")

}
func main() {
	test()
}
