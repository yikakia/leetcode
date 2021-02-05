# 643.子数组最大平均数 I
[https://leetcode-cn.com/problems/maximum-average-subarray-i/](https://leetcode-cn.com/problems/maximum-average-subarray-i/) 
## 原题
给定 `n` 个整数，找出平均数最大且长度为 `k` 的连续子数组，并输出该最大平均数。

 

**示例：** 

```

输入：[1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75

```
 

**提示：** 
- 1 <= `k` <= `n` <= 30,000。
- 所给数据范围 [-10,000，10,000]。
 
**标签**
`数组` 


## 模拟
```go
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	max := 0
	for i := range nums[:k] {
		sum += nums[i]
	}
	max = sum
	for i := range nums[k:] {
		sum += nums[i+k] - nums[i]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
```
>执行用时: 140 ms
内存消耗: 7.1 MB

题目其实就是变化一下的求固定大小的窗口的最大和。没有太大的难度