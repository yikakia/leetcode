# 剑指 Offer 54. 二叉搜索树的第k大节点
[https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/)
## 魔改中序遍历
```go
func kthLargest(root *TreeNode, k int) int {
	res, _ := dfs(root, k)
	return res
}

// 返回值 value 表示要找的数，返回值 nums 表示这个节点的子树的个数
// 当返回值 nums 为 -1 时表示找到了，value 就是要找的数
//
func dfs(root *TreeNode, k int) (value int, nums int) {
	// 到了空节点，没有子树，返回 nums 值为 0
	if root == nil {
		return -1, 0
	}
	value, nums = dfs(root.Right, k)
	// 右子树已经找到了要找的了，返回值为 value
	if nums == -1 {
		return value, -1
	}

	// 该点就是要找的
	if k == nums+1 {
		return root.Val, -1
	}
	nums++
	// 计算左子树，这个时候要减去右子树和根节点的个数
	value, tmpNums := dfs(root.Left, k-nums)
	// 左子树找到了要找的
	if tmpNums == -1 {
		return value, -1
	}
	// 添加上左子树的节点数
	nums += tmpNums
	return -1, nums
}
```
>执行用时：12 ms, 在所有 Go 提交中击败了81.45%的用户
内存消耗：6.3 MB, 在所有 Go 提交中击败了43.11%的用户

思路蛮简单的，就是先访问右子树，再访问根节点，再访问左子树。

其实做的比较取巧，因为利用了多返回值的特性。通过 `nums` 的值同时来计算子树的节点个数和是否找到了要找的元素。不然就要另外维护一个值来记录遍历了多少个节点。其实也是可以的，不过不太想用全局变量。