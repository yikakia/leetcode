package leetcode

import "testing"

// 232.用栈实现队列
// https://leetcode-cn.com/problems/implement-queue-using-stacks
type MyQueue struct {
	stackA []int
	stackB []int
	top    int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{stackA: []int{}, stackB: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if len(this.stackA) == 0 {
		if len(this.stackB) == 0 {
			this.top = x
		} else {
			this.top = this.stackB[len(this.stackB)-1]
			for i := len(this.stackB) - 1; i >= 0; i-- {
				this.stackA = append(this.stackA, this.stackB[i])
			}
			this.stackB = this.stackB[:0]
		}

	}
	this.stackA = append(this.stackA, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stackB) == 0 {
		for i := len(this.stackA) - 1; i >= 0; i-- {
			this.stackB = append(this.stackB, this.stackA[i])
		}
		this.stackA = this.stackA[:0]
	}
	res := this.stackB[len(this.stackB)-1]
	this.stackB = this.stackB[:len(this.stackB)-1]
	if len(this.stackB) > 0 {
		this.top = this.stackB[len(this.stackB)-1]
	}
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.top
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stackA) == 0 && len(this.stackB) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func TestSolution(t *testing.T) {
	t.Run("", func(t *testing.T) {
		q := Constructor()
		for i := 0; i < 5; i++ {
			q.Push(i)
			if q.Peek() != 0 {
				t.Errorf("push : %d get top %d\n", i, q.Peek())
			}
		}
		for i := 5; i < 10; i++ {
			r := q.Pop()
			if r != i-5 {
				t.Errorf("pop want : %d get %d\n", i-5, r)
			}
		}

	})

}
