package main

import "fmt"

// 5631.跳跃游戏VI
// https://leetcode-cn.com/problems/jump-game-vi/
func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))
	if k > len(nums) {
		k = len(nums)
	}
	dp[0] = nums[0]
	preMax := dp[0]
	for i := 1; i < len(nums); i++ {
		if i <= k {
			dp[i] = preMax + nums[i]
			if dp[i] > preMax {
				preMax = dp[i]
			}
		} else {
			if dp[i-k-1] == preMax {
				preMax = max(dp[i-k : i]...)
			}
			dp[i] = preMax + nums[i]
			if dp[i] > preMax {
				preMax = dp[i]
			}
		}
	}
	return dp[len(nums)-1]
}

func max(nums ...int) (res int) {
	res = nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return
}
func test() {
	type TestType struct {
		nums []int
		k    int
		want int
	}
	ts := []TestType{
		{
			nums: []int{1, -1, -2, 4, -7, 3},
			k:    2,
			want: 7,
		},
		{
			nums: []int{10, -5, -2, 4, 0, 3},
			k:    3,
			want: 17,
		},
		{
			nums: []int{1, -5, -20, 4, -1, 3, -6, -3},
			k:    2,
			want: 0,
		},
	}
	for i, t := range ts {
		get := maxResult(t.nums, t.k)
		if t.want != get {
			// 填充输出格式
			fmt.Printf("i:%d %+v get:%v\n", i, t, get)
		}
	}
	fmt.Println("Finished Test!")

}
func main() {
	test()
}
