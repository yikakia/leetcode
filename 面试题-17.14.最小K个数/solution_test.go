package leetcode

import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

// 面试题17.14.最小K个数
// https://leetcode-cn.com/problems/smallest-k-lcci
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if k >= len(arr) {
		return arr
	}
	h := &intHeap{IntSlice: arr}
	heap.Init(h)
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

type intHeap struct {
	sort.IntSlice
}

func (ih *intHeap) Pop() interface{} {
	v := (*ih).IntSlice[len((*ih).IntSlice)-1]
	(*ih).IntSlice = (*ih).IntSlice[:len((*ih).IntSlice)-1]
	return v
}
func (ih *intHeap) Push(x interface{}) {
	(*ih).IntSlice = append((*ih).IntSlice, x.(int))
}

var _ heap.Interface = (*intHeap)(nil)

func TestSolution(t *testing.T) {
	testCases := []struct {
		arr  []int
		k    int
		want []int
		desc string
	}{
		{
			arr:  []int{1, 3, 5, 7, 2, 4, 6, 8},
			k:    4,
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := smallestK(tC.arr, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
