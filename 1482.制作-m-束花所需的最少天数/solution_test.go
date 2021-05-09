package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 1482.制作m束花所需的最少天数
// https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets
func minDays(bloomDay []int, m int, k int) int {
	total := m * k
	if total > len(bloomDay) {
		return -1
	}

	var maxDay int

	for _, d := range bloomDay {
		if d > maxDay {
			maxDay = d
		}
	}

	return sort.Search(maxDay, func(testDays int) bool {
		var flowers, set int
		for _, day := range bloomDay {
			if day > testDays {
				flowers = 0
			} else {
				flowers++
				if flowers == k {
					set++
					flowers = 0
				}
			}
		}
		return set >= m
	})

}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		bloomDay []int
		m        int
		k        int
		want     int
	}{
		{
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        1,
			want:     3,
		},
		{
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        2,
			want:     -1,
		},
		{
			bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
			m:        2,
			k:        3,
			want:     12,
		},
		{
			bloomDay: []int{1000000000, 1000000000},
			m:        1,
			k:        1,
			want:     1000000000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := minDays(tC.bloomDay, tC.m, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
