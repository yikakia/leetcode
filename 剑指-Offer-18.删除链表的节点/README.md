# 剑指 Offer 18. 删除链表的节点
[https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/](https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/)
```go
func deleteNode(head *ListNode, val int) *ListNode {
	// 排除第一个就是的情况
	if head.Val == val {
		return head.Next
	}
	var father *ListNode = &ListNode{}
	var p *ListNode = head
	for p != nil {
		if p.Val == val {
			father.Next = p.Next
			break
		}
		father = p
		p = p.Next
	}
	return head
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了71.75%的用户
内存消耗：2.8 MB, 在所有 Go 提交中击败了53.73%的用户

就是对于链表的应用。对于一般性而言，定一个父指针用于指向当前节点的父亲。如果该节点就是要找的节点的话就直接修改其父亲的`Next`为自己的`Next`就可以了。

但是第一个就是的话就不同了，因为他没有父亲。如果非要这么做的话，就要新声明一个节点用于初始化父亲，并且设一个和头指针不同的值。这样在循环后判断父指针的Next和头指针相同，而Val不同，就可以判断这个父指针到底是表示头指针就是要找的，还是头指针的儿子就是要找的了。因为我这里是设头指针不动，所以就这么写了。