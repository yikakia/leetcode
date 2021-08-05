package main

import (
	"math"
	"reflect"
	"testing"
)

// 743.网络延迟时间
// https://leetcode-cn.com/problems/network-delay-time
func networkDelayTime(times [][]int, n int, k int) int {
	if n > len(times)+1 {
		return -1
	}

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = math.MaxInt32
		}
	}
	for _, time := range times {
		st, ed, cost := time[0]-1, time[1]-1, time[2]
		matrix[st][ed] = cost
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k-1] = 0

	used := make([]bool, n)
	for i := 0; i < n; i++ {
		// 找到距离最近的
		u := -1
		for tmpU, us := range used {
			if !us && (u == -1 || dist[tmpU] < dist[u]) {
				u = tmpU
			}
		}
		used[u] = true
		for v, time := range matrix[u] {
			dist[v] = min(dist[v], dist[u]+time)
		}
	}
	ret := -1
	for _, d := range dist {
		if d == math.MaxInt32 {
			return -1
		}
		ret = max(ret, d)
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		times [][]int
		n     int
		k     int
		want  int
		desc  string
	}{
		{
			times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:     4,
			k:     2,
			want:  2,
		},
		{
			times: [][]int{{1, 2, 1}},
			n:     2,
			k:     1,
			want:  1,
		},
		{
			times: [][]int{{1, 2, 1}},
			n:     2,
			k:     2,
			want:  -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := networkDelayTime(tC.times, tC.n, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
