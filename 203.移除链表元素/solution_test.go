package leetcode

import (
	"fmt"
	"testing"
)

// 203.移除链表元素
// https://leetcode-cn.com/problems/remove-linked-list-elements
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

func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := ListNode{Next: head}
	for preNode, node := &dummyNode, dummyNode.Next; node != nil; {
		for node != nil && node.Val == val {
			node = node.Next
		}
		preNode.Next = node
		preNode = preNode.Next
		if node != nil {
			node = node.Next
		}
	}
	return dummyNode.Next
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		head *ListNode
		vals []int
		val  int
		desc string
	}{
		{
			vals: []int{1, 2, 36, 3, 4, 5, 6},
			val:  6,
			// want: ,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			nodes := make([]ListNode, len(tC.vals))
			for i, val := range tC.vals {
				nodes[i].Val = val
			}
			for i := range tC.vals[1:] {
				nodes[i].Next = &nodes[i+1]
			}
			tC.head = &nodes[0]
			get := removeElements(tC.head, tC.val)
			for get != nil {
				fmt.Print(get, " ")
			}
			fmt.Println()
			// if !reflect.DeepEqual(tC.want,get){
			// 	t.Errorf("input: %+v get: %v\n",tC,get)
			// }
		})
	}
}
