# 1646.获取生成数组中的最大值
[https://leetcode-cn.com/problems/get-maximum-in-generated-array](https://leetcode-cn.com/problems/get-maximum-in-generated-array) 
## 原题
给你一个整数 `n` 。按下述规则生成一个长度为 `n + 1` 的数组 `nums` ：
-  `nums[0] = 0` 
-  `nums[1] = 1` 
- 当 `2 <= 2 * i <= n` 时， `nums[2 * i] = nums[i]` 
- 当 `2 <= 2 * i + 1 <= n` 时， `nums[2 * i + 1] = nums[i] + nums[i + 1]` 
返回生成数组 `nums` 中的 **最大** 值。

 

 **示例 1：** 

```

输入：n = 7
输出：3
解释：根据规则：
  nums[0] = 0
  nums[1] = 1
  nums[(1 * 2) = 2] = nums[1] = 1
  nums[(1 * 2) + 1 = 3] = nums[1] + nums[2] = 1 + 1 = 2
  nums[(2 * 2) = 4] = nums[2] = 1
  nums[(2 * 2) + 1 = 5] = nums[2] + nums[3] = 1 + 2 = 3
  nums[(3 * 2) = 6] = nums[3] = 2
  nums[(3 * 2) + 1 = 7] = nums[3] + nums[4] = 2 + 1 = 3
因此，nums = [0,1,1,2,1,3,2,3]，最大值 3

```
 **示例 2：** 

```

输入：n = 2
输出：1
解释：根据规则，nums[0]、nums[1] 和 nums[2] 之中的最大值是 1

```
 **示例 3：** 

```

输入：n = 3
输出：2
解释：根据规则，nums[0]、nums[1]、nums[2] 和 nums[3] 之中的最大值是 2

```
 

 **提示：** 
-  `0 <= n <= 100` 
 
**标签**
`数组` `动态规划` `模拟` 


## 水题
```go
func getMaximumGenerated(n int) int {
	switch n {
	case 0:
		return 0
	}
	nums := make([]int, n+1)
	ans := 0
	nums[1] = 1
	for i := 0; i <= n; i++ {
		t := 2 * i
		if t >= 2 && t <= n {
			nums[t] = nums[i]
			ans = max(ans, nums[t])
		}
		if t+1 >= 2 && t+1 <= n {
			nums[t+1] = nums[i] + nums[i+1]
			ans = max(ans, nums[t+1])
		}
	}
	return max(ans, nums[n])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```
>执行用时: 0 ms
内存消耗: 2 MB
