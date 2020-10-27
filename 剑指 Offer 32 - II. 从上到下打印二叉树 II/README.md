# 剑指 Offer 32 - II. 从上到下打印二叉树 II
[https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/)

## 
```go

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	r := [][]int{}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		width := len(queue)
		tmpr := make([]int, width)
		for i := 0; i < width; i++ {
			tmpnode := queue[0]
			tmpr[i] = tmpnode.Val
			queue = queue[1:]
			if tmpnode.Left != nil {
				queue = append(queue, tmpnode.Left)
			}
			if tmpnode.Right != nil {
				queue = append(queue, tmpnode.Right)
			}
		}
		r = append(r, tmpr)
	}
	return r
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.8 MB, 在所有 Go 提交中击败了66.67%的用户

和之前那道题差距不大，就是改成一次把当前的全部入了就行。