# 1269.停在原地的方案数
[https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps](https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps) 
## 原题
有一个长度为 `arrLen` 的数组，开始有一个指针在索引 `0` 处。

每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。

给你两个整数 `steps` 和 `arrLen` ，请你计算并返回：在恰好执行 `steps` 次操作以后，指针仍然指向索引 `0` 处的方案数。

由于答案可能会很大，请返回方案数 **模** `10^9 + 7` 后的结果。

 

 **示例 1：** 

```
输入：steps = 3, arrLen = 2
输出：4
解释：3 步后，总共有 4 种不同的方法可以停在索引 0 处。
向右，向左，不动
不动，向右，向左
向右，不动，向左
不动，不动，不动

```
 **示例  2：** 

```
输入：steps = 2, arrLen = 4
输出：2
解释：2 步后，总共有 2 种不同的方法可以停在索引 0 处。
向右，向左
不动，不动

```
 **示例 3：** 

```
输入：steps = 4, arrLen = 2
输出：8

```
 

 **提示：** 
-  `1 <= steps <= 500` 
-  `1 <= arrLen <= 10^6` 
 
**标签**
`动态规划` 


## 
```go
func numWays(steps int, arrLen int) int {
	if steps == 1 {
		return 1
	}
	if arrLen-1 > steps {
		arrLen = steps
	}
	sum := make([]int, arrLen)
	tmpSum := make([]int, arrLen)
	sum[0] = 1
	for i := 0; i < steps; i++ {
		copy(tmpSum, sum)
		for j := 0; j < arrLen; j++ {
			if j > 0 {
				tmpSum[j] += sum[j-1]
				tmpSum[j] %= mod
			}
			if j < arrLen-1 {
				tmpSum[j] += sum[j+1]
				tmpSum[j] %= mod
			}
		}
		copy(sum, tmpSum)
	}
	return sum[0]
}

const mod = 1000000007
```
>执行用时: 8 ms
内存消耗: 2 MB

原理很简单，就是对每一个位置都记录每一步有多少种可能到就行。唯一坑人的地方就是 steps 很小，而数组 arrLen 很大。这个时候很大的 arrLen 是不可能到的，因此忽略掉即可。
