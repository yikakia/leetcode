# 剑指 Offer 53 - I. 在排序数组中查找数字 I
[https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/](https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/)

## 二分定位，然后计算
```go
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	count := 0
	lo, hi := 0, n-1
	mid := lo + (hi-lo)/2
	for lo < hi && nums[mid] != target {
		if nums[mid] > target {
			hi = mid
			mid = lo + (hi-lo)/2
		} else {
			lo = mid + 1
			mid = lo + (hi-lo)/2
		}
	}
	if nums[mid] != target {
		return 0
	}
	for i := mid; i >= 0 && nums[i] == target; i-- {
		count++
	}
	for i := mid + 1; i < n && nums[i] == target; i++ {
		count++
	}
	return count
}

```

>执行用时：8 ms, 在所有 Go 提交中击败了91.09%的用户
内存消耗：3.9 MB, 在所有 Go 提交中击败了77.68%的用户

本来觉得是白给的题目，一路跑过去就行了。一看题目里面有说这是有序数组，那么直接上二分吧！二分查找之后的结果可能是这个数的中间位置，开头和最后。所以要向两边分别遍历，同时还要注意数组越界的问题。