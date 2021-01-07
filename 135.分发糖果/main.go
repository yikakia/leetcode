package main

import "fmt"

// 135.分发糖果
// https://leetcode-cn.com/problems/candy/

func candy(ratings []int) int {
	n := len(ratings)
	count := make([]int, n)
	for i := range count {
		count[i] = 1
	}
	changed := true
	for changed {
		changed = false
		for i := 1; i < n; i++ {
			if ratings[i] > ratings[i-1] && count[i] <= count[i-1] {
				count[i] = count[i-1] + 1
				changed = true
			}
		}
		for i := n - 2; i >= 0; i-- {
			if ratings[i] > ratings[i+1] && count[i] <= count[i+1] {
				count[i] = count[i+1] + 1
				changed = true
			}
		}
	}

	return sum(count...)
}
func sum(nums ...int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}
func test() {
	type TestType struct {
		ratings []int
		want    int
	}
	ts := []TestType{
		TestType{
			ratings: []int{1, 0, 2},
			want:    5,
		},
		TestType{
			ratings: []int{1, 2, 2},
			want:    4,
		},
		TestType{
			ratings: []int{1, 2, 3, 2, 1},
			want:    9,
		},
	}
	for _, t := range ts {
		get := candy(t.ratings)
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
