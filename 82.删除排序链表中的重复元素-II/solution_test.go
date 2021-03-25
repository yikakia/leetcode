package leetcode

import (
	"testing"
)

// 82.删除排序链表中的重复元素II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
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
	dummy := &ListNode{Next: head}

	for cur, pre := dummy.Next, dummy; cur != nil; {
		if next := cur.Next; next != nil && next.Val == pre.Next.Val {
			for next != nil && next.Val == pre.Next.Val {
				cur = cur.Next
				next = next.Next
			}
			cur = next
			pre.Next = cur
			continue
		}
		cur = cur.Next
		pre = pre.Next
	}
	return dummy.Next
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		lists []int
		// want
	}{
		{
			// want: ,
			lists: []int{1, 2, 3, 3, 4, 4, 5},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lists := make([]ListNode, len(tC.lists))
			for i := range tC.lists {
				lists[i].Val = tC.lists[i]
			}
			for i := 0; i < len(lists)-1; i++ {
				lists[i].Next = &lists[i+1]
			}
			deleteDuplicates(&lists[0])
			// get :=
			// if !reflect.DeepEqual(tC.want,get){
			// 	t.Errorf("input: %+v get: %v\n",tC,get)
			// }
		})
	}
}
