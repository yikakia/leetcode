# 978.最长湍流子数组
[https://leetcode-cn.com/problems/longest-turbulent-subarray/](https://leetcode-cn.com/problems/longest-turbulent-subarray/) 
## 原题
当 `A` 的子数组 `A[i], A[i+1], ..., A[j]` 满足下列条件时，我们称其为*湍流子数组*：
- 若 `i <= k < j`，当 `k` 为奇数时， `A[k] > A[k+1]`，且当 `k` 为偶数时，`A[k] < A[k+1]`；
- **或**若 `i <= k < j`，当 `k` 为偶数时，`A[k] > A[k+1]` ，且当 `k` 为奇数时， `A[k] < A[k+1]`。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

返回 `A` 的最大湍流子数组的**长度**。

 

**示例 1：**

```
输入：[9,4,2,10,7,8,8,1,9]
输出：5
解释：(A[1] > A[2] < A[3] > A[4] < A[5])

```
**示例 2：**

```
输入：[4,8,12,16]
输出：2

```
**示例 3：**

```
输入：[100]
输出：1

```
 

**提示：**
- `1 <= A.length <= 40000`
- `0 <= A[i] <= 10^9`
 
**标签**
`数组` `动态规划` `Sliding Window` 


## 滑动窗口
```go
func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	switch {
	case n < 2:
		return 1
	case n == 2:
		if arr[0] == arr[1] {
			return 1
		}
		return 2
	}
	compare := make([]int, n-1)
	for i := range arr[1:] {
		switch {
		case arr[i] > arr[i+1]:
			compare[i] = -1
		case arr[i] < arr[i+1]:
			compare[i] = 1
		case arr[i] == arr[i+1]:
			compare[i] = 0
		}
	}
	st, end := 0, 1
	res := 1
	for end < n-1 {
		if t := compare[end-1] * compare[end]; t == -1 {
			if t := end - st + 2; t > res {
				res = t
			}
			end++
		} else if t == 0 {
			x, y := compare[st], compare[end]
			if !(x == 0 && y == 0) {
				if res < 2 {
					res = 2
				}
			}
			end++
			st = end - 1
		} else if t == 1 {
			end++
			st = end - 1
			if res < 2 {
				res = 2
			}
		} else {
			panic("unreachable")
		}
	}
	return res
}
```
>执行用时: 80 ms
内存消耗: 7.5 MB

这道题做得很差，主要是感觉是面向样例编程了。

如果只是看增减的话是没问题的，关键是出现了相同的情况应该怎么处理。

我们先默认右部每次只增长 **1** ，那么当我们遇到了有一个相同的情况时，可能是左部有一个相同，也可能是右部有一个相同，还有可能是两个都相同。

当
- 只有一个相同的时候，存在一个长度为 **2** 的子数组。
- 而两个都相同的时候，存在一个长度为 **1** 的子数组。

这个问题解决就简单了。还要注意的是如果出现了连续增长的情况，那么存在一个长度为 **2** 的子数组。