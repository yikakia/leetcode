# 740.删除并获得点数
[https://leetcode-cn.com/problems/delete-and-earn](https://leetcode-cn.com/problems/delete-and-earn) 
## 原题
给你一个整数数组  `nums`  ，你可以对它进行一些操作。

每次操作中，选择任意一个  `nums[i]`  ，删除它并获得  `nums[i]`  的点数。之后，你必须删除 **每个** 等于  `nums[i] - 1`  或  `nums[i] + 1`  的元素。

开始你拥有 `0` 个点数。返回你能通过这些操作获得的最大点数。

 

 **示例 1：** 

```

输入：nums = [3,4,2]
输出：6
解释：
删除 4 获得 4 个点数，因此 3 也被删除。
之后，删除 2 获得 2 个点数。总共获得 6 个点数。

```
 **示例 2：** 

```

输入：nums = [2,2,3,3,3,4]
输出：9
解释：
删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
总共获得 9 个点数。

```
 

 **提示：** 
-  `1 <= nums.length <= 2 * 10^4` 
-  `1 <= nums[i] <= 10^4` 
 
**标签**
`动态规划` 


## 打家劫舍
```go
func deleteAndEarn(nums []int) int {
	var maxNum int
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}
	dp := make([]int, maxNum+1)
	for _, num := range nums {
		dp[num]++
	}
	return rob(dp)
}

func rob(dp []int) int {
	pre1, pre2 := dp[0], max(dp[0], dp[1])

	for i := 2; i < len(dp); i++ {
		pre1, pre2 = pre2, max(pre1+dp[i]*i, pre2)
	}
	return pre2
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```
>执行用时: 4 ms
内存消耗: 2.9 MB

[打家劫舍](../213.打家劫舍-II/README.md)是相邻的家不能同时偷，也就相当于这里的相邻的数字不能同时取。因此思路是相同的，对于每一个值，作为最后的结果来说都是可能被取和不可能被取两种状态。这个时候就让其记录下取它能获得的最大值和不取它能获得的最大值即可。
