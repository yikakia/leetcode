# 剑指 Offer 24. 反转链表
[https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)

## 
```go
func reverseList(head *ListNode) *ListNode {
	var curPointer *ListNode = &ListNode{}
	for head != nil {
		curPointer.Next, head.Next, head = head, curPointer.Next, head.Next
	}

	return curPointer.Next
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了66.17%的用户

其实在草稿纸上画一个就好了。定一个空指针，不断地遍历链表，然后把空指针的`Next`赋给当前节点，空指针的`Next`指向当前节点，当前节点向后遍历。当然这里有个覆盖的问题，所以不支持多值赋值的语言需要定一个临时变量来储存当前节点的值。