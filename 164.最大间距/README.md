# 164. 最大间距
[https://leetcode-cn.com/problems/maximum-gap/](https://leetcode-cn.com/problems/maximum-gap/) 
## 
```go
type pair struct{ min, max int }

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	res := 0
	maxValue := max(nums...)
	minValue := min(nums...)
	// 桶的大小
	d := max(1, (maxValue-minValue)/(n-1))
	bucketSize := (maxValue-minValue)/d + 1

	// 存储 (桶内最小值，桶内最大值) 对，(-1, -1) 表示该桶是空的
	buckets := make([]pair, bucketSize)
	for i := range buckets {
		buckets[i] = pair{-1, -1}
	}
	for _, v := range nums {
		bid := (v - minValue) / d
		if buckets[bid].min == -1 {
			buckets[bid].min = v
			buckets[bid].max = v
		} else {
			buckets[bid].min = min(buckets[bid].min, v)
			buckets[bid].max = max(buckets[bid].max, v)
		}
	}

	prev := -1
	for i, b := range buckets {
		if b.min == -1 {
			continue
		}
		if prev != -1 {
			res = max(res, b.min-buckets[prev].max)
		}
		prev = i
	}

	return res
}
func min(nums ...int) int {
	return extrema(func(k, x int) bool {
		if k > x {
			return true
		}
		return false
	}, nums...)
}
func max(nums ...int) int {
	return extrema(func(k, x int) bool {
		if k < x {
			return true
		}
		return false
	}, nums...)
}
func extrema(needUpdate func(extreme, x int) bool, nums ...int) int {
	res := nums[0]
	for _, v := range nums[1:] {
		if needUpdate(res, v) {
			res = v
		}
	}
	return res
}

```
>执行用时: 4 ms
内存消耗: 3.3 MB

参考了官方题解，简单地说就是进行桶排序，不过不用真正地进行排序，只用更新这个区间的最大值和最小值。

因为我们可以简单地得到 $A_i-A_{i-1}>=\frac{max-min}{n-1}$ 这里的**A**数组是已经排序后的数组，而最大值,最小值就不用解释了吧。**n** 是数组的个数。

因此我们可以将 $d=\lfloor\frac{max-min}{n-1}\rfloor$作为桶的大小，在里面的数的差一定小于 $\frac{max-min}{n-1}$，也就是说对于桶里面的相邻元素一定不是结果，结果只能是两个桶之间的。

结果就是取每个桶的最大值与之后的桶的最小值之差的最大值。

如此我们便可以得到了结果。实现上不难，难的是怎么想到这个规律。顺便这次也试着写了一个判断是不是极值点的函数，然后通过这个函数来生成求最大值和最小值的函数。感觉还不错。
