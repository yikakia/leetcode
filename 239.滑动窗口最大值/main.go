package main

import (
	"fmt"
	"math"
)

// 239.滑动窗口最大值
// https://leetcode-cn.com/problems/sliding-window-maximum/

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n-k+1)
	queue := make([]int, 1)
	queue[0] = math.MinInt32
	for i := 0; i < k; i++ {
		queue = adjQueue(queue, nums[i])
	}
	for i := k; i < n; i++ {
		res[i-k] = queue[0]
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		queue = adjQueue(queue, nums[i])
	}
	res[n-k] = queue[0]
	return res
}

func adjQueue(queue []int, x int) []int {
	n := len(queue)
	res := make([]int, 0)
	if n == 0 {
		res = append(res, x)
	}
	for i := n - 1; i >= 0; i-- {
		if x <= queue[i] {
			res = append(queue[:i+1], x)
			break
		} else {
			res = append(queue[:i], x)
		}
	}

	return res
}
func isSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func test() {
	type TestType struct {
		nums []int
		k    int
		want []int
	}
	ts := []TestType{
		TestType{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7}},
		TestType{
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1}},
		TestType{
			nums: []int{4, -2},
			k:    2,
			want: []int{4}},
	}
	for _, t := range ts {
		get := maxSlidingWindow(t.nums, t.k)
		if !isSame(get, t.want) {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")
}
func main() {
	test()
}
