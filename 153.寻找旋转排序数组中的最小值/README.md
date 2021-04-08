# 153.寻找旋转排序数组中的最小值
[https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array) 
## 原题
已知一个长度为 `n` 的数组，预先按照升序排列，经由 `1` 到 `n` 次 **旋转** 后，得到输入数组。例如，原数组 `nums = [0,1,2,4,5,6,7]` 在变化后可能得到：

- 若旋转 `4` 次，则可以得到 `[4,5,6,7,0,1,2]` 
- 若旋转 `4` 次，则可以得到 `[0,1,2,4,5,6,7]` 
注意，数组 `[a[0], a[1], a[2], ..., a[n-1]]` **旋转一次** 的结果为数组 `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]` 。

给你一个元素值 **互不相同** 的数组 `nums` ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 **最小元素** 。

 

 **示例 1：** 

```

输入：nums = [3,4,5,1,2]
输出：1
解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。

```
 **示例 2：** 

```

输入：nums = [4,5,6,7,0,1,2]
输出：0
解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。

```
 **示例 3：** 

```

输入：nums = [11,13,15,17]
输出：11
解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。

```
 

 **提示：** 
-  `n == nums.length` 
-  `1 <= n <= 5000` 
-  `-5000 <= nums[i] <= 5000` 
-  `nums` 中的所有整数 **互不相同** 
-  `nums` 原来是一个升序排序的数组，并进行了 `1` 至 `n` 次旋转
 
**标签**
`数组` `二分查找` 


## 二分
```go
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if right < 1 {
		return nums[0]
	}
	if right == 1 {
		return min(nums[0], nums[1])
	}

	prev := func(i int) int {
		if i > 0 {
			return nums[i-1]
		}
		return nums[len(nums)-1]
	}

	mid := right / 2

	for left <= right {
		switch {
		case nums[mid] < prev(mid):
			return nums[mid]
		case nums[mid] > nums[right]:
			left = mid + 1
		default:
			right = mid - 1
		}
		mid = (left + right) / 2
	}

	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```
>执行用时: 0 ms
内存消耗: 2.5 MB

本质就是二分，当 `nums[mid]` 比它的前一个元素小的时候，那么它就是最小的那个元素，这个时候返回即可。因为这个旋转相当于是把首尾连接在一起，因此最后会出现一个谷底，也只会存在一个谷底。

而不是要找的元素的时候就要判断要找的那个元素可能是在左边的区间还是右边的区间，当中间的元素比最右的元素大的时候，就说明右边的最大的元素都比中间的这个元素小，而我们要找的是最小的元素，这个时候就应该缩小左边界。

而其他情况下就是在左区间中了，这个时候就缩小右边界。
