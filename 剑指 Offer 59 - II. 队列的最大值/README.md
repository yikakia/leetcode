# 剑指 Offer 59 - II. 队列的最大值
[https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/](https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/)
## 双向队列
```go
type MaxQueue struct {
	q    []int
	maxq []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxq) == 0 {
		return -1
	}
	return this.maxq[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	this.maxq = adjQueue(this.maxq, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	if this.q[0] == this.maxq[0] {
		this.maxq = this.maxq[1:]
	}
	r := this.q[0]
	this.q = this.q[1:]
	return r
}

func adjQueue(queue []int, x int) []int {
	n := len(queue)
	res := make([]int, 0)
	if n == 0 {
		res = append(res, x)
	}
	for i := n - 1; i >= 0; i-- {
		if x <= queue[i] {
			res = append(queue[:i+1], x)
			break
		} else {
			res = append(queue[:i], x)
		}
	}
	return res
}
```
>执行用时：100 ms, 在所有 Go 提交中击败了89.40%的用户
内存消耗：8.6 MB, 在所有 Go 提交中击败了17.05%的用户

和之前的那道题[剑指 Offer 59 - I. 滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)差不多。只不过这里是作为一个队列来辅助的。