# 92.反转链表 II
[https://leetcode-cn.com/problems/reverse-linked-list-ii](https://leetcode-cn.com/problems/reverse-linked-list-ii) 
## 原题
给你单链表的头指针 `head` 和两个整数  `left` 和 `right` ，其中  `left <= right` 。请你反转从位置 `left` 到位置 `right` 的链表节点，返回 **反转后的链表** 。
 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/rev2ex2.jpg" style="width: 542px; height: 222px;" />
```

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

```
 **示例 2：** 

```

输入：head = [5], left = 1, right = 1
输出：[5]

```
 

 **提示：** 
- 链表中节点数目为 `n` 
-  `1 <= n <= 500` 
-  `-500 <= Node.val <= 500` 
-  `1 <= left <= right <= n` 
 

 **进阶：** 你可以使用一趟扫描完成反转吗？

 
**标签**
`链表` 


## 挨个把节点放在翻转的开头
```go
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

这里的思路是把节点挨个放到要翻转的位置上，因为可能会修改 head 指针，所以使用了 dummy 节点，从而能简化代码。