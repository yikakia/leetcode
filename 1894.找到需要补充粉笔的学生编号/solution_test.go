package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 1894.找到需要补充粉笔的学生编号
// https://leetcode-cn.com/problems/find-the-student-that-will-replace-the-chalk
func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	for i := 1; i < n; i++ {
		chalk[i] += chalk[i-1]
	}
	for k >= 0 {
		index := sort.SearchInts(chalk, k)
		if index < n {
			if chalk[index] == k {
				return (index + 1) % n
			}
			return index
		} else {
			k %= chalk[n-1]
		}
	}
	return 0
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		chalk []int
		k     int
		want  int
		desc  string
	}{
		{
			chalk: []int{5, 1, 5},
			k:     23,
			want:  0,
		},
		{
			chalk: []int{5, 1, 5},
			k:     22,
			want:  0,
		},
		{
			chalk: []int{3, 4, 1, 2},
			k:     25,
			want:  1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := chalkReplacer(tC.chalk, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
