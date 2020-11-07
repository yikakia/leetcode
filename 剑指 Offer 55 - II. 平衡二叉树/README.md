# 剑指 Offer 55 - II. 平衡二叉树
[https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/](https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/)
## 后序遍历
```go
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	r, _ := dfs(root)
	return r
}
func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return false, 0
	}
	_, depth := dfs(root.Left)

	if depth == -1 {
		return false, -1
	}
	_, tmpDepth := dfs(root.Right)
	if tmpDepth == -1 {
		return false, -1
	}
	if math.Abs(float64(depth-tmpDepth)) > 1 {
		return false, -1
	}
	return true, max(depth, tmpDepth) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了99.63%的用户
内存消耗：6.1 MB, 在所有 Go 提交中击败了24.95%的用户

简单的就是后序遍历。不过要特别判断当根节点为空时应该直接输出 `true`