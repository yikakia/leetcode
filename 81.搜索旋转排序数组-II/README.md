# 81.搜索旋转排序数组 II
[https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii) 
## 原题
已知存在一个按非降序排列的整数数组 `nums` ，数组中的值不必互不相同。

在传递给函数之前， `nums` 在预先未知的某个下标 `k` （ `0 <= k < nums.length` ）上进行了 **旋转** ，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` （下标 **从 0 开始** 计数）。例如， `[0,1,2,4,4,4,5,6,6,7]` 在下标 `5` 处经旋转后可能变为 `[4,5,6,6,7,0,1,2,4,4]` 。

给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 `nums` 中存在这个目标值 `target` ，则返回 `true` ，否则返回 `false` 。

 

 **示例 1：** 

```

输入：nums = [2,5,6,0,0,1,2], target = 0
输出：true

```
 **示例 2：** 

```

输入：nums = [2,5,6,0,0,1,2], target = 3
输出：false
```
 

 **提示：** 
-  `1 <= nums.length <= 5000` 
-  `-10^4 <= nums[i] <= 10^4` 
- 题目数据保证 `nums` 在预先未知的某个下标上进行了旋转
-  `-10^4 <= target <= 10^4` 
 

 **进阶：** 
- 这是 <a href="https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/">搜索旋转排序数组</a> 的延伸题目，本题中的  `nums`   可能包含重复元素。
- 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 
**标签**
`数组` `二分查找` 


## 先判断有序再二分查找
```go
func search(nums []int, target int) bool {
	mid, n := len(nums)/2, len(nums)
	if n == 0 {
		return false
	}
	if nums[mid] == target {
		return true
	}
	if mid == 0 {
		return false
	}
	switch {
	case nums[mid] > nums[0]:
		return bSearch(nums[:mid+1], target) || search(nums[mid+1:], target)
	case nums[mid] < nums[n-1]:
		return bSearch(nums[mid:], target) || search(nums[:mid], target)
	default:
		return search(nums[:mid], target) || search(nums[mid:], target)
	}
}

func bSearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
```
>执行用时: 4 ms
内存消耗: 3.2 MB

这道题和 [33.搜索旋转排序数组](../33.搜索旋转排序数组/README.md) 的思路差不多，就先判断二分后的区间哪部分是有序区间，然后对有序区间进行二分，对无序区间再次调用本算法即可。

不过因为本题会出现相同的元素，所以二分后的两个区间都可能是无序区间。比如 `[2,1,2,2,2]` ，这里的 `nums[mid]` 表现出来的性质就是 `nums[mid]<=nums[n] && nums[mid]<=nums[0]` 而 `[2,2,2,1,2]` 表现出来的性质也是一样的。所以需要对两个区间都再次调用本算法才行。
