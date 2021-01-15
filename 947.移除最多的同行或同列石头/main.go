package main

import (
	"fmt"
	"reflect"
)

// 947.移除最多的同行或同列石头
// https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column/
func removeStones(stones [][]int) int {
	fa, rank := map[int]int{}, map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if _, has := fa[x]; !has {
			fa[x], rank[x] = x, 1
		}
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	union := func(x, y int) {
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

	for _, p := range stones {
		union(p[0], p[1]+10001)
	}
	ans := len(stones)
	for x, fx := range fa {
		if x == fx {
			ans--
		}
	}
	return ans
}

func test() {
	type TestType struct {
		stones [][]int
		want   int
	}
	ts := []TestType{
		TestType{},
	}
	for _, t := range ts {
		get := removeStones(t.stones)
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
