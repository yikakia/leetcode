# 704.二分查找
[https://leetcode-cn.com/problems/binary-search](https://leetcode-cn.com/problems/binary-search) 
## 原题
给定一个 `n` 个元素有序的（升序）整型数组 `nums` 和一个目标值 `target` ，写一个函数搜索 `nums` 中的 `target` ，如果目标值存在返回下标，否则返回 `-1` 。
 **示例 1:** 

```
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

```
 **示例 2:** 

```
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

```
 

 **提示：** 
- 你可以假设 `nums` 中的所有元素是不重复的。
-  `n` 将在 `[1, 10000]` 之间。
-  `nums` 的每个元素都将在 `[-9999, 9999]` 之间。
 
**标签**
`数组` `二分查找` 


## 
```go
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}
```
>执行用时: 28 ms
内存消耗: 6.7 MB

当取 l = 0 ,r = n -1 的时候说明 r 所指向的元素有可能是要找的，即可能的区间是 [l,r] ，因此当 mid == r 的时候，还要进行一次判断。因此循环判断的条件应该是 l <= r

而当 l = 0 , r = n 的时候说明 r 所指向的元素是有可能要找的元素的下一个，即可能的区间是 [l,r)，因此当 mid == r 的时候说明已经越界了，并不存在。因此循环判断的条件应该是 l < r 

