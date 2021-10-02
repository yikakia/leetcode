package leetcode

import (
	"reflect"
	"testing"
)

// 1436.旅行终点站
// https://leetcode-cn.com/problems/destination-city
func destCity(paths [][]string) string {
	dst := map[string]int{}

	for _, path := range paths {
		dst[path[1]] = dst[path[1]]
		dst[path[0]]++
	}
	for k, v := range dst {
		if v == 0 {
			return k
		}
	}
	return ""
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		paths [][]string
		want  string
		desc  string
	}{
		{
			paths: [][]string{
				{"London", "New York"},
				{"New York", "Lima"},
				{"Lima", "Sao Paulo"},
			},
			want: "Sao Paulo",
		},
		{
			paths: [][]string{
				{"B", "C"},
				{"D", "B"},
				{"C", "A"},
			},
			want: "A",
		},
		{
			paths: [][]string{
				{"A", "Z"},
			},
			want: "Z",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := destCity(tC.paths)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
