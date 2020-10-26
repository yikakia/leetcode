# 剑指 Offer 28. 对称的二叉树
[https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/](https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/)

## 中序遍历后对比
```go

func isSymmetric(root *TreeNode) bool {
	r := []int{}
	r = dfs(root, r, 0)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-i-1] {
			return false
		}
	}
	return true
}
func dfs(root *TreeNode, r []int, depth int) []int {
	if root == nil {
		r = append(r, 0xfffffff-depth)
		return r
	}
	if root.Left == nil && root.Right == nil {
		r = append(r, root.Val)
		return r
	}

	r1 := append(r, dfs(root.Left, r, depth+1)...)
	r2 := append(r, dfs(root.Right, r, depth+1)...)

	r = append(r1, r...)
	r = append(r, root.Val)
	r = append(r, r2...)
	return r
}

```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：4.4 MB, 在所有 Go 提交中击败了5.19%的用户

很简单，就是根据中序遍历的顺序进行判断。但是单纯的中序遍历得到的数据会被一个精心设计的坑 `[5,4,1,null,1,null,4,2,null,2,null]`卡住

按照 `leetcode`的表示结果，这个中序遍历的结果竟然是回文的！！！

当时就惊呆了，于是琢磨了半天，把深度信息作为噪声加了进去进行比较。不得不说这个卡的精准。还有就是这个内存实在是不忍直视，占用太多了。主要是不断地append，做了不少重复的操作。

看了题解才发现可以递归直接比较。

## 递归直接比较
```go
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsSymmetric(root.Left, root.Right)
}
func dfsSymmetric(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) || (a != nil && b == nil) || (a.Val != b.Val) {
		return false
	}
	return dfsSymmetric(a.Left, b.Right) && dfsSymmetric(a.Right, b.Left)
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了74.01%的用户
内存消耗：2.9 MB, 在所有 Go 提交中击败了62.86%的用户

简单地说就是另外写一个函数用于对比当前的两个树是否镜像。每次就是把左节点的左子树和右节点的右子树对比，左节点的右子树和右节点的左子树对比。

