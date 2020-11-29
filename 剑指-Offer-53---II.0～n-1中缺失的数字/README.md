# 剑指 Offer 53 - II. 0～n-1中缺失的数字
[https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/](https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/)
## 二分查找数值比自己下标大的
```go
func missingNumber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	lo, hi := 0, n-1
	mid := lo + (hi-lo)/2
	for lo <= hi {
		if nums[mid] > mid {
			hi = mid - 1
			mid = lo + (hi-lo)/2
		} else {
			lo = mid + 1
			mid = lo + (hi-lo)/2
		}
	}

	return lo
}
```
>执行用时：20 ms, 在所有 Go 提交中击败了68.64%的用户
内存消耗：6.1 MB, 在所有 Go 提交中击败了32.01%的用户

简单地说就是通过二分查找下标比自己的数值小的。然后不断地缩小区间。最后`lo`越过`hi`的时候就表示这个时候是介于出现和没出现之间的。