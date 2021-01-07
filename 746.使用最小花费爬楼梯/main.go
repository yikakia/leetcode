package main

import "fmt"

// 746.使用最小花费爬楼梯
// https://leetcode-cn.com/problems/min-cost-climbing-stairs/

func minCostClimbingStairs(cost []int) int {
	cur, prev1, prev2 := 0, 0, 0
	for _, costi := range cost {
		cur = min(prev1, prev2) + costi
		prev1, prev2 = prev2, cur
	}
	return min(prev1, prev2)
}
func min(nums ...int) (res int) {
	res = nums[0]
	for _, num := range nums {
		if res > num {
			res = num
		}
	}
	return
}
func test() {
	type TestType struct {
		cost []int
		want int
	}
	ts := []TestType{
		TestType{
			cost: []int{10, 15, 20},
			want: 15,
		},
		TestType{
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}
	for _, t := range ts {
		get := minCostClimbingStairs(t.cost)
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
