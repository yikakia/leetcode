# 面试题 17.14.最小K个数
[https://leetcode-cn.com/problems/smallest-k-lcci](https://leetcode-cn.com/problems/smallest-k-lcci) 
## 原题
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

 **示例：** 

```
输入： arr = [1,3,5,7,2,4,6,8], k = 4
输出： [1,2,3,4]

```
 **提示：** 
-  `0 <= len(arr) <= 100000` 
-  `0 <= k <= min(100000, len(arr))` 
 
**标签**
`数组` `分治` `快速选择` `排序` `堆（优先队列）` 


## 简易做法 堆
```go
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
```
>执行用时：56 ms
内存消耗：7.3 MB

简单的做法就是用堆。