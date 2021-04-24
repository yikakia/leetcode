# 377.组合总和 Ⅳ
[https://leetcode-cn.com/problems/combination-sum-iv](https://leetcode-cn.com/problems/combination-sum-iv) 
## 原题
给你一个由 **不同** 整数组成的数组 `nums` ，和一个目标整数 `target` 。请你从 `nums` 中找出并返回总和为 `target` 的元素组合的个数。

题目数据保证答案符合 32 位整数范围。

 

 **示例 1：** 

```

输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。

```
 **示例 2：** 

```

输入：nums = [9], target = 3
输出：0

```
 

 **提示：** 
-  `1 <= nums.length <= 200` 
-  `1 <= nums[i] <= 1000` 
-  `nums` 中的所有元素 **互不相同** 
-  `1 <= target <= 1000` 
 

 **进阶：** 如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？

 
**标签**
`动态规划` 


## dp
```go
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
```
>执行用时: 0 ms
内存消耗: 2 MB

`dp[i]` 代表 `i` 有多少种组合的方式。从小到大推算有多少种组合形式。主要就是为了避免重复计算。

而推算的方式就是对 `nums` 进行遍历，取 `num` 作为结尾有多少种方式即 取 `dp[i-num]` 有多少种组合方式 累加进 `dp[i]` 。而如果不用  `nums` 中的值进行遍历的话就会出现一部分的组合被重复计算的问题。因为题目里面说了没有重复的数字，所以可以将其作为单独的元素进行考虑。