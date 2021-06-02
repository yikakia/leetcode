# 523.连续的子数组和
[https://leetcode-cn.com/problems/continuous-subarray-sum](https://leetcode-cn.com/problems/continuous-subarray-sum) 
## 原题
给定一个包含 **非负数** 的数组和一个目标 **整数**   `k` ，编写一个函数来判断该数组是否含有连续的子数组，其大小至少为 2，且总和为 `k` 的倍数，即总和为 `n * k` ，其中 `n` 也是一个 **整数** 。

 

 **示例 1：** 

```

输入：[23,2,4,6,7], k = 6
输出：True
解释：[2,4] 是一个大小为 2 的子数组，并且和为 6。

```
 **示例 2：** 

```

输入：[23,2,6,4,7], k = 6
输出：True
解释：[23,2,6,4,7]是大小为 5 的子数组，并且和为 42。

```
 

 **说明：** 
- 数组的长度不会超过 `10,000` 。
- 你可以认为所有数字总和在 32 位有符号整数范围内。
 
**标签**
`数学` `动态规划` 


##  前缀和
```go
func isNK(check, k int) bool {
	return check%k == 0
}

func checkSubarraySum(nums []int, k int) bool {
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if isNK(preSum[j+1]-preSum[i], k) {
				return true
			}
		}
	}
	return false
}
```
`TLE`了，跪在了最后一个数据，$O(n^2)$ 的时间复杂度还是有点勉强。于是用哈希表来做优化，这是用了同余定理。如果 $(a-b) mod \ k == p$ 那么可以得到的是 $ a\ mod \ k == b\ mod \ k$

```go
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	preExist := map[int]int{0: -1}
	preSum := 0
	for i, v := range nums {
		preSum += v
		preSum %= k
		if preIndex, ok := preExist[preSum]; ok {
			if i-preIndex >= 2 {
				return true
			}
		} else {
			preExist[preSum] = i
		}
	}
	return false
}
```
>执行用时: 188 ms
内存消耗: 9.2 MB

记录下之前是否存在当前的结果，如果存在则返回。