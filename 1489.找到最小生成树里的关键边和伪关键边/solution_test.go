package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 1489.找到最小生成树里的关键边和伪关键边
// https://leetcode-cn.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	for i := range edges {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) bool {
		fx, fy := find(x), find(y)
		if fx == fy {
			return false
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx

		return true
	}

	minSum := 0
	for _, e := range edges {
		x, y := e[0], e[1]
		if merge(x, y) {
			minSum += e[2]
		}
	}

	critical := []int{}
	notcritical := []int{}
	for i := range edges {
		tm := 0
		m := 0
		reset(fa, rank)
		for j, e := range edges {
			if j == i {
				continue
			}
			x, y := e[0], e[1]
			if merge(x, y) {
				tm += e[2]
				m++
			}
		}
		if tm > minSum || m+1 != n {
			critical = append(critical, edges[i][3])
		}
	}

	sort.Slice(critical, func(i, j int) bool {
		return critical[i] < critical[j]
	})

	for _, e := range edges {
		index := sort.Search(len(critical), func(x int) bool {
			return critical[x] >= e[3]
		})
		if index < len(critical) && critical[index] == e[3] {
			continue
		}

		tm := 0
		reset(fa, rank)
		merge(e[0], e[1])
		tm += e[2]
		for _, e := range edges {
			x, y := e[0], e[1]
			if merge(x, y) {
				tm += e[2]
			}
		}
		if tm == minSum {
			notcritical = append(notcritical, e[3])
		}
	}

	return [][]int{critical, notcritical}
}
func reset(fa, rank []int) {
	if len(fa) != len(rank) {
		return
	}
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		want  [][]int
		n     int
		edges [][]int
	}{
		{
			desc:  "",
			n:     5,
			edges: [][]int{[]int{0, 1, 1}, []int{1, 2, 1}, []int{2, 3, 2}, []int{0, 3, 2}, []int{0, 4, 3}, []int{3, 4, 3}, []int{1, 4, 6}},
			want:  [][]int{[]int{0, 1}, []int{2, 3, 4, 5}},
		},
		{
			n:     6,
			edges: [][]int{[]int{0, 1, 1}, []int{1, 2, 1}, []int{0, 2, 1}, []int{2, 3, 4}, []int{3, 4, 2}, []int{3, 5, 2}, []int{4, 5, 2}},
			want:  [][]int{[]int{3}, []int{0, 1, 2, 4, 5, 6}},
		},
		{
			n:     6,
			edges: [][]int{[]int{0, 1, 2}, []int{0, 2, 5}, []int{2, 3, 5}, []int{1, 4, 4}, []int{2, 5, 5}, []int{4, 5, 2}},
			want:  [][]int{[]int{0, 2, 3, 5}, []int{1, 4}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findCriticalAndPseudoCriticalEdges(tC.n, tC.edges)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
