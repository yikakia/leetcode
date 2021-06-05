# 203.移除链表元素
[https://leetcode-cn.com/problems/remove-linked-list-elements](https://leetcode-cn.com/problems/remove-linked-list-elements) 
## 原题
给你一个链表的头节点 `head` 和一个整数 `val` ，请你删除链表中所有满足 `Node.val == val` 的节点，并返回 **新的头节点** 。
 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/03/06/removelinked-list.jpg" style="width: 500px; height: 142px;" />
```

输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

```
 **示例 2：** 

```

输入：head = [], val = 1
输出：[]

```
 **示例 3：** 

```

输入：head = [7,7,7,7], val = 7
输出：[]

```
 

 **提示：** 
- 列表中的节点在范围 `[0, 10^4]` 内
-  `1 <= Node.val <= 50` 
-  `0 <= k <= 50` 
 
**标签**
`链表` 


## 
```go
func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := ListNode{Next: head}
	for preNode, node := &dummyNode, dummyNode.Next; node != nil; {
		for node != nil && node.Val == val {
			node = node.Next
		}
		preNode.Next = node
		preNode = preNode.Next
		if node != nil {
			node = node.Next
		}
	}
	return dummyNode.Next
}
```
>执行用时: 8 ms
内存消耗: 4.6 MB

要去除非法的元素就只需要记住在它之前的那个合法的节点，然后不断地往后挪就可以了。