# 230.二叉搜索树中第K小的元素
[https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst) 
## 原题
给定一个二叉搜索树的根节点 `root` ，和一个整数 `k` ，请你设计一个算法查找其中第  `k`  个最小元素（从 1 开始计数）。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg" style="width: 212px; height: 301px;" />
```

输入：root = [3,1,4,null,2], k = 1
输出：1

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg" style="width: 382px; height: 302px;" />
```

输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3

```
 

 

 **提示：** 
- 树中的节点数为 `n` 。
-  `1 <= k <= n <= 10^4` 
-  `0 <= Node.val <= 10^4` 
 

 **进阶：** 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 `k` 小的值，你将如何优化算法？

 
**标签**
`树` `深度优先搜索` `二叉搜索树` `二叉树` 


## 二叉搜索树的方式来查
```go
var cur int

func kthSmallest(root *TreeNode, k int) int {
	cur = 0
	return dokthSmallest(root, k).Val
}

func dokthSmallest(root *TreeNode, k int) *TreeNode {
	if root == nil {
		return nil
	}

	if ans := dokthSmallest(root.Left, k); ans != nil {
		return ans
	}
	cur++
	if cur == k {
		return root
	}
	if ans := dokthSmallest(root.Right, k); ans != nil {
		return ans
	}

	return nil
}
```
>执行用时：8 ms
内存消耗：6.4 MB

全局变量记录当前是第几个元素，然后每次都更新这个变量即可。当当前节点是在更新之后就是对应的第k个元素的话，那么就直接返回即可。