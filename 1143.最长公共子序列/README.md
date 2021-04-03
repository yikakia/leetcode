# 1143.最长公共子序列
[https://leetcode-cn.com/problems/longest-common-subsequence](https://leetcode-cn.com/problems/longest-common-subsequence) 
## 原题
给定两个字符串  `text1` 和  `text2` ，返回这两个字符串的最长 **公共子序列** 的长度。如果不存在 **公共子序列** ，返回 `0` 。

一个字符串的  **子序列**  是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
- 例如， `"ace"` 是 `"abcde"` 的子序列，但 `"aec"` 不是 `"abcde"` 的子序列。
两个字符串的 **公共子序列** 是这两个字符串所共同拥有的子序列。

 

 **示例 1：** 

```

输入：text1 = "abcde", text2 = "ace" 
输出：3  
解释：最长公共子序列是 "ace" ，它的长度为 3 。

```
 **示例 2：** 

```

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。

```
 **示例 3：** 

```

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。

```
 

 **提示：** 
-  `1 <= text1.length, text2.length <= 1000` 
-  `text1` 和  `text2` 仅由小写英文字符组成。
 
**标签**
`动态规划` 


## 
```go
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```
>执行用时: 0 ms
内存消耗: 10.7 MB

感觉应该是以前做过类似的，[115.不同的子序列](../115.不同的子序列/README.md)

对于子序列的问题，可能大部分都可以转化成二维数组的动规来做吧？

这题的思路和之前的题类似，就是用对应的下标记录当前对应的位置的最长公共子序列长度。

所以当找到了相同的字符的时候，这个时候的最长公共子序列就是它们两个之前的位置的长度加一。

而不同的时候，就是选择 text2 左边和当前的字符与 text1 当前的字符最长长度的最大值。