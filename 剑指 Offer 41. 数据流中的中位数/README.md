# 剑指 Offer 41. 数据流中的中位数
[https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/](https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/)

## 堆
```go
type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	maxheap *maxHeap
	minheap *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	maxh := &maxHeap{}
	minh := &minHeap{}
	heap.Init(maxh)
	heap.Init(minh)
	return MedianFinder{maxheap: maxh, minheap: minh}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxheap.Len() == this.minheap.Len() {
		heap.Push(this.maxheap, num)
		heap.Push(this.minheap, heap.Pop(this.maxheap))
	} else {
		heap.Push(this.minheap, num)
		heap.Push(this.maxheap, heap.Pop(this.minheap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxheap.Len() == this.minheap.Len() {
		return float64(((*this.maxheap)[0] + (*this.minheap)[0])) / 2.0
	}
	return float64((*this.minheap)[0])
}

```
>执行用时：108 ms, 在所有 Go 提交中击败了90.27%的用户
内存消耗：13.5 MB, 在所有 Go 提交中击败了85.29%的用户

当需要时刻找出最优先的数据的时候，就可以考虑用堆来实现。

这个地方是找中位数，于是就把数据拆分成两半。一半是大根堆，记录小的那一半数据；另一半是小根堆，记录大的那一半数据。

我们默认小根堆储存的数据量要大于等于大根堆的数据量，于是当我们要取得中位数的时候，就可以根据他们两个的数据量是否相等来确定是直接返回小根堆的堆顶，还是返回大根堆和小根堆堆顶的平均数。

而添加元素的时候根据我们的规则，就是如果两个堆的数据量相同，就先压入大根堆，再把大根堆的堆顶去掉，然后压入小根堆中。这样就让小根堆的数据量大于了大根堆。而数据量不同的时候，就先压入小根堆，再把小根堆的堆顶压入大根堆。综合这两者就可以让得到我们想要的结果。


最大的难点还是在思路上，分成两半。对于TopN 问题，用堆也是挺好的。