package leetcode

import (
	"reflect"
	"testing"
)

// 765.情侣牵手
// https://leetcode-cn.com/problems/couples-holding-hands/
type unionFind struct {
	fa   []int
	rank []int
	size int
}

func newUnionFind(n int) *unionFind {
	tFa := make([]int, n)
	tRank := make([]int, n)
	for i := 0; i < n; i++ {
		tFa[i] = i
		tRank[i] = 1
	}
	return &unionFind{fa: tFa, rank: tRank, size: n}
}
func (uf *unionFind) Find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.Find(uf.fa[x])
	}
	return uf.fa[x]
}
func (uf *unionFind) Union(x, y int) {
	fx, fy := uf.Find(x), uf.Find(y)
	if fx == fy {
		return
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.fa[fy] = fx
	uf.rank[fx] += uf.rank[fy]
	uf.size--
}

func minSwapsCouples(row []int) int {
	n := len(row) / 2
	uf := newUnionFind(len(row))
	for i := 0; i < n; i++ {
		uf.Union(row[i*2], row[i*2+1])
		uf.Union(i*2, i*2+1)
	}

	return n - uf.size
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		row  []int
		want int
	}{
		{
			row:  []int{0, 2, 1, 3},
			want: 1,
		},
		{
			row:  []int{3, 2, 0, 1},
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := minSwapsCouples(tC.row)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
