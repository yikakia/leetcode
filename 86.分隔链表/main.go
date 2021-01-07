package main

import "fmt"

// 86.分隔链表
// https://leetcode-cn.com/problems/partition-list/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	smallHead := ListNode{}
	bigHead := ListNode{}
	tmp := head
	small := &smallHead
	big := &bigHead
	for tmp != nil {
		if tmp.Val < x {
			small.Next = tmp
			small = small.Next
		} else {
			big.Next = tmp
			big = big.Next
		}
		tmp = tmp.Next
	}
	small.Next = bigHead.Next
	big.Next = nil
	return smallHead.Next
}

func test() {
	vals := [][]int{
		[]int{1, 4, 3, 2, 5, 2},
	}
	lists := make([][]ListNode, len(vals))

	for i := range vals {
		lists[i] = make([]ListNode, len(vals[i]))
		for j := range vals[i] {
			lists[i][j].Val = vals[i][j]
		}
		for j := 0; j < len(vals[i])-1; j++ {
			lists[i][j].Next = &lists[i][j+1]
		}
	}

	xs := []int{3}
	for i := range lists {
		get := partition(&lists[i][0], xs[i])
		for get != nil {
			fmt.Print(get)
			get = get.Next
		}
	}

	fmt.Println("Finished Test!")
}
func main() {
	test()

}
