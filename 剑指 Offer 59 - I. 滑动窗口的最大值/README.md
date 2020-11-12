# 剑指 Offer 59 - I. 滑动窗口的最大值
[https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)
## 维护一个队列
```go
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n-k+1)
	queue := make([]int, 1)
	for i := 0; i < k; i++ {
		queue = adjQueue(queue, nums[i])
	}
	for i := k; i < n; i++ {
		res[i-k] = queue[0]
		if queue[0] == nums[i-k] {
			queue = adjQueue(queue[1:], nums[i])
		} else {
			queue = adjQueue(queue, nums[i])
		}

	}
	res[n-k] = queue[0]
	return res
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
>执行用时：16 ms, 在所有 Go 提交中击败了95.70%的用户
内存消耗：6.3 MB, 在所有 Go 提交中击败了66.03%的用户

简单地说就是维护一个双向队列，让新加入的值插入到第一个大于等于它的数的后面，然后把这个新加入的值后面的元素删掉。

当队列里面没有比它大的元素的时候，就把这个队列所有的元素删掉，然后加入这个元素。

在主程序里面，要做的就是初始化第一个窗口，然后循环向后，每次都把队列的一个元素放入结果中，然后判断当前窗口的最左元素和队列的第一个元素是否相同，如果相同就把这第一个元素出队，并且继续把新元素插入，维持这个队列。这么说其实我的代码可以更简洁的。

```go
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n-k+1)
	queue := make([]int, 1)
	for i := 0; i < k; i++ {
		queue = adjQueue(queue, nums[i])
	}
	for i := k; i < n; i++ {
		res[i-k] = queue[0]
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		queue = adjQueue(queue, nums[i])
	}
	res[n-k] = queue[0]
	return res
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
>执行用时：16 ms, 在所有 Go 提交中击败了95.70%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了70.73%的用户

修改后内存占用还小了一丢丢，惊了。