# 剑指 Offer 30. 包含min函数的栈
[https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/](https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/)

## 每次 pop 都判断一次

```go
type MinStack struct {
	nums []int
	min  int
	len  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.nums) == 0 {
		this.min = x
	}
	if x < this.min {
		this.min = x
	}
	this.nums = append(this.nums, x)
	this.len++
}

func (this *MinStack) Pop() {
	if this.len == 0 {
		return
	}
	this.nums = this.nums[:this.len-1]
	this.len--
	if this.len == 0 {
		return
	}
	min := this.nums[0]
	for i := 0; i < this.len; i++ {
		if min > this.nums[i] {
			min = this.nums[i]
		}
	}
	this.min = min
}

func (this *MinStack) Top() int {
	return this.nums[this.len-1]
}

func (this *MinStack) Min() int {
	return this.min
}

```

>执行用时：16 ms, 在所有 Go 提交中击败了92.06%的用户
内存消耗：8.5 MB, 在所有 Go 提交中击败了18.38%的用户

虽然实现了，不过体会一下解决方案，好像只是利用的`Golang`封装好的特性？就是非传统的栈的形式。因为这个明显是要去找栈的最小元素，是不可能遍历的。

那么对于栈而言，最好的解决就是辅助栈了。

## 辅助栈

简单地说就是另外开一个栈，另外一个栈记录的是当前栈顶元素之前的最小元素。换句话说把数组看成起伏的大山，那么辅助栈就是记录这座山之前的最低的山脚。因为两个栈的数量是肯定一样的，所以这个是可以明确一一对应的。
```go
type MinStack struct {
	nums []int
	min  []int
	len  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nums: make([]int, 0), min: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if this.len == 0 || x < this.min[this.len-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[this.len-1])
	}

	this.nums = append(this.nums, x)
	this.len++
}

func (this *MinStack) Pop() {
	if this.len == 0 {
		return
	}
	this.nums = this.nums[:this.len-1]
	this.min = this.min[:this.len-1]
	this.len--
	}

func (this *MinStack) Top() int {
	if this.len <= 0 {
		return 0
	}
	return this.nums[this.len-1]
}

func (this *MinStack) Min() int {
	if this.len <= 0 {
		return 0
	}
	return this.min[this.len-1]
}
```

>执行用时：20 ms, 在所有 Go 提交中击败了73.46%的用户
内存消耗：8.5 MB, 在所有 Go 提交中击败了26.37%的用户

内存消耗的确小了一丢丢，不过时间为啥消耗更多了。看了下也没啥大的区别，很奇怪。