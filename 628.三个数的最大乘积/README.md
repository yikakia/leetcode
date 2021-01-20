# 628.三个数的最大乘积
[https://leetcode-cn.com/problems/maximum-product-of-three-numbers/](https://leetcode-cn.com/problems/maximum-product-of-three-numbers/) 
## 原题
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

**示例 1:** 

```

输入: [1,2,3]
输出: 6

```
**示例 2:** 

```

输入: [1,2,3,4]
输出: 24

```
**注意:** 
- 给定的整型数组长度范围是[3,10^4]，数组中所有的元素范围是[-1000, 1000]。
- 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
 
**标签**
`数组` `数学` 


## 模拟
```go
func maximumProduct(nums []int) int {
	const m = 1001
	max1, max2, max3 := -m, -m, -m
	min1, min2 := m, m
	for _, x := range nums {
		if x >= max1 {
			max1, max2, max3 = x, max1, max2
		} else if x >= max2 {
			max2, max3 = x, max2
		} else if x >= max3 {
			max3 = x
		}
		if x <= min1 {
			min1, min2 = x, min1
		} else if x <= min2 {
			min2 = x
		}
	}
	if min2 == m {
		min2 = 0
	}
	if max1 < 0 {
		return max1 * max2 * max3
	}

	return max1 * max(max2*max3, min1*min2)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

```
>执行用时: 44 ms
内存消耗: 6.5 MB

遍历的时候进行比较，找到最大的三个数以及最小的两个数。

最后的结果，
- 如果最大的数是负数，那么最大的乘积就是最大的三个数相乘（此时的乘积为负数，则要选取绝对值尽可能小的三个负数）
- 如果最大的数是正数，那么最大的乘积就是最大的三个数相乘与最大的数和最小的两个数相乘进行比较（最小的两个数可能为负数）