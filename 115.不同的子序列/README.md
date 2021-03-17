# 115.不同的子序列
[https://leetcode-cn.com/problems/distinct-subsequences](https://leetcode-cn.com/problems/distinct-subsequences) 
## 原题
给定一个字符串 `s` 和一个字符串 `t` ，计算在 `s` 的子序列中 `t` 出现的个数。

字符串的一个 **子序列** 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如， `"ACE"`  是  `"ABCDE"`  的一个子序列，而  `"AEC"`  不是）

题目数据保证答案符合 32 位带符号整数范围。

 

 **示例 1：** 

```

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
(上箭头符号 ^ 表示选取的字母)
rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^

```
 **示例 2：** 

```

输入：s = "babgbag", t = "bag"
输出：5
解释：
如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
(上箭头符号 ^ 表示选取的字母)
babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
    ^^^
```
 

 **提示：** 
-  `0 <= s.length, t.length <= 1000` 
-  `s` 和 `t` 由英文字母组成
 
**标签**
`字符串` `动态规划` 


## 动态规划
```go
func numDistinct(s string, t string) int {
	ls, lt := len(s), len(t)
	if ls < lt {
		return 0
	}
	dp := make([][]int, ls)
	for i := range dp {
		dp[i] = make([]int, lt)
	}
	var maxInColAtRow func(col int, row int) int = func(col, row int) int {
		for i := row; i >= 0; i-- {
			if dp[i][col] > 0 {
				return dp[i][col]
			}
		}
		return 0
	}
	for i := 0; i < ls; i++ {
		for j := 0; j < lt && j <= i; j++ {
			if s[i] == t[j] {
				if j == 0 {
					dp[i][j] = 1 + maxInColAtRow(j, i-1)
				} else {
					dp[i][j] = maxInColAtRow(j-1, i-1) + maxInColAtRow(j, i-1)
				}
			}
		}
	}
	return maxInColAtRow(lt-1, ls-1)
}
```
>执行用时: 0 ms
内存消耗: 2.6 MB

做完才发现和题解的思路是反着来的。这题就是动态规划，状态转移方程如下

$$
  dp[i][j]=\begin{cases}
  1+max(dp[0:i][j])                 ,(s[i]==t[j]且j==0)\\
  max(dp[0:i][j-1])+max(dp[0:i][j]) ,(s[i]==t[j]且j!=0)\\
  0                                 ,(s[i]!=t[j])
  \end{cases}
$$

这里的 $[0:i]$ 表示的是从 $0$ 到 $i$ ，而不是 Go 中的 从 $0$ 到 $i-1$，是数学上的左闭右闭的区间。

我的思路是用 dp[i][j] 表示 s[i] 等于 t[j] 时，以它为结尾的有多少种组合。