package leetcode

import (
	"fmt"
	"reflect"
)

// 872.叶子相似的树
// https://leetcode-cn.com/problems/leaf-similar-trees
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

func getLeafSlice(root *TreeNode) []int {
	resSlice := []int{}
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		if node.Left == nil && node.Right == nil {
			resSlice = append(resSlice, node.Val)
		}
		helper(node.Right)
	}
	helper(root)
	return resSlice
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1, s2 := getLeafSlice(root1), getLeafSlice(root2)
	fmt.Println(s1, s2)
	return reflect.DeepEqual(s1, s2)
}

// func TestSolution(t *testing.T) {
// 	testCases := []struct {
// 		desc string
// 		want
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
