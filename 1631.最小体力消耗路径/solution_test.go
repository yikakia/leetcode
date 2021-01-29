package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 1631.最小体力消耗路径
// https://leetcode-cn.com/problems/path-with-minimum-effort/
func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	edges := make([]edge, 0, n*m)
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			// 和上方的差值
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h - heights[i-1][j])})
			}
			// 和左边的差值
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].value < edges[j].value
	})

	uf := newUnionFind(n * m)
	for _, e := range edges {
		uf.union(e.start, e.end)
		if uf.inSameSet(0, n*m-1) {
			return e.value
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type edge struct {
	start, end, value int
}

type unionFind struct {
	fa, rank []int
	count    int
}

func newUnionFind(n int) *unionFind {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &unionFind{
		fa:    fa,
		rank:  rank,
		count: n}
}

func (uf *unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.fa[fy] = fx
	return true
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		want    int
		heights [][]int
	}{
		{
			desc:    "",
			heights: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			want:    2,
		},
		{
			heights: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			want:    1,
		},
		{
			heights: [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
			want:    0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := minimumEffortPath(tC.heights)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
