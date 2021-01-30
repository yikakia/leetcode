package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 778.水位上升的泳池中游泳
// https://leetcode-cn.com/problems/swim-in-rising-water/
func swimInWater(grid [][]int) int {
	n := len(grid)
	edges := make([]edge, 0, n*n)
	for i, row := range grid {
		for j := range row {
			baseIndex := i*n + j
			if i > 0 {
				edges = append(edges,
					edge{baseIndex - n, baseIndex, max(grid[i-1][j], grid[i][j])})
			}
			if j > 0 {
				edges = append(edges,
					edge{baseIndex - 1, baseIndex, max(grid[i][j-1], grid[i][j])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].value < edges[j].value })
	uf := newUnionFind(n * n)
	for _, e := range edges {
		uf.union(e.start, e.end)
		if uf.inSameSet(0, n*n-1) {
			return e.value
		}
	}
	return 0
}

type edge struct {
	start, end, value int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type unionFind struct {
	fa, rank []int
}

func newUnionFind(n int) *unionFind {
	fa, rank := make([]int, n), make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &unionFind{fa: fa, rank: rank}
}

func (uf *unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *unionFind) union(x, y int) {
	if uf.inSameSet(x, y) {
		return
	}
	fx, fy := uf.fa[x], uf.fa[y]
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.fa[fy] = fx
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		grid [][]int
	}{
		{
			desc: "",
			grid: [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}},
			want: 16,
		},
		{
			grid: [][]int{{0, 2}, {1, 3}},
			want: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := swimInWater(tC.grid)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
