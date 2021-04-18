package leetcode

import (
	"reflect"
	"testing"
)

// 220.存在重复元素III
// https://leetcode-cn.com/problems/contains-duplicate-iii
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		l, r := i-k, i+k
		if l < 0 {
			l = 0
		}
		if r > n {
			r = n
		}
		for j := l; j < r; j++ {
			if i != j && abs(nums[i]-nums[j]) <= t {
				return true
			}

		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		k, t int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			k:    3,
			t:    0,
			want: true,
		},
		{
			nums: []int{1, 0, 1, 1},
			k:    1,
			t:    2,
			want: true,
		},
		{
			nums: []int{1, 5, 9, 1, 5, 9},
			k:    2,
			t:    3,
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := containsNearbyAlmostDuplicate(tC.nums, tC.k, tC.t)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
