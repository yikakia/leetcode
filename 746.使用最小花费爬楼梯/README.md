# 746. 使用最小花费爬楼梯
[https://leetcode-cn.com/problems/min-cost-climbing-stairs/](https://leetcode-cn.com/problems/min-cost-climbing-stairs/) 
## 原题
数组的每个索引作为一个阶梯，第 `i`个阶梯对应着一个非负数的体力花费值 `cost[i]`(索引从0开始)。
每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。
您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。
**示例 1:** 
```
输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。
```
**示例 2:** 
```
输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。
```
**注意：** 
- `cost` 的长度将会在 `[2, 1000]`。
- 每一个 `cost[i]` 将会是一个Integer类型，范围为 `[0, 999]`。


## 变相的青蛙跳台阶
```go
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	copy(dp, cost)
	for i := 2; i < len(dp); i++ {
		dp[i] += min(dp[i-2 : i]...)
	}
	return min(dp[len(dp)-2:]...)
}
func min(nums ...int) (res int) {
	res = nums[0]
	for _, num := range nums {
		if res > num {
			res = num
		}
	}
	return
}
```
>执行用时: 4 ms
内存消耗: 3 MB

就是不断地找前面的最小项填补进来，最后结果是返回倒数两个中最小的那个。

## 滚动数组优化
```go
func minCostClimbingStairs(cost []int) int {
	cur, prev1, prev2 := 0, 0, 0
	for _, costi := range cost {
		cur = min(prev1, prev2) + costi
		prev1, prev2 = prev2, cur
	}
	return min(prev1, prev2)
}
func min(nums ...int) (res int) {
	res = nums[0]
	for _, num := range nums {
		if res > num {
			res = num
		}
	}
	return
}
```
>执行用时: 4 ms
内存消耗: 2.8 MB

注意到其实只用到了两个值，上一阶台阶和上两阶台阶。因此就可以用滚动数组来优化空间复杂度。