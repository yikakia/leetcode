package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 228.汇总区间
// https://leetcode-cn.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return []string{}
	}
	ranges := make([][]int, 0)
	ranges = append(ranges, []int{})
	ranges[0] = append(ranges[0], nums[0])

	for i := 1; i < n; i++ {
		num := nums[i]
		m1 := len(ranges) - 1
		m2 := len(ranges[m1]) - 1
		if num-ranges[m1][m2] == 1 {
			if m2 == 0 {
				ranges[m1] = append(ranges[m1], num)
			} else {
				ranges[m1][m2] = num
			}
			continue
		}
		ranges = append(ranges, []int{num})
	}

	res := []string{}
	for i := range ranges {
		if len(ranges[i]) == 2 {
			res = append(res, strconv.Itoa(ranges[i][0])+"->"+strconv.Itoa(ranges[i][1]))
		} else {
			res = append(res, strconv.Itoa(ranges[i][0]))
		}
	}
	return res
}
func test() {
	type TestType struct {
		nums []int
		want []string
	}
	ts := []TestType{
		TestType{
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		TestType{
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		TestType{
			nums: []int{},
			want: []string{},
		},
		TestType{
			nums: []int{-1},
			want: []string{"-1"},
		},
	}
	for _, t := range ts {
		get := summaryRanges(t.nums)
		if !reflect.DeepEqual(get, t.want) {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")
}
func main() {
	test()
}
