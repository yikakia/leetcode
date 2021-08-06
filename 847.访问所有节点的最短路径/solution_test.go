package leetcode

import (
	"reflect"
	"testing"
)

// 847.访问所有节点的最短路径
// https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, tuple{i, 1 << i, 0})
	}
	for {
		t := q[0]
		q = q[1:]
		if t.mask == (1<<n)-1 {
			return t.dist
		}
		for _, v := range graph[t.u] {
			maskV := t.mask | 1<<v
			if !seen[v][maskV] {
				q = append(q, tuple{v, maskV, t.dist + 1})
				seen[v][maskV] = true
			}
		}
	}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		graph [][]int
		want  int
		desc  string
	}{
		{
			graph: [][]int{{1, 2, 3},
				{0},
				{0},
				{0},
			},
			want: 4,
		},
		{
			graph: [][]int{{1},
				{0, 2, 4},
				{1, 3, 4},
				{2},
				{1, 2},
			},
			want: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := shortestPathLength(tC.graph)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
