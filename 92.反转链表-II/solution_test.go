package leetcode

import "testing"

// 92.反转链表II
// https://leetcode-cn.com/problems/reverse-linked-list-ii
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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		list        []int
		left, right int
	}{
		{
			list:  []int{1, 2, 3, 4, 5},
			left:  2,
			right: 4,
		},
	}
	for _, tC := range testCases {
		lists := make([]ListNode, len(tC.list))
		for i := range lists {
			lists[i].Val = tC.list[i]
			if i != len(tC.list)-1 {
				lists[i].Next = &lists[i+1]
			}
		}
		reverseBetween(&lists[0], tC.left, tC.right)
	}
}
