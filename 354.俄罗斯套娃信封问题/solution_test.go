package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 354.俄罗斯套娃信封问题
// https://leetcode-cn.com/problems/russian-doll-envelopes
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] ||
			a[0] == b[0] && a[1] > b[1]
	})

	f := []int{}

	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc      string
		envelopes [][]int
		want      int
	}{
		{
			envelopes: [][]int{{1, 2}, {2, 3}, {3, 1}},
			want:      2,
		},
		{
			envelopes: [][]int{{1, 3}, {1, 2}, {1, 1}, {2, 2}, {3, 3}},
			want:      3,
		},
		{
			envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			want:      3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := maxEnvelopes(tC.envelopes)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
