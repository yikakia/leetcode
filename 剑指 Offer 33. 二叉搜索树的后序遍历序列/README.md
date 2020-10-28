# 剑指 Offer 33. 二叉搜索树的后序遍历序列
[https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/)

## 循环验证
```go
func verifyPostorder(postorder []int) bool {
	lenth := len(postorder)
	if lenth < 3 {
		return true
	}
	rootIndex := lenth - 1
	for rootIndex != 0 {
		ptr := 0

		for postorder[ptr] < postorder[rootIndex] {
			ptr++
		}

		for postorder[ptr] > postorder[rootIndex] {
			ptr++
		}

		if ptr != rootIndex {
			return false
		}
		rootIndex--
	}
	return true
}
```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了88.17%的用户

- 二叉线索树的特征：
    - 左子树的所有节点一定小于根节点
    - 右子树的所有节点一定大于根节点

- 后序遍历的特征：
    - 最后一个元素一定是根节点
    - 第一个元素一定是最左边的孩子
    - 先遍历完左子树，然后再遍历右子树

简单地说就是从后序遍历的性质来看，最后一个必然是根节点。那么可以知道的是，它的左半部分一定是有一个递增序列和递减序列的。（由二叉线索树的性质和后序遍历的性质决定）。而对于数组中的每一个元素而言，它如果是二叉线索树的后序遍历顺序，那么也一定满足这个条件，即它的左部是一个递增序列和一个递减序列。（虽然递增序列的一部分是父节点的左子树）

那么解法就很明确了，从后往前，对每一个元素检查是否满足二叉线索树地后序遍历。
