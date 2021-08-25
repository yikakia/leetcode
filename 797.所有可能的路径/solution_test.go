package leetcode

import (
	"reflect"
	"testing"
)

// 797.所有可能的路径
// https://leetcode-cn.com/problems/all-paths-from-source-to-target
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph) - 1
	res := [][]int{}
	var dfs func(curNode int)
	visitList := make([]int, 0, len(graph))
	dfs = func(curNode int) {
		visitList = append(visitList, curNode)
		if curNode == n {
			tmp := make([]int, len(visitList))
			copy(tmp, visitList)
			res = append(res, tmp)
			return
		}
		for _, nxtNode := range graph[curNode] {
			dfs(nxtNode)
			visitList = visitList[:len(visitList)-1]
		}
	}
	dfs(0)
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		graph [][]int
		want  [][]int
		desc  string
	}{
		{
			graph: [][]int{
				{1, 2},
				{3},
				{3},
				{},
			},
			want: [][]int{
				{0, 1, 3},
				{0, 2, 3},
			},
		},
		{
			graph: [][]int{
				{4, 3, 1},
				{3, 2, 4},
				{3},
				{4},
				{},
			},
			want: [][]int{
				{0, 4},
				{0, 3, 4},
				{0, 1, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 4},
			},
		},
		{
			graph: [][]int{
				{1},
				{},
			},
			want: [][]int{
				{0, 1},
			},
		},
		{
			graph: [][]int{
				{1, 2, 3},
				{2},
				{3},
				{},
			},
			want: [][]int{
				{0, 1, 2, 3},
				{0, 2, 3},
				{0, 3},
			},
		},
		{
			graph: [][]int{
				{1, 3},
				{2},
				{3},
				{},
			},
			want: [][]int{
				{0, 1, 2, 3},
				{0, 3},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := allPathsSourceTarget(tC.graph)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
