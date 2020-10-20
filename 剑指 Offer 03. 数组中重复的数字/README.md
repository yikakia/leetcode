# 剑指 Offer 03. 数组中重复的数字
[https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)

## 快慢指针
```go
func findRepeatNumber(nums []int) int {
	start := 0
	for nums[start] == start {
		start++
	}
	fast := start
	slow := start
	for nums[nums[fast]] != nums[slow] {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	fast = start
	for nums[fast] != nums[slow] {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	return nums[fast]
}
```


## 置换
```go
func findRepeatNumber(nums []int) int {
	for index, val := range nums {
		if index == val {
			continue
		}
		if nums[val] == val {
			return val
		}
		nums[index], nums[val] = nums[val], nums[index]
	}
	return -1
}
```
将数组的值看作索引，遍历数组

当
1.  下标和值相等时，跳过。因为不在一个环里面
2.  该位置的值作为索引的新位置的值和该位置的值相等，则就找到了重复的数
3.  当没找到时，交换该位置上的值与该位置上的值指向的位置上的值。

思路就是这样。关键是为什么这么可以找到。

假设是没有重复数的数组并且按照升序排列，那么我们可以得到的是一列指向自己的数组。而当存在重复数，则存在一个或多个位置的入度大于1。那我们要做的就是对于每一个位置而言，如果不是指向自己，那么我们要做的就是把一个已经知道位置的数，放到它应该在的位置上。当遇到第二个相同的数时，我们就会知道这个数已经在前面出现过了。

换句话说，我们是通过遍历数组，把遍历过的数组所指向的位置抽象成了一个 seen 数组，用来表示这个是已经出现过的了。而我们遍历过的数组则是储存还没有被记录过的值。


简单的说，我们通过判断 `nums[val] == val` 与否来判断了 `val` 这个值出现过没有。因为当你要访问 `nums[val]` 的时候，`nums[index]` 的值就是 `val` ，如果此时 `nums[val] == val` ,那么很明显这个就是重复的数。当 `nums[val] != val` 的时候，我们就知道这个值还没有出现过，这个时候我们就应该把这个值出现过记录下来，即 `nums[val] = val` 。同时 `nums[val]` 原本的值不应该被消去，因为这个的值可能还没有被记录过，所以应该做的是交换 `nums[index]` 和 `nums[val]` 的值。
      