# 503.下一个更大元素 II
[https://leetcode-cn.com/problems/next-greater-element-ii](https://leetcode-cn.com/problems/next-greater-element-ii) 
## 原题
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

 **示例 1:** 

```

输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数； 
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。

```
 **注意:** 输入数组的长度不会超过 10000。

 
**标签**
`栈` 


## 暴力
```go
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := range nums {
		var index int
		for j := 1; j <= n; j++ {
			if nums[(i+j)%n] > nums[i] {
				index = (i + j) % n
				break
			}
		}
		if nums[index] > nums[i] {
			res[i] = nums[index]
		} else {
			res[i] = -1
		}
	}

	return res
}
```
>执行用时: 496 ms
内存消耗: 6.5 MB

我就是个傻狗，每次向右找到第一个满足条件的数就行了。如果找不到就置 `-1`。不过这个耗时太长，肯定有优化的办法。

## 单调栈
```go
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	stack := make([]int, 0, n)
	for i := 0; i < 2*n-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return res
}
```
>执行用时: 44 ms
内存消耗: 6.6 MB

单调栈是写过不少，但是对于哪些问题可以用并不清楚看了[别人的题解](https://leetcode-cn.com/problems/next-greater-element-ii/solution/cong-po-su-jie-fa-de-jiao-du-qu-li-jie-d-trht/)才明白，原来是对于 
>找最近一个比当前值大/小 

这类问题都可以用单调栈来解决。

对于这道题而言，我们是把当前没得到答案的下标存在栈中。当遍历到的元素的值比栈中元素所指向的值大的时候，就说明栈中元素所指向的值的最近一个比它大的元素就是当前遍历到的元素。

因为如果存在更近的满足条件的元素的话，它在之前就会出栈了。

在具体实现的时候，因为要找到每个元素的下一个最大值，所以要对原数组循环两遍。比如 [3,2,1] 这样的数组就是第二遍遍历到 3 的时候 就把栈里面的 [1,2] 给出栈了。但是对于最大的 [3] 却不能得到更新。因为题目要求没有更大的数就标为 -1 所以提前先初始化为 -1 即可 