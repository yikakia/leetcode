# 剑指 Offer 57. 和为s的两个数字
[https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/](https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/)
## 
```go
func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n == 1 {
		return []int{}
	}
	lo, hi := 0, n-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	lo = 0
	for s := nums[lo] + nums[hi]; lo <= hi && s != target; s = nums[lo] + nums[hi] {
		if s > target {
			hi--
		} else {
			lo++
		}
	}
	if nums[lo]+nums[hi] == target {
		return []int{nums[lo], nums[hi]}
	}
	return []int{}
}
```
>执行用时：212 ms, 在所有 Go 提交中击败了82.58%的用户
内存消耗：8.6 MB, 在所有 Go 提交中击败了91.98%的用户

因为是有序数组，那么就可以考虑二分了。又已知最小值为1，所以大于零，那么可以先通过筛去大于 `target` 的数来剪枝。这里可以用到二分查找。

二分查找之后，就是从两头开始找了。

首先我们应该知道，在最右边的指针确定的情况下，左边的指针越大，那么和就越大。

而在最左边的指针确定的情况下，右边的指针越小，那么和就越小。

现在我们可以开始讨论了，当两个指针指向的元素和大于目标值时，那么后面的指针应该向前移动。因为左边的最小值和右部的值加起来都比目标值大了，那么左部更大的数和右部加起来也肯定比目标值大，那么应该减小右部的值。

而加起来比目标值小，那么就是左部不够大，因为最右边的数是最大的数，加起来都不能满足和了，那么必须要增大左部。

那么问题来了，为什么不减小左部，增大右部呢？很简单，因为我们最开始就设定整个区间的最小值和最大值，这个时候已经不能再拓展区间了，只能缩小区间。

那么查找的时候会不会漏过值呢？你看这个一会大于一会小于的。事实上是不会的，因为你缩小最大值的时候，是在最小值加上最大值都大于了目标的情况。而增大最小值的时候，是最大值都不够的情况。
