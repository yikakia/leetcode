package leetcode

// 83.删除排序链表中的重复元素
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	for cur := head; cur != nil; {
		if next := cur.Next; next != nil && next.Val == cur.Val {
			for next = cur.Next; next != nil && next.Val == cur.Val; {
				next = next.Next
			}
			cur.Next = next
			cur = cur.Next
		} else {
			cur = cur.Next
		}

	}
	return head
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
