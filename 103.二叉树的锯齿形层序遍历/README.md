# 103. 二叉树的锯齿形层序遍历
[https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/) 
## 原题
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
例如：<br />
给定二叉树 `[3,9,20,null,null,15,7]`,
```
    3
   / \
  9  20
    /  \
   15   7
```
返回锯齿形层序遍历如下：
```
[
  [3],
  [20,9],
  [15,7]
]
```


## 
```go
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		width := len(queue)
		tmpRes := make([]int, width)
		if depth%2 == 0 {
			for i := 0; i < width; i++ {
				tmpnode := queue[0]
				tmpRes[i] = tmpnode.Val
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
				tmpRes[i-1] = tmpnode.Val
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
		res = append(res, tmpRes)
	}
	return res
}
```
>执行用时: 0 ms
内存消耗: 2.5 MB

就和广搜差不多，不过要看深度决定怎么放入结果