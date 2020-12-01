# 34. 在排序数组中查找元素的第一个和最后一个位置
[https://leetcode-cn.com/problems/reorganize-string/](https://leetcode-cn.com/problems/reorganize-string/) 
##  二分
```go
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	n := len(nums)
	if n == 0 {
		return res
	}
	l, r := 0, n-1
	mid := r - (r-l)>>1
	for l < r && nums[mid] != target {
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
        }
        mid = r - (r-l)>>1
	}
	if mid >= n {
		mid = n - 1
	}
	if nums[mid] != target {
		return res
	}
	for i := mid; i >= 0 && nums[i] == target; i-- {
		res[0] = i
	}
	for i := mid; i < n && nums[i] == target; i++ {
		res[1] = i
	}
	return res
}
```
>执行用时: 8 ms
内存消耗: 3.9 MB

就是标准的二分查找嘛。不过还可以继续优化，就是不进行拓展，而是直接找最左边和最右边。

## 优化
```go
func searchRange(nums []int, target int) []int {
	defaultRes := []int{-1, -1}
	n := len(nums)
	if n == 0 {
		return defaultRes
	}
	l, r := 0, n
	mid := l + (r-l)>>1
	for l < r {
		mid = l + (r-l)>>1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	left := l
	if left == n || nums[left] != target {
		return defaultRes
	}

	res := []int{-1, -1}
	res[0] = left
	l, r = 0, n
	mid = l + (r-l)>>1
	for l < r {
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
		mid = l + (r-l)>>1
	}
	right := mid - 1
	res[1] = right
	return res
}
```
>执行用时: 8 ms
内存消耗: 3.9 MB

有点尴尬，基本没有变化。可能是数据量太小的缘故？
