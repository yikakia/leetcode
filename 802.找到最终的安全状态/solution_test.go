package leetcode

import (
	"reflect"
	"testing"
)

// 802.找到最终的安全状态
// https://leetcode-cn.com/problems/find-eventual-safe-states
func eventualSafeNodes(graph [][]int) (res []int) {
	isSafeMap := map[int]bool{}
	for i := range graph {
		if len(graph[i]) == 0 {
			isSafeMap[i] = true
		}
	}
	var isSafe func(int) bool
	visitMap := map[int]bool{}
	isSafe = func(x int) bool {
		if safe, ok := isSafeMap[x]; ok {
			return safe
		}
		visitMap[x] = true
		for _, next := range graph[x] {
			if visitMap[next] {
				delete(visitMap, next)
				isSafeMap[next] = false
				isSafeMap[x] = false
				return false
			}
			if !isSafe(next) {
				delete(visitMap, next)
				isSafeMap[x] = false
				return false
			}
		}
		delete(visitMap, x)
		isSafeMap[x] = true
		return true
	}
	for i := range graph {
		if isSafe(i) {
			res = append(res, i)
		}
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		graph [][]int
		want  []int
		desc  string
	}{
		{
			graph: [][]int{
				{1, 2},
				{2, 3},
				{5},
				{0},
				{5},
				{},
				{},
			},
			want: []int{2, 4, 5, 6},
		},
		{
			graph: [][]int{
				{1, 2, 3, 4},
				{1, 2},
				{3, 4},
				{0, 4},
				{},
			},
			want: []int{4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := eventualSafeNodes(tC.graph)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
