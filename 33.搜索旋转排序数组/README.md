# 33.搜索旋转排序数组
[https://leetcode-cn.com/problems/search-in-rotated-sorted-array](https://leetcode-cn.com/problems/search-in-rotated-sorted-array) 
## 原题
整数数组 `nums` 按升序排列，数组中的值 **互不相同** 。

在传递给函数之前， `nums` 在预先未知的某个下标 `k` （ `0 <= k < nums.length` ）上进行了 **旋转** ，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` （下标 **从 0 开始** 计数）。例如， `[0,1,2,4,5,6,7]` 在下标 `3` 处经旋转后可能变为  `[4,5,6,7,0,1,2]` 。

给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，如果 `nums` 中存在这个目标值 `target` ，则返回它的下标，否则返回  `-1`  。

 

 **示例 1：** 

```

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

```
 **示例 2：** 

```

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
```
 **示例 3：** 

```

输入：nums = [1], target = 0
输出：-1

```
 

 **提示：** 
-  `1 <= nums.length <= 5000` 
-  `-10^4 <= nums[i] <= 10^4` 
-  `nums` 中的每个值都 **独一无二** 
- 题目数据保证 `nums` 在预先未知的某个下标上进行了旋转
-  `-10^4 <= target <= 10^4` 
 

 **进阶：** 你可以设计一个时间复杂度为 `O(log n)` 的解决方案吗？

 
**标签**
`数组` `二分查找` 


## 判断有序区间后二分查找
```go
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if r == -1 {
		return -1
	}
	if r == 0 {
		if nums[r] != target {
			return -1
		}
		return 0
	}
	mid := (l + r) / 2
	switch {
	case nums[mid] == target:
		return mid
	case nums[mid] > nums[r]:
		index := bSearch(nums[:mid], target)
		if index != -1 {
			return index
		}
		if index = search(nums[mid+1:], target); index == -1 {
			return -1
		}
		return mid + 1 + index
	default:
		index := bSearch(nums[mid:], target)
		if index != -1 {
			return mid + index
		}
		return search(nums[:mid], target)
	}
}

func bSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return -1
}
```
>执行用时: 0 ms
内存消耗: 2.5 MB

这道题的思路和 [81.搜索旋转排序数组-II](..\81.搜索旋转排序数组-II\README.md) 的类似，就是先判断出哪部分是有序的区间，然后对有序的区间进行查找，对无序的区间继续调用本函数。