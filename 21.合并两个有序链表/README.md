# 21. 合并两个有序链表
[https://leetcode-cn.com/problems/merge-two-sorted-lists/](https://leetcode-cn.com/problems/merge-two-sorted-lists/) 
## 归并
```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var res *ListNode
	flag := false
	if l1 == nil {
		res, l2 = l2, l2.Next
		flag = true
	}
	if !flag && l2 == nil {
		res, l1 = l1, l1.Next
		flag = true
	}
	if !flag {
		if l1.Val < l2.Val {
			res, l1 = l1, l1.Next
		} else {
			res, l2 = l2, l2.Next
		}
	}
	head := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next, l1 = l1, l1.Next
			head = head.Next
		} else {
			head.Next, l2 = l2, l2.Next
			head = head.Next
		}
	}
	for l1 != nil {
		head.Next, l1 = l1, l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next, l2 = l2, l2.Next
		head = head.Next
	}
	return res
}
```
>执行用时: 4 ms
内存消耗: 2.5 MB

写完才发现 sb 了，后面剩下的直接接在最后就可以了。
```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var res *ListNode
	flag := false
	if l1 == nil {
		res, l2 = l2, l2.Next
		flag = true
	}
	if !flag && l2 == nil {
		res, l1 = l1, l1.Next
		flag = true
	}
	if !flag {
		if l1.Val < l2.Val {
			res, l1 = l1, l1.Next
		} else {
			res, l2 = l2, l2.Next
		}
	}
	head := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next, l1 = l1, l1.Next
			head = head.Next
		} else {
			head.Next, l2 = l2, l2.Next
			head = head.Next
		}
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return res
}
```
这就是把简单的东西复杂化了吧
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了68.81%的用户

还有就是返回值可以不用先确定头部，直接接在 res.Next 后面就可以了。到时候返回的时候返回一个 res.Next 也没啥问题。之前刷剑指的时候思路很清晰，结果现在又复杂了。