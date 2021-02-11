package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

// 703.数据流中的第K大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	old := kl.IntSlice
	ele := old[len(old)-1]
	kl.IntSlice = old[:len(old)-1]
	return ele
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

var _ heap.Interface = (*KthLargest)(nil)

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func Test(t *testing.T) {
	kl := Constructor(3, []int{4, 5, 8, 2})
	tc := []int{3, 5, 10, 9, 4}
	for _, m := range tc {
		fmt.Print(kl.Add(m), " ")
	}
}
