# 368.最大整除子集
[https://leetcode-cn.com/problems/largest-divisible-subset](https://leetcode-cn.com/problems/largest-divisible-subset) 
## 原题
给你一个由 **无重复** 正整数组成的集合 `nums` ，请你找出并返回其中最大的整除子集 `answer` ，子集中每一元素对 `(answer[i], answer[j])` 都应当满足：

-  `answer[i] % answer[j] == 0` ，或
-  `answer[j] % answer[i] == 0` 
如果存在多个有效解子集，返回其中任何一个均可。

 

 **示例 1：** 

```

输入：nums = [1,2,3]
输出：[1,2]
解释：[1,3] 也会被视为正确答案。

```
 **示例 2：** 

```

输入：nums = [1,2,4,8]
输出：[1,2,4,8]

```
 

 **提示：** 
-  `1 <= nums.length <= 1000` 
-  `1 <= nums[i] <= 2 * 10^9` 
-  `nums` 中的所有整数 **互不相同** 
 
**标签**
`数学` `动态规划` 


## 动规
```go
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]struct{ index, lenth int }, n)
	for i := range dp {
		dp[i].index = -1
	}

	max := 0
	maxIndex := 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if t := dp[j].lenth + 1; dp[i].lenth < t {
					dp[i].index = j
					dp[i].lenth = t
				}
				if dp[i].lenth > max {
					max = dp[i].lenth
					maxIndex = i
				}
			}
		}
	}

	res := make([]int, 0, max+2)

	for i := maxIndex; i != -1; i = dp[i].index {
		res = append(res, nums[i])
	}

	sort.Ints(res)
	return res
}
```
>执行用时: 52 ms
内存消耗: 2.9 MB

记录自己合规的最长链路。整体的思路很简单就是这样。

题目没要求子集里面的顺序，这个就降低了挺多的难度。当大能够整除小的了，同时能够更新链路的话就更新当前链路。

最后输出就倒序添加进结果即可。这样的话是降序的顺序。需要升序的话就再加一个排序即可。