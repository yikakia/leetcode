# 703.数据流中的第 K 大元素
[https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/](https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/) 
## 原题
设计一个找到数据流中第 `k` 大元素的类（class）。注意是排序后的第 `k` 大元素，不是第 `k` 个不同的元素。

请实现 `KthLargest` 类：
- `KthLargest(int k, int[] nums)` 使用整数 `k` 和整数流 `nums` 初始化对象。
- `int add(int val)` 将 `val` 插入数据流 `nums` 后，返回当前数据流中第 `k` 大的元素。
 

 **示例：** 

```

输入：
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
输出：
[null, 4, 5, 5, 8, 8]

解释：
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8

```
 
 **提示：** 
- `1 <= k <= 10^4`
- `0 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `-10^4 <= val <= 10^4`
- 最多调用 `add` 方法 `10^4` 次
- 题目数据保证，在查找第 `k` 大元素时，数组中至少有 `k` 个元素
 
**标签**
`堆` `设计` 


## 
```go
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
```
>执行用时: 36 ms
内存消耗: 8.2 MB

就是开一个大小为 `k` 的小顶堆就可以实现 `TopK` 了。