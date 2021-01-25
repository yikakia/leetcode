package leetcode

import (
	"reflect"
	"testing"
)

// 959.由斜杠划分区域
// https://leetcode-cn.com/problems/regions-cut-by-slashes/
type UnionFind struct {
	fa, rank []int
	count    int
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &UnionFind{fa: fa, rank: rank, count: n}
}
func (uf *UnionFind) Find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.Find(uf.fa[x])
	}
	return uf.fa[x]
}
func (uf *UnionFind) Union(x, y int) bool {
	fx, fy := uf.Find(x), uf.Find(y)
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
func regionsBySlashes(grid []string) int {
	n := len(grid)
	uf := NewUnionFind(4 * n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			index := i*n + j
			/*
				______
				|\ 0/|
				|3\/ |
				| /\1|
				|/2 \|
				------
			*/
			// 把当前的▲和下面的▼相连
			if i < n-1 {
				bottom := index + n
				uf.Union(index*4+2, bottom*4)
			}
			// 把当前的最右边的三角形和右边的最左边的三角形相连
			if j < n-1 {
				right := index + 1
				uf.Union(index*4+1, right*4+3)
			}
			if grid[i][j] == '/' {
				uf.Union(index*4, index*4+3)
				uf.Union(index*4+1, index*4+2)
			} else if grid[i][j] == '\\' {
				uf.Union(index*4, index*4+1)
				uf.Union(index*4+2, index*4+3)
			} else {
				uf.Union(index*4, index*4+1)
				uf.Union(index*4+1, index*4+2)
				uf.Union(index*4+2, index*4+3)
			}

		}
	}
	return uf.count
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		grid []string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := regionsBySlashes(tC.grid)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}

}
