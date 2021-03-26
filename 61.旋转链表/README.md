# 61.旋转链表
[https://leetcode-cn.com/problems/rotate-list](https://leetcode-cn.com/problems/rotate-list) 
## 原题
给定一个链表，旋转链表，将链表每个节点向右移动 *k* 个位置，其中 *k* 是非负数。

 **示例 1:** 

```
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

```
 **示例 2:** 

```
输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
```
 
**标签**
`链表` `双指针` 


## 快慢指针
```go
func Len(head *ListNode) (n int) {
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := Len(head)
	k = k % n
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
		if fast == nil {
			fast = head
		}
	}
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if fast == slow {
		return head
	}
	res := slow.Next
	slow.Next = nil
	fast.Next = head
	return res
}
```
>执行用时: 4 ms
内存消耗: 2.6 MB

比较标准的快慢指针的题。其实主要要做的就是找到尾端（用来连接原来的头部）和第 `n-k` (`n` 为链表的长度)个节点（新的头节点）。于是就用快慢指针来做了，快指针先跑 `k` 个，然后快慢指针一起跑，当快指针到了尾端的时候，慢指针就是第 `n-k` 个节点。然后先把尾端和原来的头部相连，再把慢指针从原来的链表中断开即可。

需要注意的是 `k` 可能大于 `n` ，这个时候就 $ k = k%n $ 取余即可。样例里面可能有 `k` 远大于 `n` 的存在，这种时候取余就很节省时间了。

我代码里面的
```go
		if fast == nil {
			fast = head
		}
```

这一段就是最开始没有取余时的操作，取了余之后这个操作就没必要了。