package leetcode

import (
	"reflect"
	"testing"
)

// 88.合并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array
func merge(nums1 []int, m int, nums2 []int, n int) {
	a, b := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		switch {
		case a == -1:
			nums1[i] = nums2[b]
			b--
		case b == -1:
			nums1[i] = nums1[a]
			a--
		case nums1[a] > nums2[b]:
			nums1[i] = nums1[a]
			a--
		default:
			nums1[i] = nums2[b]
			b--
		}
	}
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc            string
		nums1, nums2, r []int
		m, n            int
		want            []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.r = append(tC.r, tC.nums1...)
			merge(tC.nums1, tC.m, tC.nums2, tC.n)
			if !reflect.DeepEqual(tC.want, tC.nums1) {
				t.Errorf("input: %+v get:%v\n", tC, tC.nums1)
			}
		})
	}
}
