# 295.数据流的中位数
[https://leetcode-cn.com/problems/find-median-from-data-stream](https://leetcode-cn.com/problems/find-median-from-data-stream) 
## 原题
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：
- void addNum(int num) - 从数据流中添加一个整数到数据结构中。
- double findMedian() - 返回目前所有元素的中位数。
 **示例：** 

```
addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3) 
findMedian() -> 2
```
 **进阶:** 
- 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
- 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 
**标签**
`设计` `双指针` `数据流` `排序` `堆（优先队列）` 


## 两个堆记录
```go
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
	if b.Len() == 0 || num >= (*b)[0] {
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
}


func (this *MedianFinder) FindMedian() float64 {
	b, s := *this.bigPartHeap, *this.smallPartHeap		
	if b.Len() > s.Len() {
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
```
>执行用时：308 ms
内存消耗：19.9 MB

就是两个堆分别记录大小，大顶堆记录小的数，小顶堆记录大的数，当大顶堆的大小大于小顶堆的大小时，就把大顶堆的最大值放入小顶堆，否则就把小顶堆的最小值放入大顶堆。

