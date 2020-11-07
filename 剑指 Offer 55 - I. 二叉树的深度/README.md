# 剑指 Offer 55 - I. 二叉树的深度
[https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/](https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/)
## 递归查找
```go
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}
```
>执行用时：8 ms, 在所有 Go 提交中击败了12.03%的用户
内存消耗：4.5 MB, 在所有 Go 提交中击败了19.00%的用户

就是递归左右子树，然后选择大的那个返回就行了。还有种不递归的是用栈实现的。