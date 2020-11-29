# 328. 奇偶链表
[https://leetcode-cn.com/problems/odd-even-linked-list/](https://leetcode-cn.com/problems/odd-even-linked-list/)
## 双指针同时遍历
```go
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := head
	a := head
	b := head.Next
	headb := b
	for a.Next != nil && b.Next != nil {
		a.Next, b.Next = b.Next, b.Next.Next
		a, b = a.Next, b.Next
	}
	a.Next = headb
	return res
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了87.97%的用户
内存消耗：3.2 MB, 在所有 Go 提交中击败了61.45%的用户

思路很简单，就是用双指针，一个指向奇数位，一个指向偶数位，然后用这两个分别形成一个链表，最后再把这偶数位的链表头部赋给奇数位指针的`Next`就可以了