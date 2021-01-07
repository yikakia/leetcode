package main

import (
	"fmt"
	"reflect"
)

// 547.省份数量
// https://leetcode-cn.com/problems/number-of-provinces/
func findCircleNum(isConnected [][]int) int {
	fa := make([]int, len(isConnected))
	for i := range fa {
		fa[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for i := range isConnected {
		for j := range isConnected[i] {
			if isConnected[i][j] == 1 {
				fFrom, fTo := find(i), find(j)
				fa[fFrom] = fTo
			}
		}
	}

	ans := make(map[int]bool)
	for _, x := range fa {
		ans[find(x)] = true
	}
	return len(ans)
}
func test() {
	type TestType struct {
		isConnected [][]int
		want        int
	}
	ts := []TestType{
		TestType{
			isConnected: [][]int{
				[]int{1, 1, 0},
				[]int{1, 1, 0},
				[]int{0, 0, 1},
			},
			want: 2,
		},
		TestType{
			isConnected: [][]int{
				[]int{1, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 1},
			},
			want: 3,
		},
	}
	for _, t := range ts {
		get := findCircleNum(t.isConnected)
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
