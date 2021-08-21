package leetcode

// 116.填充每个节点的下一个右侧节点指针
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	bros := []*Node{root, nil}
	for len(bros) > 0 {
		n := len(bros)
		used := false
		for i := 0; i < n-1; i++ {
			if bros[i] == nil {
				continue
			}
			used = true
			bros[i].Next = bros[i+1]
			bros = append(bros, bros[i].Left, bros[i].Right)
		}
		if !used {
			break
		}
		bros = append(bros, nil)
		bros = bros[n:]
	}
	return root
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
