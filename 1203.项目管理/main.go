package main

import (
	"fmt"
	"reflect"
)

// 1203.项目管理
// https://leetcode-cn.com/problems/sort-items-by-groups-respecting-dependencies/
func topSort(graph [][]int, deg, items []int) (orders []int) {
	q := []int{}
	for _, i := range items {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		orders = append(orders, from)
		for _, to := range graph[from] {
			deg[to]--
			if deg[to] == 0 {
				q = append(q, to)
			}
		}
	}
	return
}

func sortItems(n, m int, group []int, beforeItems [][]int) (ans []int) {
	groupItems := make([][]int, m+n) // groupItems[i] 表示第 i 个组负责的项目列表
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i // 给不属于任何组的项目分配一个组
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupGraph := make([][]int, m+n) // 组间依赖关系
	groupDegree := make([]int, m+n)
	itemGraph := make([][]int, n) // 组内依赖关系
	itemDegree := make([]int, n)
	for cur, items := range beforeItems {
		curGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != curGroupID { // 不同组项目，确定组间依赖关系
				groupGraph[preGroupID] = append(groupGraph[preGroupID], curGroupID)
				groupDegree[curGroupID]++
			} else { // 同组项目，确定组内依赖关系
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemDegree[cur]++
			}
		}
	}

	// 组间拓扑序
	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}
	groupOrders := topSort(groupGraph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}

	// 按照组间的拓扑序，依次求得各个组的组内拓扑序，构成答案
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := topSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return
}

func test() {
	type TestType struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
		want        []int
	}
	ts := []TestType{
		TestType{
			n: 5, m: 5,
			group:       []int{2, 0, -1, 3, 0},
			beforeItems: [][]int{[]int{2, 1, 3}, []int{2, 4}, []int{}, []int{}, []int{}},
			want:        []int{3, 2, 4, 1, 0},
		},
		TestType{
			n:           8,
			m:           2,
			group:       []int{-1, -1, 1, 0, 0, 1, 0, -1},
			beforeItems: [][]int{[]int{}, []int{6}, []int{5}, []int{6}, []int{3, 6}, []int{}, []int{}, []int{}},
			want:        []int{6, 3, 4, 1, 5, 2, 0, 7},
		},
	}
	for _, t := range ts {
		get := sortItems(t.n, t.m, t.group, t.beforeItems)
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
