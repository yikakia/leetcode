package main

import (
	"fmt"
	"reflect"
)

// 684.冗余连接
// https://leetcode-cn.com/problems/redundant-connection/
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	fa := make([]int, n+1)
	rank := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx
	}

	res := make([]int, 2)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if find(x) == find(y) {
			res = edge
			return res
		}
		merge(x, y)
	}

	return res
}
func test() {
	type TestType struct {
		edges [][]int
		want  []int
	}
	ts := []TestType{
		TestType{},
	}
	for _, t := range ts {
		get := findRedundantConnection(t.edges)
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
