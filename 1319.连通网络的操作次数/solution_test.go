package leetcode

import (
	"reflect"
	"testing"
)

// 1319.连通网络的操作次数
// https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected/
type UnionFindSet struct {
	fa       []int
	rank     []int
	setCount int
}

func NewUnionFind(n int) *UnionFindSet {
	uf := UnionFindSet{fa: make([]int, n), rank: make([]int, n), setCount: n}
	for i := range uf.fa {
		uf.fa[i] = i
		uf.rank[i] = 1
	}
	return &uf
}

func (uf *UnionFindSet) reset() {
	for i := range uf.fa {
		uf.fa[i] = i
		uf.rank[i] = 1
	}
}

func (uf *UnionFindSet) Find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.Find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *UnionFindSet) Merge(x, y int) bool {
	fx, fy := uf.Find(x), uf.Find(y)
	if fx == fy {
		return false
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.fa[fy] = fx
	uf.setCount--
	return true
}
func makeConnected(n int, connections [][]int) int {
	if len(connections)+1 < n {
		return -1
	}

	uf := NewUnionFind(n)
	for _, e := range connections {
		uf.Merge(e[0], e[1])
	}

	return uf.setCount - 1
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc        string
		want        int
		n           int
		connections [][]int
	}{
		{
			desc:        "",
			want:        1,
			n:           4,
			connections: [][]int{[]int{0, 1}, []int{0, 2}, []int{1, 2}},
		},
		{
			want:        2,
			n:           6,
			connections: [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 2}, []int{1, 3}},
		},
		{
			want:        -1,
			n:           6,
			connections: [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 2}},
		},
		{
			want:        0,
			n:           5,
			connections: [][]int{[]int{0, 1}, []int{0, 2}, []int{3, 4}, []int{2, 3}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := makeConnected(tC.n, tC.connections)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
