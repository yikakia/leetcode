# 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
[https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/)

## 伪快排
```go
func exchange(nums []int) []int {
	fastsort(nums)
	return nums
}
func fastsort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		isIEven := isEven(nums[i])
		for !isIEven && i < j {
			i++
			isIEven = isEven(nums[i])
		}
		isJOdd := isOdd(nums[j])
		for !isJOdd && i < j {
			j--
			isJOdd = isOdd(nums[j])
		}
		if isIEven && isJOdd {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if isEven(nums[0]) {
		nums[0], nums[i] = nums[i], nums[0]
	}
	return nums
}
func isOdd(a int) bool {
	return !isEven(a)
}

// isEven 当a是偶数时返回`true`，是奇数时返回`false`
func isEven(a int) bool {
	if a == 1 {
		return false
	}
	if a&1 == 0 {
		return true
	}
	return false
}
```
>执行用时：28 ms, 在所有 Go 提交中击败了37.16%的用户
内存消耗：6.3 MB, 在所有 Go 提交中击败了47.71%的用户

看到这个交换就可以认为是一个类似于快排的题目了。但是因为这个只用交换奇偶就可以了，所以是快排的弱化版本。只用跑一趟就可以搞定了。不过这个耗时挺慢的，我看下别人的题解，好像是一样的。不过没把这个直接当快排来写。其实我也是写完才发现这个只用跑一趟就可以搞定了，所以还是自己想复杂了。