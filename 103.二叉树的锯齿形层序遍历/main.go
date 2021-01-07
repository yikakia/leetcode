package main

// 103.二叉树的锯齿形层序遍历
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

// func test() {
// 	type TestType struct {
// 	}
// 	ts := []TestType{
// 		TestType{},
// 	}
// 	for _, t := range ts {
// 		get := zigzagLevelOrder()
// 		if get != t.want {
// 			// 填充输出格式
// 			fmt.Printf("%+v get:%v\n", t, get)
// 		}
// 	}
// 	fmt.Println("Finished Test!")

// }
func main() {
	// test()
}
