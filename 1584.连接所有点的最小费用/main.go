package main

import (
	"fmt"
	"reflect"
	"sort"
)

// 1584.连接所有点的最小费用
// https://leetcode-cn.com/problems/min-cost-to-connect-all-points/
func minCostConnectPoints(points [][]int) int {
	type edge struct {
		a, b int
		dis  int
	}
	n := len(points)
	edges := []edge{}
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range points {
		pa := points[i]
		fa[i] = i
		rank[i] = 1
		for j := i + 1; j < n; j++ {
			pb := points[j]
			edges = append(edges, edge{a: i, b: j, dis: calDis(pa, pb)})
		}
	}
	// for _, e := range edges {
	// 	fmt.Printf("%+v\n", e)
	// }
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

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis < edges[j].dis
	})

	res := 0
	left := n - 1
	for _, e := range edges {
		fx, fy := find(e.a), find(e.b)
		if fx == fy {
			continue
		}
		res += e.dis
		merge(e.a, e.b)
		left--
		if left == 0 {
			break
		}
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func calDis(a, b []int) int {
	if len(a) != 2 || len(b) != 2 {
		return 0
	}
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}
func test() {
	type TestType struct {
		points [][]int
		want   int
	}
	ts := []TestType{
		TestType{
			points: [][]int{
				[]int{0, 0}, []int{2, 2}, []int{3, 10}, []int{5, 2}, []int{7, 0},
			},
			want: 20,
		},
		TestType{
			points: [][]int{
				[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{-1, 1},
			},
			want: 4,
		},
		TestType{
			points: [][]int{
				[]int{3, 12}, []int{-2, 5}, []int{-4, 1},
			},
			want: 18,
		},
		TestType{
			points: [][]int{
				[]int{-1000000, -1000000}, []int{1000000, 1000000},
			},
			want: 4000000,
		},
		TestType{
			points: [][]int{
				[]int{0, 0},
			},
			want: 0,
		},
	}
	for _, t := range ts {
		get := minCostConnectPoints(t.points)
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
