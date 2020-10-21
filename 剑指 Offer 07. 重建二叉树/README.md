# 剑指 Offer 07. 重建二叉树
[https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)

## 原理分析
```go
func buildTree(preorder []int, inorder []int) *TreeNode {
	for root := range inorder {
		if inorder[root] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:root+1], inorder[0:root]),
				Right: buildTree(preorder[root+1:], inorder[root+1:]),
			}
		}
	}
	return nil
}
```
>执行用时：8 ms, 在所有 Go 提交中击败了59.37%的用户
内存消耗：4.2 MB, 在所有 Go 提交中击败了42.06%的用户

简单地说就是利用二叉树三种遍历的性质。
1. 先序遍历的第一个元素一定是根节点
2. 先序遍历后面跟的[1,root]的节点为左子树的节点
3. 先序遍历后面跟的[root+1,lenth]的节点为右子树的节点
4. 中序遍历的根节点左右一定分别是左子树的根节点和右子树的根节点

>注： `root` 为根节点在中序遍历中的位置,`lenth`为二叉树节点的个数

那么在已知先序遍历和中序遍历的情况下就能够通过这些性质来一步步重建二叉树。
具体的步骤是确定根节点，对中序遍历和先序遍历的结果进行切片，分别寻找左子树和右子树的根节点。如此迭代，最后得到的便是
这里也用到了递归的思想。

那么后序遍历有什么特点呢？
1. 后序遍历的最后一个节点一定是根节点。
2. 后序遍历的[0,root]的节点为左子树的节点
3. 后序遍历的[root,lenth-1]的节点为右子树的节点

根据这个我们就可以的得到一个结论，要重建二叉树，中序遍历是必须的，因为可以通过它来确定左右子树的长度。而先序遍历和后序遍历可以任选其一，因为它们只是用来确定根节点的位置的。