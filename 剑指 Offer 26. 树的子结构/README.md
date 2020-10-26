# 剑指 Offer 26. 树的子结构
[https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/)

##

```go

// 查找 B 树是否是 A 树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val {
		if dfs(A, B) {
			return true
		}
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// dfs 证明 A 树与 B 树结构相同。
func dfs(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
}
```
>执行用时：32 ms, 在所有 Go 提交中击败了12.80%的用户
内存消耗：7.1 MB, 在所有 Go 提交中击败了41.69%的用户

最开始是打算单递归的，但是发现有问题。主要是要判断
1. A B 结构是否相等
2. A 的子结构和 B 是否相等

这两个的判断条件并不相同。因为证明结构相同的时候跳出的条件之一是 B 树可以在遍历的过程中变为空节点，而证明 B 是 A 的子结构时，如果进入的时候 B 是空的，那就不能认为这个是对的（根据题意，空树不为子结构）。

因此就应该设置一个双递归，一个用于证明 A B 的结构相同，一个证明 A 的子结构和 B 的结构相同