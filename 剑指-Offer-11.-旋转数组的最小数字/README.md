# 剑指 Offer 11. 旋转数组的最小数字
[https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)

## 二分法
```go
func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	p := (l + r) / 2
	for l < r {
		if numbers[p] > numbers[r] {
			l = p + 1
		} else if numbers[p] < numbers[r] {
			r = p
		} else {
			r--
		}
		p = (l + r) / 2
	}
	return numbers[p]
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了92.38%的用户
内存消耗：3.1 MB, 在所有 Go 提交中击败了51.97%的用户

这个题本质是要你找这个数组中的最小数。那么用遍历的方法可以得到，不过时间复杂度是O(n)。

看到有序数列就想起了二分法，二分法的思想是这个是两个有序数列的拼接而成的，并且左部（可以没有）的任意值都不小于右部的任意值。那么我们可以通过判断中间的值与最右的数的大小来得知最小数应该落在`[l,mid]`，还是`[mid+1,r]`区间里面。
- 当`nums[mid]>nums[r]`的时候我们可以得知中间的这个数比右边最大的数还大，那么最小数当然应该落在`[mid+1,r]`这个区间中。
- 当`nums[mid]<nums[r]`的时候我们可以得知中间的这个数比右边最大的数小，那么最小数当然应该落在`[l,mid]`这个区间中。
- 当`nums[mid]=nums[r]`的时候我们可以得知中间的这个数和右边最大的数相等，那么它可能在`[l,r]`这个区间中。但是因为我们知道数组是递增的，所以我们可以把范围缩小一点，变成`[l,r-1]`。因为此时这个中间的下标`mid`并不会大于`r`，所以`mid <= r-1`是必然成立的。

不断重复上面的过程，最后得到的结果就是要找的最小数

## 坑点
数据的范围。假设第一次就判断出了，那么`mid`也应该及时更新。我第一次`WA`了就是因为是先计算`mid`，再进行的判断。更新并不及时。