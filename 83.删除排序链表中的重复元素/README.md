# 83.删除排序链表中的重复元素
[https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list) 
## 原题
存在一个按升序排列的链表，给你这个链表的头节点 `head` ，请你删除所有重复的元素，使每个元素 **只出现一次** 。

返回同样按升序排列的结果链表。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/list1.jpg" style="width: 302px; height: 242px;" />
```

输入：head = [1,1,2]
输出：[1,2]

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/04/list2.jpg" style="width: 542px; height: 222px;" />
```

输入：head = [1,1,2,3,3]
输出：[1,2,3]

```
 

 **提示：** 
- 链表中节点数目在范围 `[0, 300]` 内
-  `-100 <= Node.val <= 100` 
- 题目数据保证链表已经按升序排列
 
**标签**
`链表` 


## 遍历并移动下标
```go
func deleteDuplicates(head *ListNode) *ListNode {
	for cur := head; cur != nil; {
		if next := cur.Next; next != nil && next.Val == cur.Val {
			for next = cur.Next; next != nil && next.Val == cur.Val; {
				next = next.Next
			}
			cur.Next = next
			cur = cur.Next
		} else {
			cur = cur.Next
		}

	}
	return head
}
```
>执行用时: 0 ms
内存消耗: 3.1 MB

很简单的思路，当遇到了相同的元素的时候，就不断地往后找，直到遇到了不同的为止。最后把第一个元素连接到第一个不相同的元素即可。