# 474.一和零
[https://leetcode-cn.com/problems/ones-and-zeroes](https://leetcode-cn.com/problems/ones-and-zeroes) 
## 原题
给你一个二进制字符串数组 `strs` 和两个整数 `m` 和 `n` 。
<p class="MachineTrans-lang-zh-CN">请你找出并返回 `strs` 的最大子集的大小，该子集中 **最多** 有 `m` 个 `0` 和 `n` 个 `1` 。

<p class="MachineTrans-lang-zh-CN">如果 `x` 的所有元素也是 `y` 的元素，集合 `x` 是集合 `y` 的 **子集** 。
 

 **示例 1：** 

```

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。

```
 **示例 2：** 

```

输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。

```
 

 **提示：** 
-  `1 <= strs.length <= 600` 
-  `1 <= strs[i].length <= 100` 
-  `strs[i]`  仅由  `'0'` 和  `'1'` 组成
-  `1 <= m, n <= 100` 
 
**标签**
`动态规划` 


## 二维背包
```go
func findMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
>执行用时: 28 ms
内存消耗: 3.6 MB

就是二维背包问题，重量是`0`和`1`的个数，每个字符串的价值都是`1`。