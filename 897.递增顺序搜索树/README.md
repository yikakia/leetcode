# 897.递增顺序搜索树
[https://leetcode-cn.com/problems/increasing-order-search-tree](https://leetcode-cn.com/problems/increasing-order-search-tree) 
## 原题
给你一棵二叉搜索树，请你 **按中序遍历** 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/17/ex1.jpg" style="width: 600px; height: 350px;" />
```

输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/17/ex2.jpg" style="width: 300px; height: 114px;" />
```

输入：root = [5,1,7]
输出：[1,null,5,null,7]

```
 

 **提示：** 
- 树中节点数的取值范围是 `[1, 100]` 
-  `0 <= Node.val <= 1000` 
 
**标签**
`树` `深度优先搜索` `递归` 


## 
```go
func increasingBST(root *TreeNode) *TreeNode {
	dummy := TreeNode{}
	pre := &dummy

	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		pre.Right = tn
		tn.Left = nil
		pre = tn

		dfs(tn.Right)
	}
	dfs(root)
	return dummy.Right
}

```
>执行用时: 0 ms
内存消耗: 2.2 MB

最重要的就是设定一个 `pre` 记录上一个节点是什么。这里的 `pre` 就是它的左子树的最右。这个时候把它们相连即可。
