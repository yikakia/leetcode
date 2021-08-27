package leetcode

import (
	"container/heap"
	"fmt"
)

// 295.数据流的中位数
// https://leetcode-cn.com/problems/find-median-from-data-stream
type MedianFinder struct {
	bigPartHeap   *intHeap
	smallPartHeap *intHeap // 两个堆，大顶堆和小顶堆，大顶堆存储小的数，小顶堆存储大的数
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		bigPartHeap:   &intHeap{},
		smallPartHeap: &intHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	b, s := this.bigPartHeap, this.smallPartHeap
	if b.Len() == 0 || num < (*b)[0] {
		heap.Push(b, num)
		if b.Len()-s.Len() > 1 {
			heap.Push(s, -heap.Pop(b).(int))
		}
	} else {
		heap.Push(s, -num)
		if s.Len()-b.Len() > 0 {
			heap.Push(b, -heap.Pop(s).(int))
		}
	}
	fmt.Printf("add: %+v,%+v\n", b, s)
}

func (this *MedianFinder) FindMedian() float64 {
	b, s := *this.bigPartHeap, *this.smallPartHeap
	if b.Len() > s.Len() {
		fmt.Printf("%+v,%+v\n", b, s)
		return float64(b[0])
	} else {
		return float64(b[0]-s[0]) / 2
	}
}

type intHeap []int

func (h intHeap) Len() int { return len(h) }

func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// func TestSolution(t *testing.T) {
// 	testCases := []struct {

// 		want
// 		desc string
// 	}{
// 		{
//             want: ,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			get :=
// 			if !reflect.DeepEqual(tC.want,get){
// 				t.Errorf("input: %+v get: %v\n",tC,get)
// 			}
// 		})
// 	}
// }
