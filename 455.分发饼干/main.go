package main

import (
	"fmt"
	"sort"
)

// 455.分发饼干
// https://leetcode-cn.com/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	res := 0
	for i < len(g) && j < len(s) {
		for i < len(g) && j < len(s) && s[j] >= g[i] {
			res++
			i++
			j++
		}
		for i < len(g) && j < len(s) && s[j] < g[i] {
			j++
		}
	}
	return res
}
func test() {
	type TestType struct {
		g    []int
		s    []int
		want int
	}
	ts := []TestType{
		TestType{
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			want: 1,
		},
		TestType{
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			want: 2,
		},
		TestType{},
	}
	for _, t := range ts {
		get := findContentChildren(t.g, t.s)
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
