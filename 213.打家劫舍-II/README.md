# 213.打家劫舍 II
[https://leetcode-cn.com/problems/house-robber-ii](https://leetcode-cn.com/problems/house-robber-ii) 
## 原题
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 **围成一圈** ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统， **如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警** 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 **在不触动警报装置的情况下** ，能够偷窃到的最高金额。

 

 **示例 1：** 

```

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

```
 **示例 2：** 

```

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。
```
 **示例 3：** 

```

输入：nums = [0]
输出：0

```
 

 **提示：** 
-  `1 <= nums.length <= 100` 
-  `0 <= nums[i] <= 1000` 
 
**标签**
`动态规划` 


## 动规
```go
func rob(nums []int) int {
	n := len(nums)
	if n < 4 {
		return max(nums...)
	}
	dp := make([]int, n)  // 抢 i
	dp2 := make([]int, n) // 不抢 i
	res1, res2 := 0, 0    // res1 表示选第一个的最大结果 res2 表示不选第一个的最大的结果

	dp[0] = nums[0]
	for i := 1; i < n-1; i++ {
		dp[i] = nums[i] + max(dp[(i+n-2)%n], dp2[(i+n-1)%n])
		dp2[i] = max(dp[(i+n-1)%n], dp2[(i+n-1)%n])
	}
	dp2[n-1] = max(dp[n-2], dp2[n-2])
	res1 = max(max(dp[n-3:]...), max(dp2[n-3:]...))

	dp = make([]int, n)
	dp2 = make([]int, n)

	for i := 1; i < n; i++ {
		dp[i] = nums[i] + max(dp[(i+n-2)%n], dp2[(i+n-1)%n])
		dp2[i] = max(dp[(i+n-1)%n], dp2[(i+n-1)%n])
	}
	res2 = max(max(dp[n-3:]...), max(dp2[n-3:]...))

	return max(res1, res2)
}

func max(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if res < num {
			res = num
		}
	}
	return res
}
```
>执行用时: 0 ms
内存消耗: 2 MB

简单地说对于最后的结果而言，为了达成最大的收益，每一间房都可能被打劫，把视角放到一间房上面的话和最后的选择相比有两种情况，一是这间房被选，一是这间房不被选。这里就分别对两种情况进行动规，为了方便就分成选第一间房的最大收益和不选第一间房的最大收益。最后选择两者中大的那一个就可以了。