# 剑指 Offer 22. 链表中倒数第k个节点
[https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)

## 一般方法

```go

func getKthFromEnd(head *ListNode, k int) *ListNode {
	len := getLen(head)
	for i := 1; i <= len-k; i++ {
		head = head.Next
	}
	return head
}
func getLen(head *ListNode) (len int) {
	for head != nil {
		len++
		head = head.Next
	}
	return len
}
```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了36.64%的用户

思路很简单，找出长度，然后一步步走。

## 快慢指针
看题解还有种方法，是快慢指针。简单的说是快指针先走`k-1`步，然后快慢指针一起走，直到快指针到达最后一个节点。此时慢指针所在的点即为倒数第`k`个节点。这就不上代码了。