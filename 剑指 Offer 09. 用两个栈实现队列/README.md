# 剑指 Offer 09. 用两个栈实现队列
[https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

## 手动实现栈，用栈来实现队列
```go
type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list: list}
}
func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

type CQueue struct {
	pushlist Stack
	poplist  Stack
}

func Constructor() CQueue {
	poplist := NewStack()
	pushlist := NewStack()
	return CQueue{
		pushlist: *pushlist,
		poplist:  *poplist,
	}
}

func (this *CQueue) AppendTail(value int) {
	for !this.poplist.Empty() {
		this.pushlist.Push(this.poplist.Pop())
	}
	this.pushlist.Push(value)
}

func (this *CQueue) DeleteHead() int {
	for !this.pushlist.Empty() {
		this.poplist.Push(this.pushlist.Pop())
	}
	r := this.poplist.Pop()
	switch value := r.(type) {
	default:
		return 0
	case nil:
		return -1
	case int:
		return value
	}
}
```
>执行用时：812 ms, 在所有 Go 提交中击败了5.36%的用户
内存消耗：8.7 MB, 在所有 Go 提交中击败了12.31%的用户

主要是因为golang本身没有提供一个真正意义上的 push pop 的栈，而是list这样支持多种操作的。所以我手写了一个栈来实现栈的操作。不过好像时间消耗很多？看了下别人的题解好像是直接用的 `int` 数组作为栈了，没有另外封装一个栈，所以效率很高。不过就思路而言是没啥问题的。就是拿两个栈出来，一个用于出队，一个用于入队。

入队的时候把出队的栈的全部元素挨个压入用于入队的栈，然后再把要入队的元素压入用于入队的栈

出队的时候把入队的栈的全部元素挨个压入用于出队的栈，然后再把用于出队的栈进行出栈操作。