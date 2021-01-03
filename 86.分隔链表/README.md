# 86. 分隔链表
[https://leetcode-cn.com/problems/partition-list/](https://leetcode-cn.com/problems/partition-list/) 
## 原题
给你一个链表和一个特定值 `x` ，请你对链表进行分隔，使得所有小于 `x` 的节点都出现在大于或等于 `x` 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。
 
**示例：** 
```
输入：head = 1->4->3->2->5->2, x = 3
输出：1->2->2->4->3->5
```


## 双虚拟节点
```go
func partition(head *ListNode, x int) *ListNode {
	smallHead := ListNode{}
	bigHead := ListNode{}
	tmp := head
	small := &smallHead
	big := &bigHead
	for tmp != nil {
		if tmp.Val < x {
			small.Next = tmp
			small = small.Next
		} else {
			big.Next = tmp
			big = big.Next
		}
		tmp = tmp.Next
	}
	small.Next = bigHead.Next
	big.Next = nil
	return smallHead.Next
}
```
>执行用时: 0 ms
内存消耗: 2.4 MB

