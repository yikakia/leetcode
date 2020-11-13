# 剑指 Offer 42. 连续子数组的最大和
[https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/](https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/)

## 动规
```go
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = nums[n-1]
	max := dp[n-1]
	for i := n - 2; i >= 0; i-- {
		dp[i] = nums[i]
		if dp[i+1] > 0 {
			dp[i] += dp[i+1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
```
>执行用时：40 ms, 在所有 Go 提交中击败了6.58%的用户
内存消耗：7.4 MB, 在所有 Go 提交中击败了10.66%的用户

让`dp[n]`代表在`n`这个位置之后，可以得到的最大累加。解题思路简单地说就是从尾部到头部累加，为了找到最大的子区间，那么对于`dp[n]`而言，要找到最大的部分就得看它的下一位`dp[n+1]`是不是正数。如果是正数那么就是可以累加的，如果不是，那么它自己就是最大子区间。

而对于 `max` 而言，就是在遍历整个数组的时候随时维护，更新。看目前求得的`dp[n]`是不是最大的就可以了。

不过这个时间效率太低了。看了下题解，好像是应该先判断应不应该更新，再进行更新更省时间？
```go

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = nums[n-1]
	max := dp[n-1]
	for i := n - 2; i >= 0; i-- {
		if dp[i+1] > 0 {
			dp[i] = dp[i+1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
```
>执行用时：24 ms, 在所有 Go 提交中击败了86.42%的用户
内存消耗：7.4 MB, 在所有 Go 提交中击败了10.66%的用户

这么一来时间的确缩短了。看来赋值非常消耗时间啊。

而内存消耗方面，则是不用数组，而是用一个临时变量来暂存当前的区间的值。用这个临时变量是否小于零来判断是应该把当前的值作为当前区间的最大值，还是与之前的区间相加。

```go

func maxSubArray(nums []int) int {
	n := len(nums)
	max, answer := nums[n-1], nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if answer < 0 {
			answer = nums[i]
		} else {
			answer += nums[i]
		}
		if answer > max {
			max = answer
		}
	}
	return max
}
```
>执行用时：24 ms, 在所有 Go 提交中击败了86.42%的用户
内存消耗：7.1 MB, 在所有 Go 提交中击败了43.23%的用户

内存消耗的确小了，但是并不明显。看了题解感觉和别人的没啥区别，很奇怪。