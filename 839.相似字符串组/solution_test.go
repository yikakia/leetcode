package leetcode

import (
	"reflect"
	"testing"
)

// 839.相似字符串组
// https://leetcode-cn.com/problems/similar-string-groups/
func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := newUnionFind(n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if isSame(strs[i], strs[j]) {
				uf.union(i, j)
			}
		}
	}
	return uf.count
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
	return &unionFind{fa, rank, n}
}

func (uf *unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.fa[fy] = fx
	uf.rank[fx] += uf.rank[fy]
	uf.count--
}

func isSame(a, b string) bool {
	if a == b {
		return true
	}
	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
		if count > 2 {
			return false
		}
	}
	return true
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		strs []string
	}{
		{
			desc: "",
			strs: []string{"tars", "rats", "arts", "star"},
			want: 2,
		},
		{
			strs: []string{"omv", "ovm"},
			want: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numSimilarGroups(tC.strs)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
