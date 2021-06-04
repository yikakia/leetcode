package leetcode

import (
	"reflect"
	"testing"
)

// 160.相交链表
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
	Val int
	Next *ListNode
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    la := len(headA)
    lb := len(headB)
    pa,pb := headA,headB
    pa = nextN(pa,la-lb)
    pb = nextN(pb,lb-la)
    for pa!=nil{
        if pa == pb{
            return pa
        }
        pa=pa.Next
        pb=pb.Next
    }
    return nil
}

func len(list *ListNode) int{
    if list == nil{
        return 0
    }
    n := 0
    for list != nil{
        list= list.Next
        n++
    }
    return n
}

func nextN(list *ListNode,n int) *ListNode{
    if n <= 0{
        return list
    }
    for n >0{
        list = list.Next
        n--
    }
    return list
}
func TestSolution(t *testing.T) {
	testCases := []struct {
        
		want 
		desc string        
	}{
		{
            want: ,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := 
			if !reflect.DeepEqual(tC.want,get){
				t.Errorf("input: %+v get: %v\n",tC,get)
			}
		})
	}
}
