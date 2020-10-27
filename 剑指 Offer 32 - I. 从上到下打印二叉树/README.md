# 剑指 Offer 32 - I. 从上到下打印二叉树
[https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/)

## 
```go
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0)
	r := make([]int, 0)

	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		r = append(r, tmp.Val)
		if tmp.Left != nil {
			queue = append(queue, tmp.Left)
		}
		if tmp.Right != nil {
			queue = append(queue, tmp.Right)
		}
	}
	return r
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了8.99%的用户
内存消耗：2.7 MB, 在所有 Go 提交中击败了47.86%的用户

看了下好像都是用的队列，不知道为什么耗时这么多。

