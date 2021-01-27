package leetcode

import (
	"reflect"
	"testing"
)

// 1579.保证图可完全遍历
// https://leetcode-cn.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
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
		count: n,
	}
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
	uf.count--
	return true
}
func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}
func (uf *unionFind) reset() {
	uf.count = len(uf.fa)
	for i := range uf.fa {
		uf.fa[i] = i
		uf.rank[i] = 1
	}
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	res := len(edges)
	ufa, ufb := newUnionFind(n), newUnionFind(n)
	for _, e := range edges {
		if e[0] == 3 {
			if !ufa.inSameSet(e[1]-1, e[2]-1) || !ufb.inSameSet(e[1]-1, e[2]-1) {
				ufa.union(e[1]-1, e[2]-1)
				ufb.union(e[1]-1, e[2]-1)
				res--
			}
		}
	}
	for _, e := range edges {
		switch e[0] {
		case 1:
			if ufa.union(e[1]-1, e[2]-1) {
				res--
			}
		case 2:
			if ufb.union(e[1]-1, e[2]-1) {
				res--
			}
		}
	}
	if ufa.count > 1 || ufb.count > 1 {
		return -1
	}
	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		want  int
		n     int
		edges [][]int
	}{
		{
			desc:  "",
			n:     4,
			edges: [][]int{[]int{3, 1, 2}, []int{3, 2, 3}, []int{1, 1, 3}, []int{1, 2, 4}, []int{1, 1, 2}, []int{2, 3, 4}},
			want:  2,
		},
		{
			n:     4,
			edges: [][]int{[]int{3, 1, 2}, []int{3, 2, 3}, []int{1, 1, 4}, []int{2, 1, 4}},
			want:  0,
		},
		{
			n:     4,
			edges: [][]int{[]int{3, 2, 3}, []int{1, 1, 2}, []int{2, 3, 4}},
			want:  -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := maxNumEdgesToRemove(tC.n, tC.edges)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
