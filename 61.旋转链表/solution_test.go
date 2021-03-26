package leetcode

import (
	"reflect"
	"testing"
)

// 61.旋转链表
// https://leetcode-cn.com/problems/rotate-list
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

func Len(head *ListNode) (n int) {
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := Len(head)
	k = k % n
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
		if fast == nil {
			fast = head
		}
	}
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if fast == slow {
		return head
	}
	res := slow.Next
	slow.Next = nil
	fast.Next = head
	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		lists []int
		k     int
		want  []int
	}{
		{
			lists: []int{1, 2, 3, 4, 5},
			k:     2,
			want:  []int{4, 5, 1, 2, 3},
		},
		{
			lists: []int{0, 1, 2},
			k:     4,
			want:  []int{2, 0, 1},
		},
		{
			lists: []int{0},
			k:     20,
			want:  []int{0},
		},
		{
			lists: []int{1, 2},
			k:     2,
			want:  []int{1, 2},
		},
		{
			lists: []int{},
			k:     2,
			want:  []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lists := make([]ListNode, len(tC.lists))
			for i := range tC.lists {
				lists[i].Val = tC.lists[i]
				if i < len(tC.lists)-1 {
					lists[i].Next = &lists[i+1]
				}
			}
			var get *ListNode
			if len(lists) != 0 {
				get = rotateRight(&lists[0], tC.k)
			}
			resList := []int{}
			for get != nil {
				resList = append(resList, get.Val)
				get = get.Next
			}
			if !reflect.DeepEqual(tC.want, resList) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
