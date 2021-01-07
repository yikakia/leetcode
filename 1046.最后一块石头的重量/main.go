package main

import (
	"container/heap"
	"fmt"
)

// 1046.最后一块石头的重量
// https://leetcode-cn.com/problems/last-stone-weight/

type intHeap []int

func (h intHeap) Len() int { return len(h) }
func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func lastStoneWeight(stones []int) int {
	h := &intHeap{}
	for _, stone := range stones {
		heap.Push(h, stone)
	}

	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		if x > y {
			heap.Push(h, x-y)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h.Pop().(int)
}

func test() {
	type TestType struct {
		stones []int
		want   int
	}
	ts := []TestType{
		TestType{
			stones: []int{2, 7, 4, 1, 8, 1}, want: 1,
		},
	}
	for _, t := range ts {
		get := lastStoneWeight(t.stones)
		if get != t.want {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")

}
func main() {
	test()
}
