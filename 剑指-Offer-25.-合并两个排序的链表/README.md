# 剑指 Offer 25. 合并两个排序的链表
[https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/](https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/)

## 
```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var headPointer *ListNode = l1
	var curPointer *ListNode = l1
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		headPointer = l2
		curPointer = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		for l1 != nil && l2 != nil && l1.Val <= l2.Val {
			curPointer, curPointer.Next, l1 = l1, l1, l1.Next
		}
		for l1 != nil && l2 != nil && l2.Val < l1.Val {
			curPointer, curPointer.Next, l2 = l2, l2, l2.Next
		}
	}
	for l1 != nil {
		curPointer, curPointer.Next, l1 = l1, l1, l1.Next
	}
	for l2 != nil {
		curPointer, curPointer.Next, l2 = l2, l2, l2.Next
	}
	return headPointer
}
```
>执行用时：8 ms, 在所有 Go 提交中击败了77.36%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了54.50%的用户

空指针还是得特判，不然就会报错，很麻烦。具体的就是归并排序的思路，不过因为是链表，所以得考虑遍历到空指针时该怎么办。

