package leetcode

// 230.二叉搜索树中第K小的元素
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cur int

func kthSmallest(root *TreeNode, k int) int {
	cur = 0
	return dokthSmallest(root, k).Val
}

func dokthSmallest(root *TreeNode, k int) *TreeNode {
	if root == nil {
		return nil
	}

	if ans := dokthSmallest(root.Left, k); ans != nil {
		return ans
	}
	cur++
	if cur == k {
		return root
	}
	if ans := dokthSmallest(root.Right, k); ans != nil {
		return ans
	}

	return nil
}

// func TestSolution(t *testing.T) {
// 	testCases := []struct {

// 		want
// 		desc string
// 	}{
// 		{
//             want: ,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			get :=
// 			if !reflect.DeepEqual(tC.want,get){
// 				t.Errorf("input: %+v get: %v\n",tC,get)
// 			}
// 		})
// 	}
// }
