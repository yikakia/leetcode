# 938.二叉搜索树的范围和
[https://leetcode-cn.com/problems/range-sum-of-bst](https://leetcode-cn.com/problems/range-sum-of-bst) 
## 原题
给定二叉搜索树的根结点  `root` ，返回值位于范围 *`[low, high]`* 之间的所有结点的值的和。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/05/bst1.jpg" style="width: 400px; height: 222px;" />
```

输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/05/bst2.jpg" style="width: 400px; height: 335px;" />
```

输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
输出：23

```
 

 **提示：** 
- 树中节点数目在范围 `[1, 2 * 10^4]` 内
-  `1 <= Node.val <= 10^5` 
-  `1 <= low <= high <= 10^5` 
- 所有 `Node.val` **互不相同** 
 
**标签**
`树` `深度优先搜索` `递归` 


## 标准二分搜索树的应用
```go
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	if root.Val <= high {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
	return rangeSumBST(root.Left, low, high)
}
```
>执行用时: 104 ms
内存消耗: 9.1 MB

递归累加对应区间的值即可。