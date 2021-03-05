# 232.用栈实现队列
[https://leetcode-cn.com/problems/implement-queue-using-stacks](https://leetcode-cn.com/problems/implement-queue-using-stacks) 
## 原题
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列的支持的所有操作（ `push` 、 `pop` 、 `peek` 、 `empty` ）：

实现 `MyQueue` 类：
-  `void push(int x)` 将元素 x 推到队列的末尾
-  `int pop()` 从队列的开头移除并返回元素
-  `int peek()` 返回队列开头的元素
-  `boolean empty()` 如果队列为空，返回 `true` ；否则，返回 `false` 
 

 **说明：** 
- 你只能使用标准的栈操作 —— 也就是只有  `push to top` ,  `peek/pop from top` ,  `size` , 和  `is empty`  操作是合法的。
- 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 

 **进阶：** 
- 你能否实现每个操作均摊时间复杂度为 `O(1)` 的队列？换句话说，执行 `n` 个操作的总时间复杂度为 `O(n)` ，即使其中一个操作可能花费较长时间。
 

 **示例：** 

```

输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]

解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false

```
 

 **提示：** 
-  `1 <= x <= 9` 
- 最多调用 `100` 次 `push` 、 `pop` 、 `peek` 和 `empty` 
- 假设所有操作都是有效的 （例如，一个空的队列不会调用 `pop` 或者 `peek` 操作）
 
**标签**
`栈` `设计` 


## 
```go
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
```
>执行用时: 0 ms
内存消耗: 1.9 MB

用额外的一个栈来辅助进行倒序，即把栈中的元素变为队列的顺序。

我这里用 stackA 来作为正常栈，用 stackB 来辅助 Pop 时候的操作

如果操作数是平均的话，那么选任意一个操作作为 $O(1)$ 的时间复杂度即可。不然的话就需要对操作数多的那个操作进行优化，让它的时间复杂度降低。

于是这里多用了一个 top 来标记队列的第一个元素。当需要第一个元素的时候直接返回即可。这样就可以让 Peek 的时间复杂度降低。如果要 Push 的时间复杂度低的话 就让 Pop 和 Peek 的时候进行调节即可。同理要 Pop 的时间复杂度低的话 就让 Push 的时候就完成倒序即可。
