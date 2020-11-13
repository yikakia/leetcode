# 剑指 Offer 34. 二叉树中和为某一值的路径
[https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/](https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/)

## dfs
```go

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	r := make([][]int, 0)
	dfs(root, sum, []int{}, &r)
	return r
}

func dfs(root *TreeNode, sum int, tmpr []int, r *[][]int) {
	if root == nil {
		return
	}
	tmpr = append(tmpr, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(tmpr))
		copy(tmp, tmpr)
		*r = append(*r, tmp)
	}
	dfs(root.Left, sum-root.Val, tmpr, r)
	dfs(root.Right, sum-root.Val, tmpr, r)
	tmpr = tmpr[:len(tmpr)-1]
}

```

>执行用时：4 ms, 在所有 Go 提交中击败了91.67%的用户
内存消耗：4.7 MB, 在所有 Go 提交中击败了39.71%的用户

就是简单的 `dfs` 不过需要记录当前的遍历序列，还得插入以后不被修改。所以要用一个`copy`，还要一个`tmpr = tmpr[:len(tmpr)-1]`来便于回溯。因为题目只说了整数没说是正整数，所以即使当前满足了，也得继续找下去，直到找到叶子节点为止。因为可能来一个`[-1,+1]`就把你的抵消了。😀