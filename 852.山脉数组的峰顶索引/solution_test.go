package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 852.山脉数组的峰顶索引
// https://leetcode-cn.com/problems/peak-index-in-a-mountain-array
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool {
		return arr[i] > arr[i+1]
	})
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
		desc string
	}{
		{
			arr:  []int{0, 1, 0},
			want: 1,
		},
		{
			arr:  []int{0, 2, 1, 0},
			want: 1,
		},
		{
			arr:  []int{0, 10, 5, 2},
			want: 1,
		},
		{
			arr:  []int{3, 4, 5, 1},
			want: 2,
		},
		{
			arr:  []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			want: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := peakIndexInMountainArray(tC.arr)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
