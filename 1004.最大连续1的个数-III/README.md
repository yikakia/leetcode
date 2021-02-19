# 1004.最大连续1的个数 III
[https://leetcode-cn.com/problems/max-consecutive-ones-iii/](https://leetcode-cn.com/problems/max-consecutive-ones-iii/) 
## 原题
给定一个由若干 `0` 和 `1` 组成的数组 `A` ，我们最多可以将 `K` 个值从 0 变成 1 。

返回仅包含 1 的最长（连续）子数组的长度。

 

 **示例 1：** 

```
输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释： 
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
```

 **示例 2：** 

```
输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。
```
 

 **提示：** 
-  `1 <= A.length <= 20000` 
-  `0 <= K <= A.length` 
-  `A[i]` 为 `0` 或 `1` 
 
**标签**
`双指针` `Sliding Window` 


## 双指针
```go
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestOnes(A []int, K int) (res int) {
	var left, right, hasZero int
	n := len(A)
	for right < n {
		if A[right] == 0 {
			if hasZero < K {
				hasZero++
			} else {
				for hasZero >= K {
					if A[left] == 0 {
						hasZero--
					}
					left++
				}
				hasZero++
			}
		}
		right++
		res = max(res, right-left)
	}
	return
}
```
>执行用时: 72 ms
内存消耗: 6.9 MB

简单地说，双指针维护一个合法的区间，然后右指针不断地向右拓展就行。

关键就是当右指针拓展后不合法的话，就需要将左指针不断地向右移动，直到去除了一个为 `0` 的才结束。

因此还需要一个变量来统计当前区间中有多少个 `0` ，这样才能方便地判断当前区间是否合法。

而结果就每次拓展完后更新即可。