package leetcode

import (
	"reflect"
	"testing"
)

// 815.公交路线
// https://leetcode-cn.com/problems/bus-routes
func numBusesToDestination(routes [][]int, source int, target int) int {
	station2routes := map[int][]int{}
	for route := range routes {
		for _, station := range routes[route] {
			station2routes[station] = append(station2routes[station], route)
		}
	}
	type pair struct {
		station int
		steps   int
	}
	visited := map[int]bool{}
	q := []pair{{source, 0}}
	for len(q) != 0 {
		x := q[0]
		q = q[1:]
		if x.station == target {
			return x.steps
		}
		for _, route := range station2routes[x.station] {
			for i := range routes[route] {
				if !visited[routes[route][i]] && routes[route][i] != x.station {
					q = append(q, pair{routes[route][i], x.steps + 1})
					visited[routes[route][i]] = true
				}
			}
		}
	}
	return -1
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		routes [][]int
		source int
		target int
		want   int
		desc   string
	}{
		{
			routes: [][]int{{1, 2, 7}, {3, 6, 7}},
			source: 1,
			target: 6,
			want:   2,
		},
		{
			routes: [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}},
			source: 15,
			target: 12,
			want:   -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numBusesToDestination(tC.routes, tC.source, tC.target)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
