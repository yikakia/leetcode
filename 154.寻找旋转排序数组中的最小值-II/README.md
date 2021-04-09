# 154.寻找旋转排序数组中的最小值 II
[https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii) 
## 原题
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 `[0,1,2,4,5,6,7]` **** 可能变为 `[4,5,6,7,0,1,2]` )。

请找出其中最小的元素。

注意数组中可能存在重复的元素。

 **示例 1：** 

```
输入: [1,3,5]
输出: 1
```
 **示例 2：** 

```
输入: [2,2,2,0,1]
输出: 0
```
 **说明：** 
- 这道题是 <a href="https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/">寻找旋转排序数组中的最小值</a> 的延伸题目。
- 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？
 
**标签**
`数组` `二分查找` 


## 二分查找
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
		case nums[mid] < nums[right]:
			right = mid
		case nums[mid] == nums[right]:
			right--
		default:
		}
		mid = (left + right) / 2
	}

	return nums[left]
}
```
>执行用时: 4 ms
内存消耗: 3.1 MB

和[153.寻找旋转排序数组中的最小值](../153.寻找旋转排序数组中的最小值/README.md)差不多。比前面的小的就是就是要找的元素。不过还有一点不同的就是有重复的元素，所以最后就不一定能找到严格符合要求的节点，这就要求我们迭代后的 left 或者 right 就要满足条件。

我们再看一次这个题目，当当前的元素和右边界的元素进行比较的时候，
- 如果当前的元素更大，那么说明小的元素应该在右边的区间此时 left = mid + 1
- 如果当前的元素更小，说明小的元素应该在包含 mid 的左边的区间，此时 right = mid 
- 如果当前的元素和右边界的元素相同，那么并不能判断更小的元素是在左区间还是右区间，但是可以知道的一点是，将右边界左移一位的话，还是存在和右边界相同的元素，即 mid ，因此缩小了区间也不会影响结果。


