# 82.删除排序链表中的重复元素 II
[https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii) 
## 原题
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 *没有重复出现* 的数字。

 **示例 1:** 

```
输入: 1->2->3->3->4->4->5
输出: 1->2->5

```
 **示例 2:** 

```
输入: 1->1->1->2->3
输出: 2->3
```
 
**标签**
`链表` 


## 哑节点
```go
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for cur, pre := dummy.Next, dummy; cur != nil; {
		if next := cur.Next; next != nil && next.Val == pre.Next.Val {
			for next != nil && next.Val == pre.Next.Val {
				cur = cur.Next
				next = next.Next
			}
			cur = next
			pre.Next = cur
			continue
		}
		cur = cur.Next
		pre = pre.Next
	}
	return dummy.Next
}
```
>执行用时: 4 ms
内存消耗: 3 MB

因为是预先排过序的，所以遇到了重复的就可以不断地进行去除，直到遇到下一个元素。