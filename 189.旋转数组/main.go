package main

import (
	"fmt"
	"reflect"
)

// 189.旋转数组
// https://leetcode-cn.com/problems/rotate-array/

func resverse(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	resverse(nums)
	resverse(nums[:k])
	resverse(nums[k:])
}

func test() {
	type TestType struct {
		nums []int
		k    int
		want []int
	}
	ts := []TestType{
		TestType{
			nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		TestType{
			nums: []int{-1, -100, 3, 99}, k: 2,
			want: []int{3, 99, -1, -100},
		},
	}
	for _, t := range ts {
		rotate(t.nums, t.k)
		if !reflect.DeepEqual(t.nums, t.want) {
			// 填充输出格式
			fmt.Printf("%+v \n", t)
		}
	}
	fmt.Println("Finished Test!")
}
func main() {
	test()
}
