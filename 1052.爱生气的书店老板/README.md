# 1052.爱生气的书店老板
[https://leetcode-cn.com/problems/grumpy-bookstore-owner](https://leetcode-cn.com/problems/grumpy-bookstore-owner) 
## 原题
今天，书店老板有一家店打算试营业 `customers.length` 分钟。每分钟都有一些顾客（ `customers[i]` ）会进入书店，所有这些顾客都会在那一分钟结束后离开。

在某些时候，书店老板会生气。 如果书店老板在第 `i` 分钟生气，那么 `grumpy[i] = 1` ，否则 `grumpy[i] = 0` 。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。

书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 `X` 分钟不生气，但却只能使用一次。

请你返回这一天营业下来，最多有多少客户能够感到满意的数量。<br>
 

 **示例：** 

```
输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
输出：16
解释：
书店老板在最后 3 分钟保持冷静。
感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.

```
 

 **提示：** 
-  `1 <= X <= customers.length == grumpy.length <= 20000` 
-  `0 <= customers[i] <= 1000` 
-  `0 <= grumpy[i] <= 1` 
 
**标签**
`数组` `Sliding Window` 


## 滑动窗口
```go
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxSatisfied(customers []int, grumpy []int, X int) (res int) {
	n := len(customers)
	for i := range customers {
		res += customers[i] * (1 - grumpy[i])
	}
	var l, r, update, maxupdate int
	for ; r < X; r++ {
		update += customers[r] * grumpy[r]
	}
	maxupdate = update
	for ; r < n; r++ {
		update += -customers[l]*grumpy[l] + customers[r]*grumpy[r]
		maxupdate = max(maxupdate, update)
		l++
	}
	return res + maxupdate
}
```
>执行用时: 48 ms
内存消耗: 6.6 MB

题目的要求转化一下就是当 `grumpy[i] == 0` 的时候，把 `customers[i]` 的值累加进结果 `res` 中。然后开一个长度为 `X` 的窗口用来统计其中 `grumpy[i] == 1` 的时候的 `customers[i]` 的累加值。

于是用乘法来写简化代码。但是时间就长了点。于是用 `if else `来实现试试

### if else
```go
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxSatisfied(customers []int, grumpy []int, X int) (res int) {
	n := len(customers)
	for i := range customers {
		if grumpy[i] == 0 {
			res += customers[i]
		}
	}
	var l, r, update, maxupdate int
	for ; r < X; r++ {
		if grumpy[r] == 1 {
			update += customers[r]
		}
	}
	maxupdate = update
	for r < n {
		if grumpy[l] == 1 {
			update -= customers[l]
		}
		l++
		if grumpy[r] == 1 {
			update += customers[r]
		}
		r++
		maxupdate = max(maxupdate, update)
	}
	return res + maxupdate
}
```
>执行用时: 32 ms
内存消耗: 6.6 MB

快了 30 % 这个优化也太大了