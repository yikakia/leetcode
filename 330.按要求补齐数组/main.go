package main

import (
	"fmt"
)

// 330.按要求补齐数组
// https://leetcode-cn.com/problems/patching-array/
func minPatches(nums []int, n int) int {
	res := 0
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x *= 2
			res++
		}
	}
	return res
}
func test() {
	type TestType struct {
		nums []int
		n    int
		want int
	}
	ts := []TestType{
		TestType{nums: []int{1, 3}, n: 6, want: 1},
		TestType{nums: []int{1, 5, 10}, n: 20, want: 2},
		TestType{nums: []int{1, 2, 2}, n: 5, want: 0},
	}
	for _, t := range ts {
		get := minPatches(t.nums, t.n)
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
