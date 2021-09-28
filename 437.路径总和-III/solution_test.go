package leetcode

import "fmt"

// 437.路径总和III
// https://leetcode-cn.com/problems/path-sum-iii
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

//

type TreeNodeWithCount struct {
	Val   int
	Left  *TreeNodeWithCount
	Right *TreeNodeWithCount
	count map[int]int
}

var res = 0
var _targetSum = 0

func pathSum(root *TreeNode, targetSum int) int {
	res = 0
	_targetSum = targetSum
	ret := conv(root)
	pathSum2(ret)
	return res
}

func conv(root *TreeNode) *TreeNodeWithCount {
	if root == nil {
		return nil
	}

	ret := &TreeNodeWithCount{}
	ret.count = map[int]int{}

	ret.count[root.Val]++
	ret.Left = conv(root.Left)
	ret.Right = conv(root.Right)
	ret.Val = root.Val

	return ret
}

func pathSum2(root *TreeNodeWithCount) {
	if root == nil {
		return
	}
	pathSum2(root.Left)
	pathSum2(root.Right)
	res += root.count[_targetSum]
	add(root, root.Left)
	add(root, root.Right)
}

func add(root, son *TreeNodeWithCount) {
	if son != nil {
		for k, v := range son.count {
			root.count[k+root.Val] += v
			fmt.Println(root, k+root.Val)
			if k+root.Val == _targetSum {
				res += v
			}
		}
	}
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
