# 剑指 Offer 27. 二叉树的镜像
[https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/)

## 递归
```go
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了60.94%的用户

就是简单的递归，然后把无脑把左右子树交换。