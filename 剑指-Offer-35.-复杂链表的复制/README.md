# 剑指 Offer 35. 复杂链表的复制
[https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/](https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/)

## 复制节点并松弛节点
```go

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		tmp := &Node{Val: p.Val, Next: p.Next}
		p.Next = tmp
		p = p.Next.Next
	}

	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	// 拆分
	r := head.Next
	p = head
	for p.Next.Next != nil {
		trueNext := p.Next.Next
		p.Next.Next = trueNext.Next
		p.Next = trueNext
		p = trueNext
	}
	p.Next.Next = nil
	p.Next = nil
	return r
}

```

>执行用时：4 ms, 在所有 Go 提交中击败了15.45%的用户
内存消耗：3.4 MB, 在所有 Go 提交中击败了57.14%的用户

简单地说就是在目前的链表中先复制每一个节点插入到链表中。然后把每个新复制的节点的`Random`指向应该指向的位置。然后把每个原始节点的`Next`和每个复制出来的节点的`Next`分别指向应该指向的位置就可以了。