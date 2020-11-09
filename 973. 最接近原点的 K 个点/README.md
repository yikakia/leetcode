# 973. 最接近原点的 K 个点
[https://leetcode-cn.com/problems/k-closest-points-to-origin/](https://leetcode-cn.com/problems/k-closest-points-to-origin/)
## 堆排序
```go

type pointHeap [][]int

func (h pointHeap) Len() int {
	return len(h)
}
func (h pointHeap) Less(i, j int) bool {
	a := h[i][0]*h[i][0] + h[i][1]*h[i][1]
	b := h[j][0]*h[j][0] + h[j][1]*h[j][1]
	return a < b
}
func (h pointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *pointHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *pointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, K int) [][]int {
	var pointheap pointHeap = points
	heap.Init(&pointheap)
	r := make([][]int, K)
	for i := 0; i < K; i++ {
		r[i] = heap.Pop(&pointheap).([]int)
	}
	return r[:K]
}

```
>执行用时：132 ms, 在所有 Go 提交中击败了64.56%的用户
内存消耗：7.7 MB, 在所有 Go 提交中击败了57.69%的用户

简单地说就是对这种求最 K 个的问题，都可以用堆排序。Golang 中用 heap 容器就可以了，只要实现了几个接口就可以了。