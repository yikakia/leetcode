# 1721.交换链表中的节点
[https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list/](https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list/) 
## 原题
给你链表的头节点 `head` 和一个整数 `k` 。

**交换**  链表正数第 `k` 个节点和倒数第 `k` 个节点的值后，返回链表的头节点（链表 **从 1 开始索引** ）。

 

**示例 1：** 
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/01/10/linked1.jpg" style="width: 722px; height: 202px;" />
```

输入：head = [1,2,3,4,5], k = 2
输出：[1,4,3,2,5]

```
**示例 2：** 

```

输入：head = [7,9,6,6,7,8,3,0,9,5], k = 5
输出：[7,9,6,6,8,7,3,0,9,5]

```
**示例 3：** 

```

输入：head = [1], k = 1
输出：[1]

```
**示例 4：** 

```

输入：head = [1,2], k = 1
输出：[2,1]

```
**示例 5：** 

```

输入：head = [1,2,3], k = 2
输出：[1,2,3]

```
 

**提示：** 
- 链表中节点的数目是 `n`
- `1 <= k <= n <= 10^5`
- `0 <= Node.val <= 100`
 
**标签**
`链表` 


## 

```go
func lenth(head *ListNode) (n int) {
	for head != nil {
		n++
		head = head.Next
	}
	return
}

func getKAndFather(head *ListNode, k int) (*ListNode, *ListNode) {
	if k < 1 {
		return nil, nil
	}
	if k == 1 {
		return head, nil
	}
	kFNode := head
	for i := 1; i < k-1; i++ {
		kFNode = kFNode.Next
	}
	kNode := kFNode.Next
	return kNode, kFNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	n := lenth(head)
	if n == 1 {
		return head
	}
	if k > n/2 {
		k = n - k + 1
	}
	dum := &ListNode{}
	dum.Next = head

	kNode, kFNode := getKAndFather(head, k)
	nkNode, nkFnode := getKAndFather(head, n-k+1)

	if kFNode == nil {
		kFNode = dum
	}
	kFNode.Next = nkNode
	nkFnode.Next = kNode
	kNode.Next, nkNode.Next = nkNode.Next, kNode.Next
	return dum.Next
}
```
>执行用时: 208 ms
内存消耗: 13.2 MB

遍历两次。用一个辅助函数来返回第k个节点和它的父节点。
不过看了一下才发现原来只用交换数值就可以了。早知道就不用断链了。

## 双指针
还有种方法是双指针，简单地说就是先找到第k个节点，并且记录下来。然后再开一个指针，从头开始，两个指针一起向后遍历，直到到头为止。这个时候交换两者的值即可。