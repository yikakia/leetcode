# 2. 两数相加
[https://leetcode-cn.com/problems/add-two-numbers/](https://leetcode-cn.com/problems/add-two-numbers/) 
## 
```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{}
	carry := 0
	p := &res
	for l1 != nil && l2 != nil {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = l1.Val + l2.Val + carry
		carry = p.Val / 10
		p.Val %= 10
		l1, l2 = l1.Next, l2.Next
	}
	for l1 != nil {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		l1 = l1.Next
	}
	for l2 != nil {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		l2 = l2.Next
	}
	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}
	return res.Next
}
```
>执行用时: 20 ms
内存消耗: 4.9 MB

简单地说就是模拟运算，因为是相加所以进位方面还是比较简单的。
