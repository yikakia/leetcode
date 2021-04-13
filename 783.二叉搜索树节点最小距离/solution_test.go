package leetcode

// 783.二叉搜索树节点最小距离
// https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if res > num {
			res = num
		}
	}
	return res
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 1e6
	}
	res := min(minDiffInBST(root.Left), minDiffInBST(root.Right))
	lMax := maxLeft(root)
	if lMax != nil {
		res = min(res, root.Val-lMax.Val)
	}
	rMin := minRight(root)
	if rMin != nil {
		res = min(res, rMin.Val-root.Val)
	}
	return res
}

func maxLeft(root *TreeNode) *TreeNode {
	cur := root.Left
	if cur == nil {
		return nil
	}
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

func minRight(root *TreeNode) *TreeNode {
	cur := root.Right
	if cur == nil {
		return nil
	}
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
