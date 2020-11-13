# 剑指 Offer 06. 从尾到头打印链表
[https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

## 记录然后输出
```go
func reversePrint(head *ListNode) []int {
	len := 0
	p := head
	for p != nil {
		len++
		p = p.Next
	}
	r := make([]int, len)
	p = head
	i := 0
	for p != nil {
		r[len-i-1] = p.Val
		p = p.Next
		i++
	}
	return r
}
```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.8 MB, 在所有 Go 提交中击败了95.38%的用户

水题一道，主要是链表的知识。可以通过先通过遍历确定长度，然后开辟数组，再遍历一次，从后到前放入数字。

当然也可以用栈来存放，这样的效率会慢一点，因为有了栈的创建和进栈出栈的操作。