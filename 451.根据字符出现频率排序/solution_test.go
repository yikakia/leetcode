package leetcode

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

// 451.根据字符出现频率排序
// https://leetcode-cn.com/problems/sort-characters-by-frequency
func frequencySort(s string) string {
	type pair struct {
		ch  rune
		frq int
	}
	frqs := map[rune]int{}
	for _, ch := range s {
		frqs[ch]++
	}
	ps := make([]pair, 0, len(frqs))
	for k, v := range frqs {
		ps = append(ps, pair{k, v})
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].frq > ps[j].frq
	})

	res := ""
	for _, p := range ps {
		res += strings.Repeat(string(p.ch), p.frq)
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		s    string
		want string
		desc string
	}{
		{
			s:    "tree",
			want: "eert",
		},
		{
			s:    "cccaaa",
			want: "cccaaa",
		},
		{
			s:    "Aabb",
			want: "bbAa",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := frequencySort(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
