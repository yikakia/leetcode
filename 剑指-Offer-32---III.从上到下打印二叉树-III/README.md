# 剑指 Offer 32 - III. 从上到下打印二叉树 III
[https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/)

## 顺序入队，根据深度切换放入数组的方式。
```go 
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	r := [][]int{}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		width := len(queue)
		tmpr := make([]int, width)
		if depth%2 == 0 {
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
			depth++
		} else {
			for i := width; i > 0; i-- {
				tmpnode := queue[0]
				tmpr[i-1] = tmpnode.Val
				queue = queue[1:]
				if tmpnode.Left != nil {
					queue = append(queue, tmpnode.Left)
				}
				if tmpnode.Right != nil {
					queue = append(queue, tmpnode.Right)
				}
			}
			depth++
		}
		r = append(r, tmpr)
	}
	return r
}
```

>执行用时：4 ms, 在所有 Go 提交中击败了7.69%的用户
内存消耗：2.8 MB, 在所有 Go 提交中击败了72.93%的用户

和上一题差距不大，就是改了个根据奇偶决定放入数组 `tmpr` 的方式。不过这个时间有点丢人，我看看题解。

好像思路都一样啊，有个标识符用来决定是不是要翻转数组。不过我是直接改变构建数组的方式，别人是构建完数组后进行翻转。