# 剑指 Offer 52. 两个链表的第一个公共节点
[https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/](https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/)

## 先调整到相同的长度再进行计算
```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var r *ListNode
	lenA, lenB := getLen(headA), getLen(headB)
	tmpA, tmpB := headA, headB
	for i := 0; i <= lenA-lenB-1; i++ {
		tmpA = tmpA.Next
	}
	for i := 0; i <= lenB-lenA-1; i++ {
		tmpB = tmpB.Next
	}
	for tmpA != nil && tmpA != tmpB {
		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}
	if tmpA != nil && tmpA == tmpB {
		r = tmpA
	}
	return r
}

func getLen(head *ListNode) int {
	r := 0
	for head != nil {
		r++
		head = head.Next
	}
	return r
}
```

>执行用时：36 ms, 在所有 Go 提交中击败了97.21%的用户
内存消耗：7.7 MB, 在所有 Go 提交中击败了36.22%的用户

简单地说就是先把两个调整到同样的长度，然后同时往后找。最后判断两个是否相等就好了。如果相等，那么这个就是我们要找的第一个公共点。